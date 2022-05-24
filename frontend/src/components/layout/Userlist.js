// import React, { useEffect/* , useState */ } from "react";
import React from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsUserList/* , webSocketConnect */ } from "../../websocket.js"

function Userlist({user}) {
  // const [users, setUsers] = useState([]);

  // useEffect(() => {
  //   requestUsers();
  // }, []);

  // const requestUsers = () => {
  //   fetch("http://localhost:4000/v1/api/user/", {
  //     method: "GET",
  //     // headers: { "Content-Type": "application/json" },
  //     // body: JSON.stringify(),
  //   })
  //     .then((response) => {
  //       if (response.ok) {
  //         console.log("okay");
  //         return response.json();
  //       }
  //       throw new Error("Something went wrong");
  //     })
  //     .then((json) => setUsers(json))
  //     .catch((error) => console.log(error));
  // };

  // console.log(users);

  // useEffect(() => {
  //   webSocketConnect.wsGetUsers();
  //   console.log(wsUserList);
  // }, []); 

  if (wsUserList) {
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
          {wsUserList.map((person) => (
            //<Post key={p.id} json={p} />
            //<li className={classes.online}>{user.username}</li>
            <ChatModal key={person.id} name={person.username} user={user}/>
          ))}
        </ul>
      </div>
    );
  } else {
    // empty user list sidebar
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
        </ul>
      </div>
    )
  }
}

export default Userlist;
