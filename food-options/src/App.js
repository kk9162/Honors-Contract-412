import AddUser from './AddUser';
import AddExercise from './AddExercise';
import UserSearch from './SearchUser';
import UserSearchByExercise from './SearchUserByExercise';
import './App.css';

function App() {
  return (
    <div>
      <h1>Food Options</h1>
      <h3> Add a User</h3>
      <AddUser></AddUser>
      <h3> Add an Exercise </h3>
      <AddExercise></AddExercise>
      <h3> Search User by Name</h3>
      <UserSearch></UserSearch>
      <h3> Search Exercises That Burn More Than: </h3>
      <UserSearchByExercise></UserSearchByExercise>
   
    </div>
  );
}

export default App;
