import { msgUpdate } from "./components/layout/ChatModal";
import { usrUpdate } from "./components/layout/Userlist";
// import { UserContext } from "../../UserContext";
// import React, { useContext } from "react";

export let wsUserList = []
export let wsActiveUserList
export let wsMessageList = []

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
        // console.log("Active user is:", activeUser);
        // console.log("Active user ID is:", activeUser.id);
        // socket.send(JSON.stringify(`"activeUserID":"${activeUser.id}"`))
        wsGetUsers()
        usrUpdate()
        // sendActiveUserID(usrID)
    }
    
    socket.onclose = (e) => {
        console.log("WebSocket Connection Closed: ", e);
    }
    
    socket.onerror = (err) => {
        console.log("WebSocket Error: ", err);
    }

    socket.onmessage = (msg) => {
        console.log("Backend has responded: ", msg);
        // console.log("Backend has responded with data: ", msg.data);
        let incomingJson = JSON.parse(msg.data)

        console.log(incomingJson);

        if (incomingJson.type === "wsReturnedUsers") {
            wsUserList = incomingJson.body
            wsActiveUserList = incomingJson.pool
            usrUpdate()
        }
        if (incomingJson.type === "wsReturnedMessages") {
            console.log("returned messages incomingJson.body:", incomingJson.body);
            wsMessageList = incomingJson.body
            msgUpdate()
        }

        console.log("wsUserList =", wsUserList);
        console.log("wsMessageList =", wsMessageList);
    }

    webSocketConnect.sendMessage = sendMessage;
    webSocketConnect.wsGetUsers = wsGetUsers;
    webSocketConnect.wsGetChatMessages = wsGetChatMessages;
    webSocketConnect.sendActiveUserID = sendActiveUserID;

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
        Date(Date.now()) // this need fixing to shorter length
        );

        socket.send(newMessage);
            
        wsGetChatMessages(
            document.querySelector("#send-button").getAttribute("data-user-id"),
            document.querySelector("#send-button").getAttribute("data-target-id")
        )
        msgUpdate()

        document.getElementById("chat-text").textContent = "";
    }

    function wsGetUsers() {
        //JSON for getting users from db query
        let msg = {
        type: "wsGetUsers",
        };

        socket.send(JSON.stringify(msg));
    }
    
    function sendActiveUserID(usrID) {
        //JSON for getting users from db query
        let msg = {
        type: "sendUser",
        activeUser: usrID,
        };
        if (usrID) {
            socket.send(JSON.stringify(msg));
            console.log("Sent ActiveUserID over websocket to backend:", usrID);
        } else {
            console.log("Error: no ActiveUserID to send.");
        }
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
}
