import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export default function Post() {
  const params = useParams();
  const [post, setPost] = useState({});

  useEffect(() => {
    fetchSinglePost();
    console.log(post);
  }, []);

  const fetchSinglePost = () => {
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

  return (
    <div className="card">
      <h2 className="mb-1 text-lg font-bold">{post.title}</h2>
      <p className="font-medium line-clamp-4">{post.body}</p>
    </div>
  );
}
