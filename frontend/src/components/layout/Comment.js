import { useEffect, useState } from "react";

import styles from "./Comment.module.css";
import globalStyles from "../../App.module.css";

const Comment = ({ comment }) => {
  const [user, setUser] = useState({});
  const [date, setDate] = useState(new Date());

  useEffect(() => {
    getUser();
    convertDateTime(comment.creation_time);
    console.log("created: ", comment.creation_time);
  }, []);

  const getUser = () => {
    fetch(`http://localhost:4000/v1/api/user/${comment.user_id}`)
      .then((res) => {
        if (!res.ok) {
          throw Error("Error fetching user");
        }
        return res.json();
      })
      .then((user) => setUser(user[0]))
      .catch((e) => console.log(e));
  };

  const convertDateTime = (dateTime) => {
    dateTime = dateTime.split(" ");

    const date = dateTime[0].split("-");
    const yyyy = date[0];
    const mm = date[1] - 1;
    const dd = date[2];

    const time = dateTime[1].split(":");
    const h = time[0];
    const m = time[1];
    const s = parseInt(time[2]); //get rid of that 00.0;

    setDate(new Date(yyyy, mm, dd, h, m, s));
  };

  return (
    <div className={`${globalStyles.flex_column} ${styles.comment}`}>
      <div className={`${globalStyles.flex_row} ${styles.meta}`}>
        <p className={styles.username}>{user.username}</p>
        <span className={styles.separator}>·</span>
        <p className={styles.dateTime}>
          <i>{date.toLocaleString([], { hour12: false })}</i>
        </p>
      </div>
      <p>{comment.body}</p>
    </div>
  );
};

export default Comment;
