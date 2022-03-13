import { Link } from "react-router-dom";
import React from "react";

export default function Post(props) {
  return (
    <div className="flex flex-col justify-between flex-1 basis-52 md:basis-80 basis-96 p-7 bg-blue-200">
      <h2 className="mb-1 text-lg font-bold">
        <Link to="#!">{props.json.title}</Link>
      </h2>
      <p className="font-medium line-clamp-4">{props.json.body}</p>
      <p>
        <span className="text-gray-500">
          <Link to="#!">3 comments</Link>
        </span>
      </p>
    </div>
  );
}
