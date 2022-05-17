import {useContext} from "react";
import {UserContext} from "../UserContext";
import {useNavigate} from "react-router-dom";

function Profile() {
    const {user, setUser} = useContext(UserContext)
    const navigate = useNavigate()

    if (!user) {
        navigate("/login", {replace: true})
    }

    return (
        <div>
            <h1>This is your user profile.</h1>
            <ul>
                <li>Username: {user?.username}</li>
                <li>E-mail: {user?.email}</li>
                <li>First name: {user?.first_name}</li>
                <li>Last name: {user?.last_name}</li>
                <li>Gender: {user?.gender}</li>
            </ul>
        </div>
    );
}

export default Profile;
