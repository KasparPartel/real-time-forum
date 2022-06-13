import { useState } from "react";

const Comment = () => {
  const [comment, setComment] = useState("");

  // getComment fetches comment with given id
  const getComment = () => {
    fetch(`http://localhost:4000/comment/${postId}/${id}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => setComment(data.comment));
  };

  return (
    <div>
      <p>{comment}</p>
    </div>
  );
};

export default Comment;
