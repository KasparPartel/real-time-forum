import React, { useEffect, useState } from "react";
import classes from "./Userlist.module.css";

function Userlist() {
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
    <div className="flex flex-wrap py-5 gap-4">
      <ul className={classes.userlist}>
        {users.map((user) => (
          //<Post key={p.id} json={p} />
          <li className={classes.online}>{user.username}</li>
        ))}
      </ul>
    </div>
  );
}

export default Userlist;
