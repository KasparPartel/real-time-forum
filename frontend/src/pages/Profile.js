import { useContext } from "react";
import { UserContext } from "../UserContext";
import { useNavigate } from "react-router-dom";

import styles from "./Profile.module.css";
import globalStyles from "../App.module.css";

export default function Profile() {
  const { user } = useContext(UserContext);
  const navigate = useNavigate();

  if (!user) {
    navigate("/login", { replace: true });
  }

  return (
    <div>
      <h2 className={styles.header}>This is your user profile</h2>
      <ul className={`${globalStyles.flex_column} ${styles.ul}`}>
        <li className={styles.ul__li}>
          <b>Username:</b> {user?.username}
        </li>
        <li className={styles.ul__li}>
          <b>E-mail:</b> {user?.email}
        </li>
        <li className={styles.ul__li}>
          <b>First name:</b> {user?.first_name}
        </li>
        <li className={styles.ul__li}>
          <b>Last name:</b> {user?.last_name}
        </li>
        <li className={styles.ul__li}>
          <b>Gender:</b> {user?.gender}
        </li>
        <li className={styles.ul__li}>
          <b>Age:</b> {user?.age}
        </li>
      </ul>
    </div>
  );
}
