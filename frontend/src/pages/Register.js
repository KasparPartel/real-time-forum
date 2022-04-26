function Register() {
  return (
    <div className="registerBox">
      <header>Register to our real-time-forum!</header>
      <form id="register" action="forum/register" method="POST">
        <div>
          <label>E-Mail: </label>
          <input type="email" id="email" name="email" required />
        </div>

        <div>
          <label>Nickname: </label>
          <input
            type="text"
            id="username"
            name="username"
            pattern="^[a-zA-Z0-9]+$"
            title="Alphanumerical characters only."
            required
          />
        </div>

        <div>
          <label>Age: </label>
          <input type="number" min="1" max="120" id="age" name="age" required />
        </div>

        <div>
          <label>Gender: </label>
          <select name="gender" id="gender">
            <option value="male">Male</option>
            <option value="female">Female</option>
            <option defaultValue="unknown" /* selected="selected" */>
              Prefer not to specify
            </option>
          </select>
        </div>

        <div>
          <label>First name: </label>
          <input type="text" id="firstname" name="firstname" required />
        </div>

        <div>
          <label>Last name: </label>
          <input type="text" id="lastname" name="lastname" required />
        </div>

        <div>
          <label>Password: </label>
          <input type="password" id="password" name="password" required />
        </div>

        <div>
          <label>Re-enter password: </label>
          <input type="password" id="password2" name="password2" required />
        </div>

        <input type="submit" value="Sign up" />
      </form>
    </div>
  );
}

export default Register;
