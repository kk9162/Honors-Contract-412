import React, { useState } from 'react';

function ExerciseSearch() {
    const [calorieBurn, setCalorieBurn] = useState('');
    const [searchResult, setSearchResult] = useState(null);
    const [error, setError] = useState(null);
  
    const handleSearch = async () => {
      try {
        const response = await fetch(`http://localhost:8080/search-exercises?calorieBurn=${calorieBurn}`);
        if (!response.ok) {
          throw new Error('No exercises with calorie burn greater than that amount are present');
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
          type="number"
          value={calorieBurn}
          onChange={(e) => setCalorieBurn(e.target.value)}
          placeholder="Enter minimum calorie burn"
        />
        <label>calories  </label>
        <button onClick={handleSearch}>Search</button>
        {error && <p>{error}</p>}
        {searchResult && (
          <div>
            {searchResult.map((exercise, index) => (
             <div key={index}>
             <p>
               Name: {exercise.eName} 
               {exercise.calorieBurn && `, Calorie Burn: ${exercise.calorieBurn}`} 
             </p>
           </div>
            ))}
          </div>
        )}
      </div>
    );
  }  

export default ExerciseSearch;
