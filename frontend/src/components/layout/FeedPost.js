import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import styles from "./FeedPost.module.css";
import globalStyles from "../../App.module.css";

// FeedPost is component for post previews in Feed
const FeedPost = ({ json }) => {
  const [comments, setComments] = useState([]);
  const [category, setCategory] = useState({});
  const [author, setAuthor] = useState({});

  // const [likes, setLikes] = useState([]);
  // const [dislikes, setDislikes] = useState([]);

  useEffect(
    () => {
      getComments();
      getCategory();
      getAuthor();
      console.log("styles", styles);
      // getLikes();
      // getDislikes();
    },
    // eslint-disable-next-line
    []
  );

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

  const getCategory = () => {
    fetch(`http://localhost:4000/v1/api/categories/${json["category_id"]}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setCategory(data[0]);
      });
  };

  const getAuthor = () => {
    fetch(`http://localhost:4000/v1/api/user/${json["user_id"]}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setAuthor(data[0]);
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
    <div className={styles.div__post}>
      <div className={`${globalStyles.flex_row} ${styles.post__heading}`}>
        <h2 className={styles.post__header}>
          <Link to={`/post/${json.id}`}>{json.title}</Link>
        </h2>
        <div className={`${globalStyles.flex_row} ${styles.post__meta}`}>
          <p className="post__author">
            <i>Author: {author.username}</i>
          </p>
          |
          <p className={styles.post__category}>
            <i>Category: {category.title}</i>
          </p>
        </div>
      </div>
      <p className={styles.post__body}>{json.body}</p>
      <p className={styles.post__comments}>
        <b>{comments?.length || 0}</b> comments
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
