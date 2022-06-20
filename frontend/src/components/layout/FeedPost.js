import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

// FeedPost is component for post previews in Feed
const FeedPost = ({ json }) => {
  const [comments, setComments] = useState([]);

  useEffect(() => {
    getComments();
  }, []);

  const getComments = () => {
    fetch(`http://localhost:4000/v1/api/comments/${json.id}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setComments(data);
        console.log("comments", comments);
      });
  };

  return (
    <div className="card">
      <h2 className="mb-1 text-lg font-bold">
        <Link to={`/post/${json.id}`}>{json.title}</Link>
      </h2>
      <p className="font-medium line-clamp-4">{json.body}</p>
      <p>
        <span className="comment-amount text-gray-500">
          <b>{comments ? comments.length : 0}</b> comments <span> | </span>
        </span>
        <span className="like-amount text-gray-500">
          <b>{json.likeAmount ? json.likeAmount : 0}</b> likes <span> | </span>
        </span>
        <span className="dislike-amount text-gray-500">
          <b>{json.dislikeAmount ? json.dislikeAmount : 0}</b> dislikes
        </span>
      </p>
    </div>
  );
};

export default FeedPost;
