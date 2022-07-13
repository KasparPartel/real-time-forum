import React, { /* useEffect, */ useState, useContext } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect, /* wsMessageList, */ wsConnected } from "../../websocket.js"
import { UserContext } from "../../UserContext";

// let orderMsg = false
function ChatModal(props) {
  const {user} = useContext(UserContext)
  // webSocketConnect.wsGetChatMessages(user.id, props.id)
  
  // let msgBubble = "nonew"

  // getActiveUser(user)

  const [render, setRender] = useState(false)
  // console.log("render:", render);

  const toggleRender = () => {
    if (wsConnected) {
      setRender(!render)
      // setRender(render + 1)
      // orderMsg = true
      webSocketConnect.sendModal(user.id, props.id, messagelist.length)
      webSocketConnect.wsGetChatMessages(user.id, props.id)
      // webSocketConnect.wsGetUsers(user.id)
    }
  }

  // useEffect(() => {
  //   webSocketConnect.wsGetChatMessages(user.id, props.id); 
  // }, [user.id, props.id])

  // useEffect(() => {
  //   if (wsConnected) {
  //     orderMsg = true
  //     webSocketConnect.wsGetChatMessages(user.id, props.id);
  //     // webSocketConnect.sendActiveUserID(user.id); // fires too often, but works
  //     if (user.id > 0) {
  //       webSocketConnect.wsGetUsers()
  //     }
  //     console.log("Asked for messages:", user.id, props.id);
  //   }
  // }, [user.id, props.id])

  
  const [messagelist, setMessagelist] = useState([])
  
  const [modal, setModal] = useState(false);
  const toggleModal = () => {
    toggleRender()
    setModal(!modal);
    if (modal) {
      webSocketConnect.sendModal(user.id, props.id)
    }
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

  // console.log("Chatmodal var user is:", props.user.username);
  // console.log("Chatmodal var user is:", user.username);
  // console.log("Chatmodal var target is:", props.name);
  // console.log("Chatmodal messagelist length is:", messagelist?.length);

  // let msgLength = messagelist.length

  ChatModal.setMessagelist = setMessagelist;
  ChatModal.user = user;
  // ChatModal.orderMsg = orderMsg;
  // ChatModal.msgBubble = msgBubble;
  ChatModal.msgLength = messagelist?.length

  console.log("ChatModal.msgLength", ChatModal.msgLength);
  // ChatModal.setRender = setRender;

  let textList = messagelist


  return (
    <>
      {/* {render && toggleRender()} */}
      
      {/* <li className={[classes[props.class], msgBubble].join(' ')} onClick={toggleModal}> */}
      <li className={classes[props.class]} onClick={toggleModal}>
        {props.name}
      </li>
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

export function msgUpdate(messages) {
  ChatModal.setMessagelist(messages)
  // if (ChatModal.orderMsg === false) {
  //   ChatModal.msgBubble = "newstuff" //classes.newmsg
  // }
  // ChatModal.orderMsg = false
  // ChatModal.setMessagelist(wsMessageList)
}

export default ChatModal;
// export default ChatModal.user;

