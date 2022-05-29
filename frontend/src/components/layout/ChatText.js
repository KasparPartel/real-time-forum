import React from "react";
import classes from "./ChatText.module.css";

function ChatText(props) {

  if (props.userid === String(props.loginuser)) {
      return (
          <>
            <p className={classes.usermessage}>
              {props.body}
            </p>
          </>
      );
  } else {
    return (
      <>
        <p className={classes.targetmessage}>
          {props.body}
        </p>
      </>
  );
  }

}

export default ChatText;