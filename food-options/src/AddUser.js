import React, { useState } from 'react';

const AddUserForm = () => {
  const [userData, setUserData] = useState({
    name: '',
    age: '',
    height: '',
    weight: '',
    gender: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUserData({ ...userData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/add-user', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(userData)
      });

      if (response.ok) {
        console.log('User added successfully');
        // Reset form fields
        setUserData({
          name: '',
          age: '',
          height: '',
          weight: '',
          gender: ''
        });
      } else {
        console.error('Error adding user:', response.statusText);
      }
    } catch (error) {
      console.error('Error adding user:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="name" placeholder="Name" value={userData.name} onChange={handleChange} />
      <input type="text" name="age" placeholder="Age" value={userData.age} onChange={handleChange} />
      <input type="text" name="height" placeholder="Height in CM" value={userData.birthday} onChange={handleChange} />
      <input type="text" name="weight" placeholder="Weight in KG" value={userData.email} onChange={handleChange} />
      <select name="gender" value={userData.gender} onChange={handleChange} required>
        <option value="">Select Gender</option>
        <option value="FEMALE">Female</option>
        <option value="MALE">Male</option>
      </select>     
        <button type="submit">Add User</button>
    </form>
  );
};

export default AddUserForm;