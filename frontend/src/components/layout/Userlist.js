import React, { useState, useEffect } from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsUserList, /* wsActiveUserList, */ /* webSocketConnect, wsConnected */ } from "../../websocket.js"

function Userlist({user}) {
  
  console.log("Rendering Userlist!");
  
  const [userlist, setUserlist] = useState([])
  // const [activeuserlist, setActiveUserlist] = useState(wsActiveUserList)
  
  console.log("userlist:", userlist)
  // console.log("wsUserList:", wsUserList)

  useEffect(() => {
    setUserlist(wsUserList)
  }, [])

  Userlist.setUserlist = setUserlist;

  if (user && userlist) {
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
          {userlist.map((target) => (target.id !== user.id &&
              // <ChatModal class={target.class} key={target.id} id={target.id} name={target.username} user={user}/>
              <ChatModal active={target.active} newmessage={target.newmessage} targetkey={target.id} key={target.id} id={target.id} name={target.username} target={target} user={user}/>
          ))}
        </ul>
      </div>
    )
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

export function usrUpdate(userlist) {
  Userlist.setUserlist([...userlist])
  // Userlist.setActiveUserlist(wsActiveUserList)
}

export default Userlist;
