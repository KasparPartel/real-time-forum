import { Fragment, useContext, useState } from "react";

import styles from "./CreateComment.module.css";

const CreateComment = ({ postID, userID, getComments }) => {
  const [commentForm, setCommentForm] = useState("");
  const [formData, setFormData] = useState({
    user_id: userID.toString(),
    post_id: postID.toString(),
  });

  // handleChange sets new object based on changed form data
  const handleChange = (e) => {
    setCommentForm(e.target.value);

    let formDataCopy = Object.assign({}, formData);
    let name = e.target.getAttribute("name");

    formDataCopy[name] = e.target.value;
    setFormData(formDataCopy);
  };

  // handleSubmit sends data to api and navigates to new post page
  const handleSubmit = (e) => {
    e.preventDefault();

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

    setCommentForm("");
  };

  return (
    <Fragment>
      <form onSubmit={handleSubmit} className={styles.form}>
        <input
          type="text"
          name="body"
          onChange={handleChange}
          value={commentForm}
        />
        <input type="submit" value="Comment" />
      </form>
    </Fragment>
  );
};

export default CreateComment;
