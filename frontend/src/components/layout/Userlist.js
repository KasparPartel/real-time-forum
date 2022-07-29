import React, { useState } from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsMessageList } from "../../websocket.js"

function Userlist({user}) {
  // monitors change of userlist (set onmessage in websocket.js)
  const [userlist, setUserlist] = useState([])
  // exposes setUserlist function to usrUpdate
  Userlist.setUserlist = setUserlist;

  if (user && userlist) {
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
          {userlist.map((target) => (target.id !== user.id &&
              <ChatModal messages={wsMessageList} active={target.active} newmessage={target.newmessage} targetkey={target.id} key={target.id} id={target.id} name={target.username} target={target} user={user}/>
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

export function usrUpdate(list) {
  // called from websocket.js with userlist change from db (onmessage)
  if (list) {
    Userlist.setUserlist(list)
  }
}

export default Userlist;