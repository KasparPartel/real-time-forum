import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../UserContext";

import styles from "./Login.module.css";

export default function Login({ setCookie }) {
  const [formData, setFormData] = useState({});
  const { user } = useContext(UserContext);
  const navigate = useNavigate();

  if (user) {
    navigate("/", { replace: true });
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
        if (res.ok) {
          return res.json();
        } else if (res.status === 303) {
          // Must do redirect
          console.log("User already logged in!");
          return navigate("/", { replace: true });
        }
      })
      .then((data) => {
        console.log("Login successful!");
        setCookie("session_token", data.token, { path: "/" });
        return navigate("/", { replace: true });
      });
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.header}>Log in to our real-time-forum!</h2>
      <form className={styles.form} onSubmit={handleSubmit}>
        <div>
          <label className={styles.form__label} htmlFor="username">
            Username/Email:{" "}
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
            Password:{" "}
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
