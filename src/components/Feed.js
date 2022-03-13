import React, { useEffect, useState } from "react";

import Post from "./Post";

export default function Feed() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    requestPosts();
  }, []);

  const requestPosts = () => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => {
        if (response.ok) {
          return response.json();
        }
        throw new Error("Something went wrong");
      })
      .then((json) => setPosts(json))
      .catch((error) => console.log(error));
  };

  return (
    <div className="flex flex-wrap py-5 gap-4">
      {posts.map((p) => (
        <Post json={p} />
      ))}
    </div>
  );
}
