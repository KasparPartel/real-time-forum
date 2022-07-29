import React, { useEffect, useState, useContext } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect/* , wsMessageList */} from "../../websocket.js"
import { UserContext } from "../../UserContext";

// let orderMsg = false
function ChatModal(props) {

  
  const {user} = useContext(UserContext)
  console.log("Rendering ChatModal!", user.id, props.target);

  // console.log("wsMessageList in Chatmodal:", wsMessageList);

  
  // target int = props.id
  
  const [messagelist, setMessagelist] = useState([])
  // const [render, setRender] = useState(false)
  
  // let wsMessages = []
  // for (const [key, value] of Object.entries(wsMessageList)) {
  //   if (key === props.id.toString()) {
  //     wsMessages = [...value]
  //   }
  // }
  
  console.log("SETTING MESSAGES BASE:", props.messages);
  // console.log("SETTING MESSAGES BASE:", wsMessageList);
  // console.log("SETTING MESSAGES id:", props.id);
  console.log("SETTING MESSAGES:", props.messages[4]);
  useEffect(() => {
    setMessagelist(props.messages[4])
  }, [props.id, props.messages])

  let messages = props.messages[props.id]
  console.log("messages!!!", messages);

  // function assignMessages() {
  //   console.log("Looping!");
  //   for (const [key, value] of Object.entries(wsMessageList)) {
  //     if (key === props.id.toString()) {
  //       console.log("SETTING MESSAGES");
  //       setMessagelist(value)
  //     }
  //   }
  // }
  
  
  const [modal, setModal] = useState(false);
  const toggleModal = () => {
    // toggleRender()
    console.log("toggleModal:", modal);
    //if (modal === false) {
      webSocketConnect.sendModal(user.id, props.id)
      console.log("modal sent");
      // webSocketConnect.removeUnread([props.name])
    //}
    getMessages()
    setModal(!modal);
  };
  
  function getMessages() {
    webSocketConnect.wsGetChatMessages(user.id, props.id)
    webSocketConnect.wsGetChatMessages(props.id, user.id)
  }

  function sendClick() {
    webSocketConnect.sendMessage()
    getMessages()
  }

  if (modal) {
    document.body.classList.add("active-modal");
    // webSocketConnect.sendModal(user.id, props.id);
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
  // ChatModal.render = render;
  // ChatModal.setRender = setRender;
  // ChatModal.orderMsg = orderMsg;
  // ChatModal.msgBubble = msgBubble;
  ChatModal.msgLength = messagelist?.length
  
  // ChatModal.toggleRender = toggleRender;
  // function toggleRender() {
  //   // setRender(!render)
  //   console.log("toggleRender fired!");
  // }

  console.log("ChatModal.msgLength, user.id, target.id", ChatModal.msgLength, user.id, props.id);


  return (
    <>
      {/* {render && toggleRender()} */}
      
      {/* <li className={[classes[props.class], msgBubble].join(' ')} onClick={toggleModal}> */}
      <li className={`${props.active ? classes.active : ""} ${props.newmessage ? classes.unread : ""}`/* classes[props.class] */} onClick={toggleModal}>
        {props.name}
      </li>
      {(modal /* && messagelist */) && (
        <div className={classes.chatmodal}>
          <div onClick={toggleModal} className={classes.overlay}></div>
          <div className={classes.chatmodalcontent}>
            <h2>{props.user.username}, you're chatting with: {props.name}</h2>
              
              <div className={classes.chattexts}>
              
                <div>
                {messages?.map((message) => (
                  <ChatText key={message.id} body={message.body} userid={message.user_id} 
                  target={message.target} time={message.creation_time} loginuser={props.user.id}/>
                ))}
                </div>
              
              </div>

            <div className={classes.chatfield}>
              <textarea id="chat-text" name="chat-text" rows="4" cols="50" placeholder="Enter your message here">
              </textarea>
              <br></br>
              <button id="send-button" onClick={sendClick} 
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
      {/* {render && toggleRender()} */}
    </>
  );
}

export function msgUpdate(messages) {
  // ChatModal.setMessagelist(messages)
  // ChatModal.setRender(!ChatModal.render)
  // if (ChatModal.orderMsg === false) {
  //   ChatModal.msgBubble = "newstuff" //classes.newmsg
  // }
  // ChatModal.orderMsg = false
  if (messages) {
    ChatModal.setMessagelist([...messages])
  }
  // ChatModal.toggleRender()
  // ChatModal.setMessagelist(wsMessageList)
}

export default ChatModal;
// export default ChatModal.user;

