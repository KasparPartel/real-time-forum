import { useState, useEffect } from "react";

function Profile() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    requestUsers();
  }, []);

  const requestUsers = () => {
    fetch("http://localhost:4000/api/user/", {
      method: "GET",
      // headers: { "Content-Type": "application/json" },
      // body: JSON.stringify(),
    })
      .then((response) => {
        if (response.ok) {
          console.log("okay");
          return response.json();
        }
        throw new Error("Something went wrong");
      })
      .then((json) => setUsers(json))
      .catch((error) => console.log(error));
  };

  return (
    <div>
      <h1>This is your user profile.</h1>
      <ul>
        <li>Username: {users[0].username}</li>
        <li>E-mail: {users[0].email}</li>
        <li>First name: {users[0].first_name}</li>
        <li>Last name: {users[0].last_name}</li>
        <li>Gender: {users[0].gender}</li>
      </ul>
    </div>
  );
}

export default Profile;
