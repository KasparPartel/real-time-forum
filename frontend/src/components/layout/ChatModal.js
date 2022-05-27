import React, { useEffect, useState } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect, wsMessageList } from "../../websocket.js"

function ChatModal(props) {

  useEffect(() => {
    webSocketConnect.wsGetChatMessages(props.user.id, props.id)
  }, [props.user.id, props.id])

  const [messagelist, setMessagelist] = useState(wsMessageList)
  const [modal, setModal] = useState(false);
  const toggleModal = () => {
    setModal(!modal);
    if (modal) {
      webSocketConnect.wsGetChatMessages(props.user.id, props.id)
      setMessagelist(wsMessageList)
    }
  };

  if (modal) {
    document.body.classList.add("active-modal");
  } else {
    document.body.classList.remove("active-modal");
  }

  console.log("Chatmodal var user is:", props.user.username);

  ChatModal.setMessagelist = setMessagelist;

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

            
              {messagelist.map((message) => (
                <ChatText key={message.id} body={message.body} userid={message.user_id} 
                target={message.target} time={message.creation_time} loginuser={props.user.id}/>
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

export function msgUpdate() {
  ChatModal.setMessagelist(wsMessageList)
}

export default ChatModal;
