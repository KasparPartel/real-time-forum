function Login() {
  return (
    <div class="registerBox">
      <header>Log in to our real-time-forum!</header>
      <form id="register" action="forum/register" method="POST">
        <div>
          <label>E-Mail: </label>
          <input type="email" id="email" name="email" required />
        </div>

        <div>or</div>

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
          <label>Password: </label>
          <input type="password" id="password" name="password" required />
        </div>

        <input type="submit" value="Log in" />
      </form>
    </div>
  );
}

export default Login;
