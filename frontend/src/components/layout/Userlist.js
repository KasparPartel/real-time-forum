import React, { useState, useEffect } from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsUserList } from "../../websocket.js"

function Userlist({user}) {
  
  const [userlist, setUserlist] = useState(wsUserList)
  
  console.log("userlist:", userlist)
  console.log("wsUserList:", wsUserList)

  useEffect(() => {
    setUserlist(wsUserList)
  }, [])
  
  Userlist.setUserlist = setUserlist;

  if (user) {
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
          {userlist.map((target) => (
            target.id !== user.id &&
              <ChatModal key={target.id} id={target.id} name={target.username} user={user}/>
          ))}
        </ul>
      </div>
    );
  } else {
    // empty user list sidebar if not logged in
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
        </ul>
      </div>
    )
  }
}

export function usrUpdate() {
  Userlist.setUserlist(wsUserList)
}

export default Userlist;
