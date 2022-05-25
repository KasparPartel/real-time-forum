import React, { useState } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect } from "../../websocket.js"

// import {useContext} from "react";
// import {UserContext} from "../../UserContext";

function ChatModal(props) {
  
  //const [user, setUser] = useState({})

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

  console.log("Chatmodal var user is:", props.user);

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
            <ChatText /* key={target.id} id={target.id} name={target.username} user={user} *//>
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
