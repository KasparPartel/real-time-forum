//import

export function webSocketConnect(port) {

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection on port:", port);
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket on port:", port);
        // socket.send("Hello from the FrontEnd!")
        let msg = {
            type: "wsGetUsers",
            body: "get users query string here!",
            };
    
        socket.send(JSON.stringify(msg));
    
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
    }

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
        "saveChatMessage",
        document.querySelector("#chat-text").value,
        1,
        document.querySelector("#target-name").getAttribute("data-id"),
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
