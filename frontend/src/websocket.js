import { usrUpdate } from "./components/layout/Userlist";
import { loggedUser } from "./App";

export let wsUserList = []
export let wsActiveUserList
export let wsMessageList = {}
export let wsConnected = false
let sortedUsers = []

// creates a websocket connection to backend
export function webSocketConnect(port) {
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection on port:", port);
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket on port:", port);
        // bool is checked in functions, if websocket is connected or not 
        wsConnected = true
        // tells the backend, which user is connected over this websocket
        sendActiveUserID(loggedUser.id)
        // gets userlist from db
        wsGetUsers(loggedUser.id)
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
        let incomingJson = JSON.parse(msg.data)

        // if users are returned over wbesocket, they are sorted and sent to Userlist.js
        if (incomingJson.type === "wsReturnedUsers") {
            // .body is userlist, .pool is string of active websocket users connected to backend
            sortedUsers = wsSortUsers(loggedUser, incomingJson.body, incomingJson.pool)
            usrUpdate(sortedUsers)
        }
        // if messages are returned over websocket, they are saved into wsMessageList object with id as key
        if (incomingJson.type === "wsReturnedMessages") {
            let key = parseInt(incomingJson.target)
            wsMessageList[key] = {}
            wsMessageList[key].body = incomingJson.body
            wsMessageList[key].count = incomingJson.body?.length
        }
        // saving message into db changes order of userlist, therefore new query
        if (incomingJson.type === "wsMessageSaved") {
            wsGetUsers(loggedUser.id)
        }
        // saving modal click into db changes flags of userlist, therefore new query
        if (incomingJson.type === "wsModalSaved") {
            wsGetUsers(loggedUser.id)
        }
    }

    // this exposes sub-functions to export
    webSocketConnect.sendMessage = sendMessage;
    webSocketConnect.wsGetUsers = wsGetUsers;
    webSocketConnect.wsGetChatMessages = wsGetChatMessages;
    webSocketConnect.sendActiveUserID = sendActiveUserID;
    webSocketConnect.sendModal = sendModal;

    // this sends save chat message to db
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
        
        // clears modal message field
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
        //JSON for connecting user and websocket connection in backend
        let msg = {
            type: "sendUser",
            activeUser: usrID,
        };

        if (usrID) {
            socket.send(JSON.stringify(msg));
        } else {
            console.log("Error: no ActiveUserID to send");
        }
    }

    function sendModal(usrID, trgtID) {
        // JSON for sending user Modal click to db (to change userlist flags in db)
        let msg = {
            type: "sendModal",
            activeUser: usrID,
            targetUser: trgtID,
        };
        if (usrID && trgtID) {
            socket.send(JSON.stringify(msg));
        } else {
            console.log("Error: Modal attribute missing when sending");
        }
    }
    
    function wsGetChatMessages(usr, trgt, count) {
        // JSON query to db for a list of messages between users
        if (count === undefined || !wsConnected) {
            return
        }
        if (count < 10) {
            count = 10
        }
        let msg = {
            type: "wsGetChatMessages",
            user_id: String(usr),
            target_id: String(trgt),
            count: String(count),
        };    
        socket.send(JSON.stringify(msg));
    }

    function wsSortUsers(mainUser, usersList, activeUsersList) {
        // sorts users in incoming userlist (from db):
        // first: last user chatted with
        // next: users chatted with in chronological order
        // last: users not chatted with in alphabetical order
        // active users (connected websocket) get "active: true"
        // users with unread messages get "newmessage: true"

        if (!mainUser || !usersList || !activeUsersList) {
            console.log("Error: user sorting data missing");
            return
        }

        if (mainUser.history.length === 0) {
            return usersList.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))    
        }

        let historyUsers = []
        let strangerUsers = []
        let orderedUsers = []
        let combinedUsers = []
        let historyarray = []
        let unreadUserArray = []
        let historySplit = []
        let mainUserHistory
        // filters into array connected webscoket user id-s
        let activeUserArray = activeUsersList.split(",").map(function(item) {return parseInt(item, 10);})
        // finds history string of logged in user
        usersList.forEach((usr) => {
            if (usr.id === mainUser.id) {
                mainUserHistory = usr.history
            }
        })
        historySplit = mainUserHistory.split(",")
        // filters into array which users have unread messages for logged in user
        historySplit.forEach((el) => {
            if (el.split("-")[1] === "1") {
                unreadUserArray.push(parseInt(el.split("-")[0]))
            }
        })
        // adds active key and newmessage key to users based on filtered arrays
        usersList.forEach((usr) => {
            activeUserArray.forEach((loginID) => {
                if (usr.id === loginID) {
                usr.active = true
                }
            })
            unreadUserArray.forEach((unreadID) => {
                if (usr.id === unreadID) {
                usr.newmessage = true
                }
            })
        })
        // truns history string into array of id-s for chronological order
        historySplit.forEach((el) => {
            historyarray.push(parseInt(el.split("-")[0], 10));
        }) 
        // pushed users with prior history into ordered array
        historyarray.forEach((historyID) => {
            for (let i = 0; i < usersList.length; i++) {
                if (historyID === usersList[i].id) {
                    historyUsers.push(usersList[i])
                    usersList.splice(i, 1)
                } 
            }
        })
        // orders users with no history aplhabetically
        strangerUsers = [...usersList]
        orderedUsers = strangerUsers.sort((a,b) => (a.username > b.username) ? 1 : ((b.username > a.username) ? -1 : 0))
        // combines users arrays with and without prior history
        combinedUsers = historyUsers.concat(orderedUsers);

        return combinedUsers
    }
}
