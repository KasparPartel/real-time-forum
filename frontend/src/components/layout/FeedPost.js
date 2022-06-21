import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

// FeedPost is component for post previews in Feed
const FeedPost = ({ json }) => {
  const [comments, setComments] = useState([]);
  // const [likes, setLikes] = useState([]);
  // const [dislikes, setDislikes] = useState([]);

  useEffect(() => {
    getComments();
    // getLikes();
    // getDislikes();
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

  // const getLikes = () => {
  //   fetch(`http://localhost:4000/v1/api/like/${json.id}`, {
  //     method: "GET",
  //   })
  //     .then((res) => res.json())
  //     .then((data) => {
  //       setLikes(data);
  //     });
  // };

  // const getDislikes = () => {
  //   fetch(`http://localhost:4000/v1/api/dislike/${json.id}`, {
  //     method: "GET",
  //   })
  //     .then((res) => res.json())
  //     .then((data) => {
  //       setDislikes(data);
  //     });
  // };

  return (
    <div className="card">
      <h2 className="mb-1 text-lg font-bold">
        <Link to={`/post/${json.id}`}>{json.title}</Link>
      </h2>
      <p className="font-medium line-clamp-4">{json.body}</p>
      <p>
        <span className="comment-amount text-gray-500">
          <b>{comments?.length || 0}</b> comments
        </span>
        {/* <span className="like-amount text-gray-500">
          <b>{likes?.length || 0}</b> likes <span> | </span>
        </span>
        <span className="dislike-amount text-gray-500">
          <b>{dislikes?.length || 0}</b> dislikes
        </span> */}
      </p>
    </div>
  );
};

export default FeedPost;
