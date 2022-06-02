import { useContext, useState } from "react";
import { UserContext } from "../UserContext";
import { useNavigate } from "react-router-dom";

function CreatePost() {
  const { user } = useContext(UserContext);
  const navigate = useNavigate();

  if (!user) {
    navigate("/login", { replace: true });
  }

  const [formData, setFormData] = useState({
    user_id: user.id.toString(),
    filename: "",
  });

  const handleChange = (e) => {
    let formDataCopy = Object.assign({}, formData);
    let name = e.target.getAttribute("name");

    formDataCopy[name] = e.target.value;
    setFormData(formDataCopy);
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    console.log(formData);

    fetch("http://localhost:4000/v1/api/post/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else {
          console.log("Cannot add post!");
          return;
        }
      })
      .then((data) => navigate(`/post/${data}`, { replace: true }));
  };

  return (
    <div className="create-post">
      <h2>Create a new post here:</h2>

      <form id="create-post" onSubmit={handleSubmit}>
        <label>Post title: </label>
        <input
          type="text"
          id="post-title"
          name="title"
          onChange={handleChange}
          required
        />
        <br></br>
        {/* <label>Post topic: </label>
            <select name="post-topic" id="post-topic">
                <option value="Memes">Memes</option>
                <option value="Useful">Useful</option>
                <option defaultValue="Random">Random</option>
            </select>
            <br></br> */}
        <textarea
          id="post-body"
          name="body"
          rows="4"
          cols="50"
          placeholder="Enter your post here"
          onChange={handleChange}
          required
        ></textarea>
        <br></br>

        <input type="submit" value="Create Post" />
      </form>
    </div>
  );
}

export default CreatePost;
