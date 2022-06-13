import { useState } from "react";

const CommentTree = () => {
  const [comments, setComments] = useState({});

  return (
    <div>
      {comments.map((comment) => (
        <Comment key={comment.id} />
      ))}
    </div>
  );
};
