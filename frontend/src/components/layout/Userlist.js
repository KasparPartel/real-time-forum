import React, { useState, useEffect } from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsUserList } from "../../websocket.js"

function Userlist({user}) {
  
  const [userlist, setUserlist] = useState(wsUserList)
  
  console.log("userlist:", userlist)
  console.log("wsUserList:", wsUserList)
  
  console.log("userlist[3]?.username:", userlist[3]?.username)
  console.log("userlist[3]?.login_date", userlist[3]?.login_date);
  console.log("Date.parse(target.login_date)", Date.parse(userlist[3]?.login_date))
  console.log("Date.parse(target.logout_date)", Date.parse(userlist[3]?.logout_date));
  console.log("active", Date.parse(userlist[3]?.login_date) > Date.parse(userlist[3]?.logout_date));
  console.log("passive", Date.parse(userlist[3]?.login_date) < Date.parse(userlist[3]?.logout_date));

  console.log("userlist[2]?.username:", userlist[2]?.username)
  console.log("userlist[2]?.login_date", userlist[2]?.login_date);
  console.log('userlist[2]?.login_date..substring(0, 19)replaceAll("-", "/")', userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/"))
  console.log("Date.parse(target.login_date)", Date.parse(userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/")))
  console.log("Date.parse(target.logout_date)", Date.parse(userlist[2]?.logout_date.substring(0, 19).replaceAll("-", "/")));
  console.log("active", Date.parse(userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/")) > Date.parse(userlist[2]?.logout_date.substring(0, 19).replaceAll("-", "/")));
  console.log("passive", Date.parse(userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/")) < Date.parse(userlist[2]?.logout_date.substring(0, 19).replaceAll("-", "/")));



  useEffect(() => {
    setUserlist(wsUserList)
  }, [])

  // userlist.map((target) => (() => {
            
  //   console.log("target.id", target.id);
  //   console.log("target.username", target.username);
  //   console.log("Date.parse(target.login_date)", Date.parse(target.login_date));
  //   console.log("Date.parse(target.logout_date)", Date.parse(target.logout_date));
  // }))

  let activeusers = []
  let passiveusers = []

  userlist.forEach((usr) => {
    if (Date.parse(usr.login_date.substring(0, 19).replaceAll("-", "/")) > 
    Date.parse(usr.logout_date.substring(0, 19).replaceAll("-", "/"))) {
      activeusers.push(usr)
      usr.class = "active"
    }
    if (Date.parse(usr.login_date.substring(0, 19).replaceAll("-", "/")) <= 
    Date.parse(usr.logout_date.substring(0, 19).replaceAll("-", "/"))) {
      passiveusers.push(usr)
      usr.class = "passive"
    }
  })

  console.log("activeusers:", activeusers);
  console.log("passiveusers:", passiveusers);
  
  Userlist.setUserlist = setUserlist;

  if (user && userlist) {
    return (
      //<div className="user-list">
        <ul className={classes.userlist}>
          {userlist.map((target) => (target.id !== user.id &&
              <ChatModal class={target.class} key={target.id} id={target.id} name={target.username} user={user}/>
          ))}
        </ul>
      //</div>
      // ,
      // <div className="user-list">
      //   <ul className={classes.passiveuserlist}>
      //     {passiveusers.map((target) => (
      //       (target.id !== user.id && Date.parse(target.login_date.substring(0, 19).replaceAll("-", "/")) <= 
      //       Date.parse(target.logout_date.substring(0, 19).replaceAll("-", "/"))) &&
      //         <ChatModal key={target.id} id={target.id} name={target.username} user={user}/>
      //     ))}
      //   </ul>
      // </div>
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
