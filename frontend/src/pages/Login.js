import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../UserContext";

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
    <div className="registerBox">
      <header>Log in to our real-time-forum!</header>
      <form id="login" onSubmit={handleSubmit}>
        <div>
          <label>E-Mail: </label>
          <input
            type="email"
            id="email"
            name="email"
            onChange={handleChange}
            required
          />
        </div>

        <div>or</div>

        <div>
          <label>Username: </label>
          <input
            type="text"
            id="username"
            name="username"
            pattern="^[a-zA-Z0-9]+$"
            title="Alphanumerical characters only."
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label>Password: </label>
          <input
            type="password"
            id="password"
            name="password"
            onChange={handleChange}
            required
          />
        </div>

        <input type="submit" value="Log in" />
      </form>
    </div>
  );
}
