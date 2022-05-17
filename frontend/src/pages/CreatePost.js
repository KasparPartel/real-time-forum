import {useContext} from "react";
import {UserContext} from "../UserContext";
import {useNavigate} from "react-router-dom";

function CreatePost() {
    const {user} = useContext(UserContext)
    const navigate = useNavigate()

    if (!user) {
        navigate("/login", {replace: true})
    }

    return (
        <div className="create-post">
            <h1>Create a new post here:</h1>

            <label>Post title: </label>
            <input type="text" id="firstname" name="firstname" required/>
            <br></br>
            <label>Post topic: </label>
            <select name="post-topic" id="post-topic">
                <option value="Memes">Memes</option>
                <option value="Useful">Useful</option>
                <option defaultValue="Random">Random</option>
            </select>
            <br></br>
            <textarea id="post-body" name="post-body" rows="4" cols="50">
        Enter your post here
      </textarea>
            <br></br>
            <button>Submit</button>
        </div>
    );
}

export default CreatePost;
