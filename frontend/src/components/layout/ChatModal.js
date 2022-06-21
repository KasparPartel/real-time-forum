import React, { useEffect, useState, useContext } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect, wsMessageList/* , getActiveUser */ } from "../../websocket.js"
import { UserContext } from "../../UserContext";

function ChatModal(props) {
  const {user} = useContext(UserContext)
  // webSocketConnect.wsGetChatMessages(user.id, props.id)

  // getActiveUser(user)

  const [render, setRender] = useState(false)
  console.log("render:", render);

  const toggleRender = () => {
    setRender(!render)
    // setRender(render + 1)
    webSocketConnect.wsGetChatMessages(user.id, props.id)
  }

  useEffect(() => {
    webSocketConnect.wsGetChatMessages(user.id, props.id)
    console.log("Asked for messages:", user.id, props.id);
  }, [user.id, props.id])

  
  const [messagelist, setMessagelist] = useState([])
  
  const [modal, setModal] = useState(false);
  const toggleModal = () => {
    toggleRender()
    setModal(!modal);
    // if (modal) {
    //   webSocketConnect.wsGetChatMessages(user.id, props.id)
    //   // setMessagelist(wsMessageList)
    // }
  };

  if (modal) {
    document.body.classList.add("active-modal");
  } else {
    document.body.classList.remove("active-modal");
  }

  console.log("Chatmodal var user is:", props.user.username);
  console.log("Chatmodal var user is:", user.username);
  console.log("Chatmodal var target is:", props.name);
  console.log("Chatmodal messagelist length is:", messagelist?.length);

  ChatModal.setMessagelist = setMessagelist;
  ChatModal.user = user;
  // ChatModal.setRender = setRender;

  let textList = messagelist

  return (
    <>
      {/* {render && toggleRender()} */}
      
      <p className={classes.username} onClick={toggleModal}>
        {props.name}
      </p>
      {(modal /* && messagelist */) && (
        <div className={classes.chatmodal}>
          <div onClick={toggleModal} className={classes.overlay}></div>
          <div className={classes.chatmodalcontent}>
            <h2>{props.user.username}, you're chatting with: {props.name}</h2>
              
              <div className={classes.chattexts}>
              {textList && (
                <div>
                {textList.map((message) => (
                  <ChatText key={message.id} body={message.body} userid={message.user_id} 
                  target={message.target} time={message.creation_time} loginuser={props.user.id}/>
                ))}
                </div>
              )}
              </div>

            <div className={classes.chatfield}>
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
      {render && toggleRender()}
    </>
  );
}

export function msgUpdate() {
  ChatModal.setMessagelist(wsMessageList)
}

export default ChatModal;
// export default ChatModal.user;

