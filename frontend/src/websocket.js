import { msgUpdate } from "./components/layout/ChatModal";
import { usrUpdate } from "./components/layout/Userlist";

export let wsUserList = []
export let wsMessageList = []

export function webSocketConnect(port) {

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection on port:", port);
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket on port:", port);
        wsGetUsers()
        usrUpdate()
        // wsGetChatMessages()
    }
    
    socket.onclose = (e) => {
        console.log("WebSocket Connection Closed: ", e);
    }
    
    socket.onerror = (err) => {
        console.log("WebSocket Error: ", err);
    }

    socket.onmessage = (msg) => {
        console.log("Backend has responded: ", msg);
        console.log("Backend has responded with data: ", msg.data);
        let incomingJson = JSON.parse(msg.data)

        if (incomingJson.type === "wsReturnedUsers") {
            wsUserList = incomingJson.body
            usrUpdate()
        }
        if (incomingJson.type === "wsReturnedMessages") {
            wsMessageList = incomingJson.body
            msgUpdate()
        }

        console.log("wsUserList =", wsUserList);
        console.log("wsMessageList =", wsMessageList);
    }

    webSocketConnect.sendMessage = sendMessage;
    webSocketConnect.wsGetUsers = wsGetUsers;
    webSocketConnect.wsGetChatMessages = wsGetChatMessages;

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

    function wsGetUsers() {
        //JSON for getting users from db query
        let msg = {
        type: "wsGetUsers",
        };

        socket.send(JSON.stringify(msg));
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
