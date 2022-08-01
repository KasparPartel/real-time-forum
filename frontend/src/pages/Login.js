import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../UserContext";
import { webSocketConnect, wsConnected } from "../websocket.js"

import styles from "./Login.module.css";

export default function Login({ setCookie }) {
  const { user } = useContext(UserContext);
  const navigate = useNavigate();

  const [formData, setFormData] = useState({});
  const [errorMsg, setErrorMsg] = useState("");

  if (user) {
    navigate("/", { replace: true });
    if (wsConnected) {

      webSocketConnect.sendActiveUserID(user.id);
    }
  }

  const handleChange = (e) => {
    let formDataCopy = Object.assign({}, formData);
    let name = e.target.getAttribute("name");

    formDataCopy[name] = e.target.value;
    setFormData(formDataCopy);
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    fetch("http://localhost:4000/v1/api/login/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(formData),
    })
      .then((res) => {
        if (!res.ok) {
          setErrorMsg("Error logging in!");
          throw Error(res.statusText);
        }
        return res.json();
      })
      .then((data) => {
        console.log("Login successful!");
        setCookie("session_token", data.token, { path: "/" });
        return navigate("/", { replace: true });
      })
      .catch((e) => {
        console.log(e);
      });
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.header}>Log in to our real-time-forum!</h2>
      {errorMsg && <p className={styles.error}>{errorMsg}</p>}
      <form className={styles.form} onSubmit={handleSubmit}>
        <div>
          <label className={styles.form__label} htmlFor="username">
            Username/Email:
          </label>
          <br />
          <input
            type="text"
            className={styles.form__input}
            name="username"
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className={styles.form__label} htmlFor="password">
            Password:
          </label>
          <br />
          <input
            type="password"
            className={styles.form__input}
            name="password"
            onChange={handleChange}
            required
          />
        </div>

        <input className={styles.btn_submit} type="submit" value="Log in" />
      </form>
    </div>
  );
}
