//import

export let wsUserList
export let wsMessageList

export function webSocketConnect(port) {

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection on port:", port);
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket on port:", port);
        wsGetUsers()
    
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
        // console.log("Backend has responded with data.body: ", msg.data.body);
        let incomingJson = JSON.parse(msg.data)
        // console.log(incomingJson);
        // console.log(incomingJson.body[1]);
        if (incomingJson.type === "wsReturnedUsers") {
            wsUserList = incomingJson.body
        }
        if (incomingJson.type === "wsReturnedMessages") {
            wsMessageList = incomingJson.body
        }
        console.log("wsUserList =", wsUserList);
        console.log("wsMessageList =", wsMessageList);
    }

    // function wsUserList(data) {
    //     return JSON.parse(data.body)
    // }

    webSocketConnect.sendMessage = sendMessage;
    webSocketConnect.wsGetUsers = wsGetUsers;

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

        document.getElementById("chat-text").textContent = "";
    }

    function wsGetUsers() {
        //JSON for getting users from db query
        let msg = {
        type: "wsGetUsers",
        body: "get users query string here!",
        };

        socket.send(JSON.stringify(msg));
    }

}
