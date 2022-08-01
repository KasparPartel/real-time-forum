import React from "react";
import classes from "./ChatText.module.css";

function ChatText(props) {

  if (props.userid === String(props.loginuser)) {
      return (
          <>
            <li className={classes.usermessage}>
              {props.body} <br>
              </br>
              @ {new Date(Date.parse(props.time)).toLocaleString('en-GB')}
            </li>
          </>
      );
  } else {
    return (
      <>
        <li className={classes.targetmessage}>
          {props.body} <br>
            </br>
            @ {new Date(Date.parse(props.time)).toLocaleString('en-GB')}
        </li>
      </>
  );
  }

}

export default ChatText;