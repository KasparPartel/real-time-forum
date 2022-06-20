const Comment = ({ json }) => {
  return (
    <div>
      <p>
        {json.body}
        <br />
        <span>
          <i>{json.creation_time}</i>
        </span>
        <br />
        <span>User: {json.user_id}</span>
      </p>
    </div>
  );
};

export default Comment;
