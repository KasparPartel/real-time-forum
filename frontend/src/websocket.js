import { msgUpdate } from "./components/layout/ChatModal";
import { usrUpdate } from "./components/layout/Userlist";
import { loggedUser } from "./App";

// import { UserContext } from "../../UserContext";
// import React, { useContext } from "react";

export let wsUserList = []
export let wsActiveUserList
export let wsMessageList = []
export let wsConnected = false

// let activeUser

// export const getActiveUser = (usr) => {
//   activeUser = usr
// }

export function webSocketConnect(port) {
    // const {user} = useContext(UserContext)

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection on port:", port);
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket on port:", port);
        wsConnected = true
        // console.log("Active user is:", activeUser);
        // console.log("Active user ID is:", activeUser.id);
        // socket.send(JSON.stringify(`"activeUserID":"${activeUser.id}"`))
        sendActiveUserID(loggedUser.id)
        wsGetUsers(loggedUser.id)
        // usrUpdate()
        
        if (loggedUser) {
            sendActiveUserID(loggedUser.id)
        }
    }
    
    socket.onclose = (e) => {
        console.log("WebSocket Connection Closed: ", e);
        wsConnected = false;
    }
    
    socket.onerror = (err) => {
        console.log("WebSocket Error: ", err);
        wsConnected = false;
    }

    socket.onmessage = (msg) => {
        console.log("Backend has responded: ", msg);
        // console.log("Backend has responded with data: ", msg.data);
        let incomingJson = JSON.parse(msg.data)

        // console.log(incomingJson);

        if (incomingJson.type === "wsReturnedUsers") {
            // wsUserList = wsSortUsers(loggedUser, incomingJson.body, incomingJson.pool, incomingJson.unread)
            console.log("loggedUser", loggedUser);   

            let tempBody = [...incomingJson.body]
            // let tempPool = [...incomingJson.pool]

            let sortedUsers = wsSortUsers(loggedUser, tempBody, incomingJson.pool/* , unreadstring */)
            // wsUserList = incomingJson.body
            // wsUserList = [...sortedUsers]
            // wsActiveUserList = incomingJson.pool
            console.log("incomingJson", incomingJson)
            console.log("sortedUsers", sortedUsers)
            usrUpdate([...sortedUsers])
        }
        if (incomingJson.type === "wsReturnedMessages") {
            console.log("returned messages incomingJson.body:", incomingJson.body);
            wsMessageList = incomingJson.body
            // let messages = incomingJson.body

            // msgUpdate(messages)
            // usrUpdate()
            msgUpdate()
        }
        if (incomingJson.type === "wsMessageSaved") {
            wsGetUsers(loggedUser.id)
        }

        console.log("wsUserList =", wsUserList);
        // console.log("wsMessageList =", wsMessageList);
    }

    webSocketConnect.sendMessage = sendMessage;
    webSocketConnect.wsGetUsers = wsGetUsers;
    webSocketConnect.wsGetChatMessages = wsGetChatMessages;
    webSocketConnect.sendActiveUserID = sendActiveUserID;
    webSocketConnect.sendModal = sendModal;
    webSocketConnect.wsSortUsers = wsSortUsers;

    function sendMessage() {
        function composeMessage(Type, Body, User_id, Target_id, Creation_time) {
            let msg = {
                type: String(Type),
                body: String(Body),
                user_id: String(User_id),
                target_id: String(Target_id),
                creation_time: String(Creation_time),
            };
            return JSON.stringify(msg);
        }

        let newMessage = composeMessage(
        "wsSaveChatMessage",
        document.querySelector("#chat-text").value,
        document.querySelector("#send-button").getAttribute("data-user-id"),
        document.querySelector("#send-button").getAttribute("data-target-id"),
        Date(Date.now()) 
        );

        socket.send(newMessage);
            
        wsGetChatMessages(
            document.querySelector("#send-button").getAttribute("data-user-id"),
            document.querySelector("#send-button").getAttribute("data-target-id")
        )
        msgUpdate()

        document.getElementById("chat-text").textContent = "";
    }

    function wsGetUsers(usrID) {
        //JSON for getting users from db query
        let msg = {
        type: "wsGetUsers",
        activeUser: usrID,
        };

        socket.send(JSON.stringify(msg));
    }
    
    function sendActiveUserID(usrID) {
        //JSON for getting users from db query
        let msg = {
            type: "sendUser",
            activeUser: usrID,
        };
        console.log("sending usrID", usrID);
        if (usrID) {
            socket.send(JSON.stringify(msg));
            console.log("Sent ActiveUserID over websocket to backend:", usrID);
        } else {
            console.log("Error: no ActiveUserID to send.");
        }
    }

    function sendModal(usrID, trgtID/* , messageLength */) {
        // console.log("Trying to send Modal: user, target, length", usrID, trgtID, messageLength);
        console.log("Trying to send Modal: user, target", usrID, trgtID);
        // JSON for sending user Modal message length
        let msg = {
            type: "sendModal",
            activeUser: usrID,
            targetUser: trgtID,
            // modalLength: messageLength,
        };
        if (usrID && trgtID /* && messageLength */) {
            socket.send(JSON.stringify(msg));
            console.log("sending Modal: user, target", usrID, trgtID);
            // console.log("sending Modal: user, target, length", usrID, trgtID, messageLength);
            // console.log("Sent ActiveUserID over websocket to backend:", usrID);
        } else {
            console.log("Error: Modal attribute missing when sending");
        }

        // wsGetUsers(usrID);
    }
    
    function wsGetChatMessages(usr, trgt) {
        console.log("Sending wsGetChatMessages", usr, trgt);
        let msg = {
            type: "wsGetChatMessages",
            // body: "empty",
            user_id: String(usr),
            target_id: String(trgt),
        };     
        socket.send(JSON.stringify(msg));
    }

    function wsSortUsers(mainUser, usersList, activeUsersList) {

        // console.log("wsSortUsers(mainUser, usersList, activeUsersList, unreadUsersList):",
        // mainUser, usersList, activeUsersList, unreadUsersList);

        if (!mainUser || !usersList || !activeUsersList) {
            console.log("Error: user sorting data missing");
            return
        }

        if (mainUser.history.length === 0) {
            return usersList.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))    
        }

        console.log("wsSortUsers started!");
        console.log("mainUser:", mainUser);
        console.log("usersList:", usersList);
        console.log("activeUsersList:", activeUsersList);
        // console.log("unreadUsersList:", unreadUsersList);

        // let activeusers = []
        // let passiveusers = []
        // let activehistory = []
        // let passivehistory = []
        let historyUsers = []
        let strangerUsers = []
        let orderedUsers = []
        // let activenames = []
        // let passivenames = []
        // let activesorted = []
        // let passivesorted = []
        let combinedUsers = []
        let historyarray = []
        let activeUserArray = activeUsersList.split(",").map(function(item) {return parseInt(item, 10);})
        // let unreadUserArray = unreadUsersList.split(",").map(function(item) {return parseInt(item, 10);})
        let unreadUserArray = []
        // let historySplit = mainUser.history.split(",") // Error! Main User has old history!!!
        let historySplit = []
        let mainUserHistory

        usersList.forEach((usr) => {
            if (usr.id === mainUser.id) {
                mainUserHistory = usr.history
            }
        })
        historySplit = mainUserHistory.split(",")

        console.log("historySplit", historySplit);
        
        historySplit.forEach((el) => {
            console.log("el.split("-")[1]", el.split("-")[1]);
            if (el.split("-")[1] === "1") {
                console.log("el.split("-")[0]", el.split("-")[0]);
                unreadUserArray.push(parseInt(el.split("-")[0]))
            }
        })

        console.log("activeUserArray", activeUserArray);
        console.log("unreadUserArray", unreadUserArray);

        usersList.forEach((usr) => {
            activeUserArray.forEach((loginID) => {
                if (usr.id === loginID) {
                usr.active = true
                // activeusers.push(usr)
                }
            })
            unreadUserArray.forEach((unreadID) => {
                if (usr.id === unreadID) {
                usr.newmessage = true
                }
            })
        })

        // usersList.forEach((usr) => {
        //     if (!activeusers.includes(usr) && !passiveusers.includes(usr)) {
        //         usr.class = "passive"
        //         passiveusers.push(usr)
        //     }
        // })

        // if (mainUser.history.length === 0) {
        //     console.log("This user has no prior chat history!");
        //     activesorted = activeusers.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))
        //     passivesorted = passiveusers.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))
        //     combinedUsers = activesorted.concat(passivesorted)    
        //     return combinedUsers
        // }

        // if (user && user.history !== undefined && user != null) {
        //     historyarray = user.history.split(",").flatMap((item) => item === "" ? [] : parseInt(item, 10));
        // }
        // historyarray = mainUser.history.split(",").split("-")[0].flatMap((item) => item === "" ? [] : parseInt(item, 10));
        historySplit.forEach((el) => {
            historyarray.push(parseInt(el.split("-")[0], 10));
        }) 

        console.log("historyarray", historyarray);
        // console.log("activeusers", activeusers);
        // console.log("passiveusers", passiveusers);
        // console.log("activenames", activenames);
        // console.log("activehistory", activehistory);
        // console.log("passivenames", passivenames);
        // console.log("passivehistory", passivehistory);
        
        usersList.forEach((usr) => {
            historyarray.forEach((loginID) => {
                if (usr.id === loginID) {
                    historyUsers.push(usr)
                }
            })
            historyarray.forEach(() => {
                if (!strangerUsers.includes(usr) && !historyUsers.includes(usr)) {
                    strangerUsers.push(usr)
                }
            })
        })
        
        orderedUsers = strangerUsers.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))
        combinedUsers = historyUsers.concat(orderedUsers);

        console.log("historyUsers", historyUsers);
        console.log("strangerUsers", strangerUsers);
        console.log("orderedUsers", orderedUsers);

        // historyarray.forEach((item) => {
            // console.log("historyarray item", item);
            // console.log("typeof historyarray item", typeof(item));
            
                     
            // activeusers.forEach((usr) => {
            //     if (usr.id === item) {
            //         activehistory.push(usr)
            //     } else if (!historyarray.includes(usr.id)) {
            //         let includes = false
            //         activenames.forEach((el) => {
            //             if (usr.id === el.id) includes = true;
            //         })
            //         if (!includes) activenames.push(usr);
            //     }
            // })
            // passiveusers.forEach((usr) => {
            //     if (usr.id === item) {
            //         passivehistory.push(usr)
            //     } else if (!historyarray.includes(usr.id)) {
            //         let includes = false
            //         passivenames.forEach((el) => {
            //             if (usr.id === el.id) includes = true;
            //         })
            //         if (!includes) passivenames.push(usr);
            //     }
            // })
        // })
        
        // console.log("historyarray", historyarray);
        // console.log("activenames", activenames);
        // console.log("activehistory", activehistory);
        // console.log("passivenames", passivenames);
        // console.log("passivehistory", passivehistory);

        // activesorted = activenames.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))
        // passivesorted = passivenames.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))

        // combinedUsers = activehistory.concat(activesorted, passivehistory, passivesorted)
        
        return combinedUsers
    }
}
