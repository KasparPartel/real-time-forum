import React, { useState, useContext } from "react";
import classes from "./ChatModal.module.css";
import ChatText from "./ChatText";
import { webSocketConnect } from "../../websocket.js";
import { UserContext } from "../../UserContext";

function ChatModal(props) {
  const { user } = useContext(UserContext);
  // gets messages to be rendered from Userlist.js
  let messages = props.messages[props.id]?.body;
  let basecount = 10;

  if (!props.messages[props.id]) {
    webSocketConnect.wsGetChatMessages(user.id, props.id, basecount);
  }

  // controls clicking on name in list of users (opens modal)
  const [modal, setModal] = useState(false);
  const toggleModal = () => {
    // sends modal opened (message viewed) to db
    webSocketConnect.sendModal(user.id, props.id);
    // gets current messages from db
    webSocketConnect.wsGetChatMessages(user.id, props.id, basecount);
    setModal(!modal);
  };

  // handles send message button click
  function sendClick() {
    webSocketConnect.sendMessage();
    if (props.messages[props.id].count) {
      webSocketConnect.wsGetChatMessages(
        user.id,
        props.id,
        props.messages[props.id]?.count
      );
    } else {
      webSocketConnect.wsGetChatMessages(user.id, props.id, basecount);
    }
    if (props.messages[user.id]?.count !== undefined) {
      webSocketConnect.wsGetChatMessages(
        props.id,
        user.id,
        props.messages[user.id]?.count
      );
    }
    webSocketConnect.wsGetUsers(user.id);
  }

  // opens and closes modal on component render
  if (modal) {
    document.body.classList.add("active-modal");
  } else {
    document.body.classList.remove("active-modal");
  }

  const onScroll = wsThrottle(() => {
    if (props.messages[props.id].count) {
      props.messages[props.id].count = props.messages[props.id].count + 10;
      webSocketConnect.wsGetChatMessages(
        user.id,
        props.id,
        props.messages[props.id]?.count
      );
    }
    webSocketConnect.wsGetUsers(user.id);
  }, 1000);

  function wsThrottle(func, wait) {
    let waiting = false;
    return function () {
      if (waiting) {
        return;
      }
      waiting = true;
      setTimeout(() => {
        func.apply(this, arguments);
        waiting = false;
      }, wait);
    };
  }

  return (
    <>
      <li
        className={`${props.active ? classes.active : classes.inactive} ${
          props.newmessage ? classes.unread : ""
        }`}
        onClick={toggleModal}
      >
        {props.name}
      </li>
      {modal && (
        <div className={classes.chatmodal}>
          <div onClick={toggleModal} className={classes.overlay}></div>
          <div className={classes.chatmodalcontent}>
            <h2>
              {props.user.username}, you're chatting with: {props.name}
            </h2>
            <div onScroll={onScroll} className={classes.chattexts}>
              <div>
                {messages?.map((message) => (
                  <ChatText
                    key={message.id}
                    body={message.body}
                    userid={message.user_id}
                    target={message.target}
                    time={message.creation_time}
                    loginuser={props.user.id}
                  />
                ))}
              </div>
            </div>
            <div className={classes.chatfield}>
              <textarea
                id="chat-text"
                name="chat-text"
                rows="4"
                cols="50"
                placeholder="Enter your message here"
              ></textarea>
              <br></br>
              <button
                id="send-button"
                onClick={sendClick}
                data-user-id={props.user.id}
                data-target-id={props.id}
              >
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
