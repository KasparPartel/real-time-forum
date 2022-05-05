import React, { useEffect, useState } from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";

function Userlist() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    requestUsers();
  }, []);

  const requestUsers = () => {
    fetch("http://localhost:4000/v1/api/user/", {
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
    <div className="user-list">
      <ul className={classes.userlist}>
        {users.map((user) => (
          //<Post key={p.id} json={p} />
          //<li className={classes.online}>{user.username}</li>
          <ChatModal key={user.id} target={user.id} name={user.username} />
        ))}
      </ul>
    </div>
  );
}

export default Userlist;
