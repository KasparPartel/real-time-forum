import React, { useState, useEffect } from "react";
import classes from "./Userlist.module.css";
import ChatModal from "./ChatModal";
import { wsUserList, /* wsActiveUserList, */ /* webSocketConnect, wsConnected */ } from "../../websocket.js"

function Userlist({user}) {
  
  
  const [userlist, setUserlist] = useState(wsUserList)
  // const [activeuserlist, setActiveUserlist] = useState(wsActiveUserList)
  
  console.log("userlist:", userlist)
  console.log("wsUserList:", wsUserList)
  
  // console.log("userlist[3]?.username:", userlist[3]?.username)
  // console.log("userlist[3]?.login_date", userlist[3]?.login_date);
  // console.log("Date.parse(target.login_date)", Date.parse(userlist[3]?.login_date))
  // console.log("Date.parse(target.logout_date)", Date.parse(userlist[3]?.logout_date));
  // console.log("active", Date.parse(userlist[3]?.login_date) > Date.parse(userlist[3]?.logout_date));
  // console.log("passive", Date.parse(userlist[3]?.login_date) < Date.parse(userlist[3]?.logout_date));

  // console.log("userlist[2]?.username:", userlist[2]?.username)
  // console.log("userlist[2]?.login_date", userlist[2]?.login_date);
  // console.log('userlist[2]?.login_date..substring(0, 19)replaceAll("-", "/")', userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/"))
  // console.log("Date.parse(target.login_date)", Date.parse(userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/")))
  // console.log("Date.parse(target.logout_date)", Date.parse(userlist[2]?.logout_date.substring(0, 19).replaceAll("-", "/")));
  // console.log("active", Date.parse(userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/")) > Date.parse(userlist[2]?.logout_date.substring(0, 19).replaceAll("-", "/")));
  // console.log("passive", Date.parse(userlist[2]?.login_date.substring(0, 19).replaceAll("-", "/")) < Date.parse(userlist[2]?.logout_date.substring(0, 19).replaceAll("-", "/")));



  useEffect(() => {
    setUserlist(wsUserList)
  }, [])
  // useEffect(() => {
  //   if (user !== null && user.id !== undefined && wsConnected) {
  //     webSocketConnect.sendActiveUserID(user.id);
  //   }
  // }, [user])

  // userlist.map((target) => (() => {
            
  //   console.log("target.id", target.id);
  //   console.log("target.username", target.username);
  //   console.log("Date.parse(target.login_date)", Date.parse(target.login_date));
  //   console.log("Date.parse(target.logout_date)", Date.parse(target.logout_date));
  // }))

  // let activeusers = []
  // let passiveusers = []
  // let activehistory = []
  // let passivehistory = []
  // let activenames = []
  // let passivenames = []
  // let activesorted = []
  // let passivesorted = []
  // let combinedUsers = []
  // let activeUserArray = activeuserlist?.split(",").map(function(item) {return parseInt(item, 10);})

  // userlist?.forEach((usr) => {
  //   activeUserArray.forEach((loginID) => {
  //     if (usr.id === loginID) {
  //       usr.class = "active"
  //       activeusers.push(usr)
  //     }
  //   })
  // })
  // userlist?.forEach((usr) => {
  //   if (!activeusers.includes(usr) && !passiveusers.includes(usr)) {
  //     usr.class = "passive"
  //     passiveusers.push(usr)
  //   }
  // })

  
  // console.log("activeusers:", activeusers);
  // console.log("passiveusers:", passiveusers);
  // console.log("Logged in users connected to websocket pool:", activeUserArray);
  // console.log("Logged in user history:", user?.history);
  // console.log("typeof Logged in user history:", typeof(user?.history));
  
  // let historyarray = []
  // if (user.history !== undefined) {
  // if (user.history !== undefined && user != null) {
  // if (user && user.history !== undefined && user != null) {
  //   historyarray = user.history.split(",").flatMap((item) => item === "" ? [] : parseInt(item, 10));
  // }
  // console.log("Logged in user history array:", historyarray);

  // historyarray.forEach((item) => {
  //   activeusers.forEach((usr) => {
  //     if (usr.id !== item && !activenames.includes(usr)) {
  //       activenames.push(usr)
  //     }
  //     if (usr.id === item) {
  //       activehistory.push(usr)
  //     }
  //   })
  //   passiveusers.forEach((usr) => {
  //     if (usr.id !== item && !passivenames.includes(usr)) {
  //       passivenames.push(usr)
  //     }
  //     if (usr.id === item) {
  //       passivehistory.push(usr)
  //     }
  //   })
  // })

  // activesorted = activenames.sort()
  // passivesorted = passivenames.sort()

  // activesorted = activenames.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))
  // passivesorted = passivenames.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))

  // console.log("activehistory", activehistory);
  // console.log("passivehistory", passivehistory);
  // console.log("activesorted", activesorted);
  // console.log("passivesorted", passivesorted);

  // combinedUsers = activehistory.concat(activesorted)
  // combinedUsers = activehistory.concat(activesorted, passivehistory, passivesorted)
  // combinedUsers = userlist
  // let combinedUsers = []
  // if (wsConnected) {
  //   combinedUsers = webSocketConnect.wsSortUsers(user, userlist, activeuserlist)
  // }
  // console.log("combinedUsers", combinedUsers);

  Userlist.setUserlist = setUserlist;
  // Userlist.setActiveUserlist = setActiveUserlist;

  if (user && userlist) {
    return (
      <div className="user-list">
        <ul className={classes.userlist}>
          {userlist.map((target) => (target.id !== user.id &&
              <ChatModal class={target.class} key={target.id} id={target.id} name={target.username} user={user}/>
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

export function usrUpdate() {
  Userlist.setUserlist(wsUserList)
  // Userlist.setActiveUserlist(wsActiveUserList)
}

export default Userlist;
