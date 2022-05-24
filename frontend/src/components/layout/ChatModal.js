import React, { useState } from "react";
import classes from "./ChatModal.module.css";
// import { webSocketConnect } from "../../websocket.js"

function ChatModal(props) {
  
  // webSocketConnect("ws://localhost:4000/v1/api/ws");

  const [modal, setModal] = useState(false);

  const toggleModal = () => {
    setModal(!modal);
  };

  if (modal) {
    document.body.classList.add("active-modal");
  } else {
    document.body.classList.remove("active-modal");
  }

  return (
    <>
      <p className={classes.username} onClick={toggleModal}>
        {props.name}
      </p>
      {modal && (
        <div className={classes.chatmodal}>
          <div onClick={toggleModal} className={classes.overlay}></div>
          <div className={classes.chatmodalcontent}>
            <h2>You're chatting with: {props.name}</h2>
            <p>Here we will put the chat history with this user.</p>
            <div className={classes.chatfield}>
              {/* <label for="chat-text"></label> */}

              <textarea id="chat-text" name="chat-text" rows="4" cols="50" defaultValue="Enter your message here">
                
              </textarea>
              <br></br>
              <button>Submit</button>
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
