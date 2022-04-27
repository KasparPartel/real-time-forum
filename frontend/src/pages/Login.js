import Cookies from "js-cookie"
import {Navigate} from 'react-router-dom'

function Login() {
    const sessionToken = Cookies.get("session_token")
    console.log(Cookies.get())

    if (sessionToken) {
        return <Navigate to='/'/>
    } else {
        return (
            <div className="registerBox">
                <header>Log in to our real-time-forum!</header>
                <form id="register" action="http://localhost:4000/v1/api/login/" method="POST">
                    <div>
                        <label>E-Mail: </label>
                        <input type="email" id="email" name="email" required/>
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
                            required
                        />
                    </div>

                    <div>
                        <label>Password: </label>
                        <input type="password" id="password" name="password" required/>
                    </div>

                    <input type="submit" value="Log in"/>
                </form>
            </div>
        );
    }
}


export default Login;
