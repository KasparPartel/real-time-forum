import { Fragment, useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../../UserContext";
import CommentTree from "./CommentTree";

const CreateComment = ({ postID, userID, getComments }) => {
  const [formData, setFormData] = useState({
    user_id: userID.toString(),
    post_id: postID.toString(),
  });

  // handleChange sets new object based on changed form data
  const handleChange = (e) => {
    let formDataCopy = Object.assign({}, formData);
    let name = e.target.getAttribute("name");

    formDataCopy[name] = e.target.value;
    setFormData(formDataCopy);
  };

  // handleSubmit sends data to api and navigates to new post page
  const handleSubmit = (e) => {
    e.preventDefault();

    console.log(formData);

    fetch("http://localhost:4000/v1/api/comments/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    }).then((res) => {
      if (res.status === 201) {
        getComments();
        return;
      } else {
        console.log("Cannot add comment!");
        return;
      }
    });
  };

  return (
    <Fragment>
      <form onSubmit={handleSubmit}>
        <input type="text" name="body" onChange={handleChange} />
        <input type="submit" value="Comment" />
      </form>
    </Fragment>
  );
};

export default CreateComment;
