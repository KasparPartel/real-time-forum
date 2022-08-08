import { useContext, useState } from "react";
import { UserContext } from "../UserContext";
import { useNavigate } from "react-router-dom";

import styles from "./Register.module.css";

function Register() {
  const [formData, setFormData] = useState({ gender: "male" });
  const [errorMsg, setErrorMsg] = useState("");

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

    fetch("http://localhost:4000/v1/api/user/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(formData),
    })
      .then((res) => {
        if (!res.ok) {
          setErrorMsg("Cannot register user!");
          throw Error(res.statusText);
        }
        return navigate("/login", { replace: true });
      })
      .catch((e) => {
        console.log(e);
      });
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.header}>Register new user</h2>
      {errorMsg && <p className={styles.error}>{errorMsg}</p>}
      <form className={styles.form} onSubmit={handleSubmit}>
        <div>
          <label className={styles.form__label} htmlFor="email">
            E-Mail:{" "}
          </label>
          <br />
          <input
            type="email"
            className={styles.form__input}
            name="email"
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className={styles.form__label} htmlFor="username">
            Username:{" "}
          </label>
          <br />
          <input
            type="text"
            className={styles.form__input}
            name="username"
            pattern="^[a-zA-Z0-9]+$"
            minLength="5"
            title="Alphanumerical characters only."
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className={styles.form__label} htmlFor="age">
            Age:{" "}
          </label>
          <br />
          <input
            type="number"
            pattern="^[0-9]+$"
            min="1"
            max="120"
            className={styles.form__input}
            name="age"
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className={styles.form__label} htmlFor="gender">
            Gender:{" "}
          </label>
          <br />
          <select
            name="gender"
            className={styles.form__select}
            onChange={handleChange}
          >
            <option value="male">Male</option>
            <option value="female">Female</option>
            {/*<option defaultValue="unknown" /* selected="selected">
              Prefer not to specify
            </option> */}
          </select>
        </div>

        <div>
          <label className={styles.form__label} htmlFor="first_name">
            First name:{" "}
          </label>
          <br />
          <input
            type="text"
            className={styles.form__input}
            name="first_name"
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className={styles.form__label} htmlFor="last_name">
            Last name:{" "}
          </label>
          <br />
          <input
            type="text"
            className={styles.form__input}
            name="last_name"
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
            minLength="8"
            onChange={handleChange}
            required
          />
        </div>

        {/* <div>
          <label>Re-enter password: </label>
          <input type="password" id="password2" name="password2" required />
        </div> */}

        <input className={styles.btn_submit} type="submit" value="Sign up" />
      </form>
    </div>
  );
}

export default Register;
