package websockets

import (
    "fmt"
    "encoding/json"
)

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan Message
}

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
    }
}

func (pool *Pool) Start() {
    for {
        // var activeClientID string
        select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            fmt.Println("POOL: Size of Connection Pool: ", len(pool.Clients))
            fmt.Println("pool.Clients", pool.Clients)
            fmt.Println("connected client ID:", client.ID)
            // activeClientID = client.ID

            for user, _ := range pool.Clients {
            // for client := range pool.Clients {
                fmt.Println(client)
                user.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
                
                if user.ID == client.ID {
                    user.Conn.WriteJSON(Message{Type: 1, Body: `"newClient":` + string(client.ID)})
                }

            }
            break
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            fmt.Println("POOL: Size of Connection Pool: ", len(pool.Clients))
            for client, _ := range pool.Clients {
            // for client := range pool.Clients {
                client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
            }
            break
        case message := <-pool.Broadcast:
            fmt.Println("POOL: Sending message to all clients in Pool")
            fmt.Println("POOL: incoming message:", message)
            fmt.Println("POOL: incoming message.body:", message.Body)
            fmt.Println("POOL: incoming message.client:", message.Conn)

            byt := []byte(message.Body)
            fmt.Println("Trying to unmarshal")
            fmt.Println("message.Body", message.Body)
            fmt.Println("message.Body byt", byt)
            var dat map[string]interface{}
            if err := json.Unmarshal(byt, &dat); err != nil {
                panic(err)
            }
            fmt.Println("Unmarshaled data:", dat)

            // here we separate incoming messages by type 

            // if the frontend sends user ID for this ws conn
            if dat["type"] == "sendUser" {
                
                for client, _ := range pool.Clients {

                    // if recerived user Id conn is same as in Client struct, save user ID in Client
                    if client.Conn == message.Conn {
    
                        client.UserID = int(dat["activeUser"].(float64))
                        fmt.Println("Active user received and saved to client.UserID:", client.UserID)
                    }
                }
                

                // activeUserID := int(dat["activeUser"].(float64))
                // fmt.Println("Active user received:", activeUserID)
            }
            

            for client, _ := range pool.Clients {
            // for client := range pool.Clients {
                // if err := client.Conn.WriteJSON(message); err != nil {
                
                if err := client.Conn.WriteJSON("POOL: MIRRORED"); err != nil {
                    fmt.Println(err)
                    return
                }
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    fmt.Println("What up everybody?")
                    return
                }
            }
        }
    }
}