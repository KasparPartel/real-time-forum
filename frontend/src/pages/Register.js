import {useContext, useState} from "react";
import {UserContext} from "../UserContext";
import {useNavigate} from "react-router-dom";

function Register() {
  const [formData, setFormData] = useState({gender: "male"})
  const {user} = useContext(UserContext)
  const navigate = useNavigate()

  if (user) {
    navigate("/", {replace: true})
  }

  const handleChange = (e) => {
    let formDataCopy = Object.assign({}, formData)
    let name = e.target.getAttribute("name")

    formDataCopy[name] = e.target.value
    setFormData(formDataCopy)
}

  const handleSubmit = (e) => {
    e.preventDefault()

    console.log(formData)

    fetch("http://localhost:4000/v1/api/user/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      body: JSON.stringify(formData)
    })
    .then(res => {
      if (res.ok) {
        return navigate("/login", {replace: true})
      } else {
        console.log("Cannot register user!");
        return
      }
    })
  }

  return (
    <div className="registerBox">
      <header>Register to our real-time-forum!</header>
      <form id="register" onSubmit={handleSubmit}>
        <div>
          <label>E-Mail: </label>
          <input type="email" id="email" name="email" onChange={handleChange} required />
        </div>

        <div>
          <label>Username: </label>
          <input
            type="text"
            id="username"
            name="username"
            pattern="^[a-zA-Z0-9]+$"
            minLength="5"
            title="Alphanumerical characters only."
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label>Age: </label>
          <input type="number" pattern="^[0-9]+$" min="1" max="120" id="age" name="age" onChange={handleChange} required />
        </div>

        <div>
          <label>Gender: </label>
          <select name="gender" id="gender" onChange={handleChange}>
            <option value="male">Male</option>
            <option value="female">Female</option>
            {/*<option defaultValue="unknown" /* selected="selected">
              Prefer not to specify
            </option> */}
          </select>
        </div>

        <div>
          <label>First name: </label>
          <input type="text" id="firstname" name="first_name" onChange={handleChange} required />
        </div>

        <div>
          <label>Last name: </label>
          <input type="text" id="lastname" name="last_name" onChange={handleChange} required />
        </div>

        <div>
          <label>Password: </label>
          <input type="password" id="password" name="password" minLength="5" onChange={handleChange} required />
        </div>

        {/* <div>
          <label>Re-enter password: </label>
          <input type="password" id="password2" name="password2" required />
        </div> */}

        <input type="submit" value="Sign up" />
      </form>
    </div>
  );
}

export default Register;
