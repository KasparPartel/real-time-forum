import React from "react";
import classes from "./ChatText.module.css";
// import ChatModal from "./ChatModal";
// import { wsUserList/* , webSocketConnect */ } from "../../websocket.js"

function ChatText(props) {


    return (
        <>
          <p className={classes.usermessage}>
            Example text (replaced with db history)
          </p>
        </>
    );

}

export default ChatText;