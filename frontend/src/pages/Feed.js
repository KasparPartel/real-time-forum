import React, { useEffect, useState } from "react";
import FeedPost from "../components/layout/FeedPost";

export default function Feed() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    getPosts();
  }, []);

  // getPosts fetches all posts from api
  const getPosts = () => {
    fetch("http://localhost:4000/v1/api/post/", {
      method: "GET",
    })
      .then((response) => {
        if (response.ok) {
          console.log("okay");
          return response.json();
        }
        throw new Error("Something went wrong");
      })
      .then((json) => setPosts(json))
      .catch((error) => console.log(error));
  };

  return (
    <div className="flex flex-wrap py-5 gap-4">
      {posts?.map((post) => (
        <FeedPost key={post.id} json={post} />
      ))}
    </div>
  );
}
