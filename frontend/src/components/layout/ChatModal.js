import React, { useEffect, useState } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect, wsMessageList } from "../../websocket.js"

// import {useContext} from "react";
// import {UserContext} from "../../UserContext";

let messages = []

export function changeMessages() {
  messages = wsMessageList
}

// export function forceUpdate() {
//   useForceUpdate()
// }

function useForceUpdate(){
  const [value, setValue] = useState(0); // integer state
  return () => setValue(value => value + 1); // update the state to force render
}

function ChatModal(props) {
  
  const forceUpdate = useForceUpdate()

  useEffect(() => {
    webSocketConnect.wsGetChatMessages(props.user.id, props.id)
  }, [props.user.id, props.id])

  // useEffect(() => {
  //   useForceUpdate()
  // }, [messages])

  // const [messages, setMessages] = useState([])
  
  // useEffect(() => {
    //   retrieveMessages(props.user.id, props.id)
  // }, [props.user.id, props.id])
  
  // async function retrieveMessages(usr, trgt) {
    //   try {

  //   } catch {
    //     console.log("Error retrieving messages")
    //   }
    // }

  // const [messages, setMessages] = useState([])
  
  // useEffect(() => {
  //   webSocketConnect.wsGetChatMessages(props.user.id, props.id)
  //   setMessages(wsMessageList)
  // }, [props.user.id, props.id])


  const [modal, setModal] = useState(false);
  const toggleModal = () => {
    setModal(!modal);
    if (modal) {
      webSocketConnect.wsGetChatMessages(props.user.id, props.id)
      // setMessages(wsMessageList)
      changeMessages()
    }
    forceUpdate()
  };

  useEffect(() => {

  }, [modal])

  if (modal) {
    document.body.classList.add("active-modal");
  } else {
    document.body.classList.remove("active-modal");
  }

  console.log("Chatmodal var user is:", props.user.username);

  return (
    <>
      <p className={classes.username} onClick={toggleModal}>
        {props.name}
      </p>
      {modal && (
        <div className={classes.chatmodal}>
          <div onClick={toggleModal} className={classes.overlay}></div>
          <div className={classes.chatmodalcontent}>
            <h2>{props.user.username}, you're chatting with: {props.name}</h2>
            <p>Here we will put the chat history with this user.</p>

            
              {messages.map((message) => (
                <ChatText key={message.id} body={message.body} user={message.user_id} 
                target={message.target} time={message.creation_time} loginuser={props.user}/>
              ))}
            

            <div className={classes.chatfield}>
              {/* <label for="chat-text"></label> */}

              <textarea id="chat-text" name="chat-text" rows="4" cols="50" placeholder="Enter your message here">
                
              </textarea>
              <br></br>
              <button id="send-button" onClick={webSocketConnect.sendMessage} 
                data-user-id={props.user.id} data-target-id={props.id}>
                Submit
                </button>
              <button className={classes.closemodal} onClick={toggleModal}>
                CLOSE
              </button>
            </div>
          </div>
        </div>
      )}
    </>

  );
}

export default ChatModal;
