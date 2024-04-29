import React, { useState } from 'react';

const AddExercise = () => {
  const [exerciseData, setExerciseData] = useState({
    eName: '',
    purpose: '',
    calorieBurn: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setExerciseData({ ...exerciseData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/add-exercise', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(exerciseData)
      });

      if (response.ok) {
        console.log('Exercise added successfully');
        // Reset form fields
        setExerciseData({
          eName: '',
          purpose: '',
          calorieBurn: ''
        });
      } else {
        console.error('Error adding exercise:', response.statusText);
      }
    } catch (error) {
      console.error('Error adding exercise:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="eName" placeholder="Name" value={exerciseData.eName} onChange={handleChange} />
      <input type="text" name="purpose" placeholder="Purpose" value={exerciseData.purpose} onChange={handleChange} />
      <input type="text" name="calorieBurn" placeholder="Calorie Burn" value={exerciseData.calorieBurn} onChange={handleChange} />
      <button type="submit">Add Exercise</button>
    </form>
  );
};

export default AddExercise;
