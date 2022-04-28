//import

export function webSocketConnect(port) {

    //let socket = new WebSocket("ws://localhost:4000/ws")
    let socket = new WebSocket(port)
    console.log("Attempting WebSocket Connection...");
    
    socket.onopen = () => {
        console.log("Successfully Connected to Websocket");
        socket.send("Hello from the FrontEnd!")
    
    }
    
    socket.onclose = (e) => {
        console.log("WebSocket Connection Closed: ", e);
    }
    
    socket.onerror = (err) => {
        console.log("WebSocket Error: ", err);
    }

    socket.onmessage = (msg) => {
        console.log("Backend has responded: ", msg);
    }

}
