import { useState } from 'react';

function UserSearch() {
    const [searchName, setSearchName] = useState('');
    const [searchResult, setSearchResult] = useState(null);
    const [error, setError] = useState(null);
  
    const handleSearch = async () => {
      try {
        const response = await fetch(`http://localhost:8080/search-users?name=${searchName}`);
        if (!response.ok) {
          throw new Error('No users with that name are present');
        }
        const data = await response.json();
        setSearchResult(data);
        setError(null);
      } catch (error) {
        setError(error.message);
        setSearchResult(null);
      }
    };
  
    return (
      <div>
        <input
          type="text"
          value={searchName}
          onChange={(e) => setSearchName(e.target.value)}
        />
        <button onClick={handleSearch}>Search</button>
        {error && <p>{error}</p>}
        {searchResult && (
          <div>
            {searchResult.map((user, index) => (
             <div key={index}>
             <p>
               Name: {user.name} 
               {user.age && `, Age: ${user.age}`} 
               {user.height && `, Height: ${user.height}`} 
               {user.weight && `, Weight: ${user.weight}`} 
               {user.gender && `, Gender: ${user.gender}`}
             </p>
           </div>
            ))}
          </div>
        )}
      </div>
    );
  }  

export default UserSearch;
