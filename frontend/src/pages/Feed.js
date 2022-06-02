import React, { useEffect, useState } from "react";
import FeedPost from "./FeedPost";

function Feed(props) {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    requestPosts();
  }, []);

  const requestPosts = () => {
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
      {posts?.map((p) => (
        <FeedPost key={p.id} json={p} />
      ))}
    </div>
  );
}

export default Feed;
