//import

export function webSocketConnect(port) {

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection...");
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket");
        // socket.send("Hello from the FrontEnd!")
        let testMessage = composeMessage(
            // 1,
            "Test message body",
            1,
            2,
            Date(Date.now())
        )

        sendMessage(testMessage)

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
    
    function sendMessage(string) {
        socket.send(string)
    }
}



