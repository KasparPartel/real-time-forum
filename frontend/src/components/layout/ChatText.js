import React from "react";
import classes from "./ChatText.module.css";
// import ChatModal from "./ChatModal";
// import { wsUserList/* , webSocketConnect */ } from "../../websocket.js"

function ChatText(props) {

    //if (props.user === props.loginuser) {
        return (
            <>
              <p className={classes.usermessage}>
                {props.body}
              </p>
            </>
        );
    //}
    
    // if (props.user === props.target) {
    //     return (
    //         <>
    //           <p className={classes.targetmessage}>
    //             {props.body}
    //           </p>
    //         </>
    //     );
    // }

}

export default ChatText;