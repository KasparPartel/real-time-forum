import {useContext} from "react";
import {UserContext} from "../UserContext";
import {useNavigate} from "react-router-dom";

function Profile() {
    const {user} = useContext(UserContext)
    const navigate = useNavigate()

    if (!user) {
        navigate("/login", {replace: true})
    }

    return (
        <div>
            <h1>This is your user profile.</h1>
            <ul>
                <li><b>Username:</b> {user?.username}</li>
                <li><b>E-mail:</b> {user?.email}</li>
                <li><b>First name:</b> {user?.first_name}</li>
                <li><b>Last name:</b> {user?.last_name}</li>
                <li><b>Gender:</b> {user?.gender}</li>
                <li><b>Age:</b> {user?.age}</li>
            </ul>
        </div>
    );
}

export default Profile;
