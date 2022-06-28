import { Fragment, useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import CommentTree from "../components/layout/CommentTree";
import CreateComment from "../components/layout/CreateComment";

export default function Post() {
  const params = useParams();

  const [post, setPost] = useState({});
  const [comments, setComments] = useState([]);
  const [category, setCategory] = useState({});
  const [author, setAuthor] = useState({});

  // const [likes, setLikes] = useState([]);
  // const [dislikes, setDislikes] = useState([]);

  const user_id = sessionStorage.getItem("user_id");

  useEffect(() => {
    getSinglePost();
  }, []);

  useEffect(() => {
    getComments();
    getCategory();
    getUser();

    // getLikes();
    // getDislikes();
  }, [post]);

  // getSinglePost fetches post with given id from api and sets post object
  const getSinglePost = () => {
    fetch(`http://localhost:4000/v1/api/post/${params.id}`, {
      method: "GET",
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else {
          console.log("cannot fetch post");
          return;
        }
      })
      .then((data) => {
        setPost(data[0]);
      });
  };

  const getComments = () => {
    fetch(`http://localhost:4000/v1/api/comments/${params.id}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setComments(data);
        console.log("comments", comments);
      });
  };

  const getCategory = () => {
    fetch(`http://localhost:4000/v1/api/categories/${post["category_id"]}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setCategory(data[0]);
      });
  };

  const getUser = () => {
    fetch(`http://localhost:4000/v1/api/user/${post["user_id"]}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        setAuthor(data[0]);
      });
  };

  // const getLikes = () => {
  //   fetch(`http://localhost:4000/v1/api/like/${params.id}`, {
  //     method: "GET",
  //   })
  //     .then((res) => res.json())
  //     .then((data) => {
  //       setLikes(data);
  //     });
  // };

  // const getDislikes = () => {
  //   fetch(`http://localhost:4000/v1/api/dislike/${params.id}`, {
  //     method: "GET",
  //   })
  //     .then((res) => res.json())
  //     .then((data) => {
  //       setDislikes(data);
  //     });
  // };

  return (
    <div className="card">
      <h2 className="mb-1 text-lg font-bold">{post.title} </h2>
      <p>
        <i>Author: {author.username}</i>
      </p>
      <p>
        <i>Category: {category.title}</i>
      </p>
      <p className="font-medium line-clamp-4">{post.body}</p>
      <hr />
      {user_id ? (
        <CreateComment
          postID={params.id}
          userID={user_id}
          getComments={getComments}
        />
      ) : (
        <Fragment></Fragment>
      )}
      <CommentTree comments={comments} />
    </div>
  );
}
