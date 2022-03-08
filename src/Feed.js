import { useEffect, useState } from "react";

export default function Feed() {
  const [posts, setPosts] = useState(null);

  useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => {
        if (response.ok) {
          return response.json();
        }
        throw response;
      })
      .then((json) => {
        setPosts(json);
      });
  }, []);

  return (
    <div className="flex flex-wrap px-10 py-5 gap-4">
      {posts.map((p) => (
        <div className="max-w-sm h-min p-7 bg-green-100">
          <h2 className="mb-3 text-lg font-bold">
            <a href="#!">{p.title}</a>
          </h2>
          <p className="font-medium">{p.body}</p>
        </div>
      ))}
    </div>
  );
}
