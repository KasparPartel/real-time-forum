// import React, { useEffect/* , useState */ } from "react";
import React from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsUserList/* , webSocketConnect */ } from "../../websocket.js"

function Userlist({user}) {
  // !!! this need to be rewritten to render compoment after wsUserList is retrieved from db
  // !!! also needs to get new wsUserList status periodically
  if (wsUserList && user) {
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
          {wsUserList.map((person) => (
            person.id !== user.id &&
              <ChatModal key={person.id} id={person.id} name={person.username} user={user}/>
          ))}
        </ul>
      </div>
    );
  } else {
    // empty user list sidebar if not logged in
    // !!! need to check for login
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
        </ul>
      </div>
    )
  }
}

export default Userlist;
