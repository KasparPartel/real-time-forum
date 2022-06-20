const Comment = ({ json }) => {
  return (
    <div>
      <p>{json.body}</p>
      <p>{json.author}</p>
    </div>
  );
};

export default Comment;
