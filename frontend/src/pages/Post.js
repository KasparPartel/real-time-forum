import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import CommentTree from "../components/layout/CommentTree";
import CreateComment from "../components/layout/CreateComment";
import { UserContext } from "../UserContext";

export default function Post() {
  const params = useParams();
  const { user } = useContext(UserContext);

  const [post, setPost] = useState({});
  const [comments, setComments] = useState([]);

  useEffect(() => {
    getSinglePost();
    console.log(post);
    getComments();
  }, []);

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
    fetch(`http://localhost:4000/v1/api/comments/${post.id}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => setComments(data));
  };

  return (
    <div className="card">
      <h2 className="mb-1 text-lg font-bold">{post.title}</h2>
      <p className="font-medium line-clamp-4">{post.body}</p>
      <hr />
      <CreateComment
        postID={params.id}
        userID={user.id}
        getComments={getComments}
      />
      <CommentTree postID={params.id} getComments={getComments} />
    </div>
  );
}
