package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

const (
   
)

// go struct that represents a User
type User struct {
    Name     string `json:"name"`
    Age      string `json:"age"`
    Height   string `json:"height"`
    Weight   string `json:"weight"`
    Gender   string `json:"gender"`
}

// go struct that represent an Exercise
type Exercise struct {
    ExerciseName     string `json:"eName"`
    Purpose  string `json:"purpose"`
    Calorie  string `json:"calorieBurn"`
}

var db *sql.DB 

// initiate a connection to sql database
func init() {

    // Open a connection to the database. Open function returns pointer to database and also error if any
    var err error
    db, err = sql.Open("mysql", "root:412@tcp(localhost:3306)/HONORS_Food_DB")  // password is 412 and database is
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    // health check for database
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connection to the database successful!")
}


func main() {
    // init will run automatically

    // close the database connection
    defer db.Close()

    // allow multiple endpoints. 
    r := mux.NewRouter()

	// matches enpdoint to function
    r.HandleFunc("/add-user", addUserHandler).Methods("POST") 
	r.HandleFunc("/add-exercise", addExerciseHandler).Methods("POST")
	r.HandleFunc("/search-users", searchUsersHandler).Methods("GET")
	r.HandleFunc("/search-exercises", searchExercisesHandler).Methods("GET")


    // Enable CORS
    // cors is like the gatekeeper for backend. when the front end sends a request to backend, 
    // cors will ensure that the origin is the same. to solve error, let requests come from localhost:3000
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Allow requests from localhost:3000
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )

    // Set up primary HTTP server
    http.Handle("/", corsHandler(r))
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func addExerciseHandler(w http.ResponseWriter, r *http.Request) {
    // create a new exercise
    var newExercise Exercise

    // decode the JSON data from HTTP request into an exercise struct
    err := json.NewDecoder(r.Body).Decode(&newExercise)
    if err != nil {
        http.Error(w, "Error decoding request body", http.StatusBadRequest)
        return
    }


	// Print the contents of the newExercise variable
fmt.Printf("Exercise Name: %s\n", newExercise.ExerciseName)
fmt.Printf("Exercise Purpose: %s\n", newExercise.Purpose)
fmt.Printf("Exercise Calorie Burn: %s\n", newExercise.Calorie)

    // Insert exercise into database
    result, err := db.Exec("INSERT INTO exercise (eName, purpose, calorieBurn) VALUES (?, ?, ?)",
		newExercise.ExerciseName, newExercise.Purpose, newExercise.Calorie)
    if err != nil {
        http.Error(w, "Error inserting exercise into database", http.StatusInternalServerError)
        log.Println("Error inserting exercise into database:", err)
        return
    }

    // Check the number of rows affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "Error retrieving rows affected", http.StatusInternalServerError)
        log.Println("Error retrieving rows affected:", err)
        return
    }

    fmt.Printf("%d row(s) inserted\n", rowsAffected)
    fmt.Println("Successfully inserted a new exercise!")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newExercise)
}


func addUserHandler(w http.ResponseWriter, r *http.Request) {
    // create a new user
    var newUser User

    // decode the JSON data from HTTP request into a User struct
    err := json.NewDecoder(r.Body).Decode(&newUser)
    
    if err != nil {
        http.Error(w, "Error decoding request body", http.StatusBadRequest)
        return
    }
	
   // Insert user into database
	result, err := db.Exec("INSERT INTO users (userName, userAge, userHeight, userWeight, userGender) VALUES (?, ?, ?, ?, ?)",
	newUser.Name, newUser.Age, newUser.Height, newUser.Weight, newUser.Gender)
	if err != nil {
	http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
	log.Println("Error inserting user into database:", err)
	return
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
	http.Error(w, "Error retrieving rows affected", http.StatusInternalServerError)
	log.Println("Error retrieving rows affected:", err)
	return
	}

	fmt.Printf("%d row(s) inserted\n", rowsAffected)
    fmt.Println("Successfully inserted a new user!")


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

func searchUsersHandler(w http.ResponseWriter, r *http.Request) {
    // Get the user's name from the query parameter
    name := r.URL.Query().Get("name")

    // Query the database for users with the provided name
    rows, err := db.Query("SELECT userName, userAge, userHeight, userWeight, userGender FROM users WHERE userName = ?", name)
    if err != nil {
        http.Error(w, "Error searching for users", http.StatusInternalServerError)
        log.Println("Error searching for users:", err)
        return
    }
    defer rows.Close()

    // Iterate over the query results and collect them in a slice
    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Name, &user.Age, &user.Height, &user.Weight, &user.Gender); err != nil {
            http.Error(w, "Error scanning users", http.StatusInternalServerError)
            log.Println("Error scanning users:", err)
            return
        }
        users = append(users, user)
    }

    // Encode the search results as JSON and send the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func searchExercisesHandler(w http.ResponseWriter, r *http.Request) {
    calorieBurn := r.URL.Query().Get("calorieBurn")

    // Query the database for exercises with the provided calories
    rows, err := db.Query("SELECT eName, calorieBurn FROM exercise WHERE calorieBurn > ?", calorieBurn)
    if err != nil {
        http.Error(w, "Error searching for exercises", http.StatusInternalServerError)
        log.Println("Error searching for exercises:", err)
        return
    }
    defer rows.Close()

    // Iterate over the query results and collect them in a slice
	var exercises []Exercise
    for rows.Next() {
        var exercise Exercise
        if err := rows.Scan(&exercise.ExerciseName, &exercise.Calorie); err != nil {
            http.Error(w, "Error scanning exercises", http.StatusInternalServerError)
            log.Println("Error scanning exercises:", err)
            return
        }
        exercises = append(exercises, exercise)
    }

    // Encode the search results as JSON and send the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(exercises)
}


