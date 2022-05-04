//import

export function webSocketConnect(port) {

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection on port:", port);
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket on port:", port);
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

    webSocketConnect.sendMessage = sendMessage

    function sendMessage() {
        
        function composeMessage (/* Id, */ Body, User_id, Target_id, Creation_time) {
            let msg = {
                // id: String(Id),
                body: String(Body),
                user_id: String(User_id),
                target_id: String(Target_id),
                creation_time: String(Creation_time)
            }
            return JSON.stringify(msg)
        }
        
        let newMessage = composeMessage(
            // 1,
            document.querySelector("#chat-text").value,
            1,
            2,
            Date(Date.now())
        )
        
        socket.send(newMessage) 

        document.getElementById("chat-text").textContent = ""
    }
    
}




