import { useEffect, useState } from "react";

import Comment from "./Comment";

const CommentTree = ({ comments }) => {
  return (
    <div>
      {comments?.map((comment) => (
        <Comment key={comment.id} comment={comment} />
      ))}
    </div>
  );
};

export default CommentTree;
