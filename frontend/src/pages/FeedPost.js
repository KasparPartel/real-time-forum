import { Link } from "react-router-dom";

export default function FeedPost(props) {
  return (
    <div className="card">
      <h2 className="mb-1 text-lg font-bold">
        <Link to={`/post/${props.json.id}`}>{props.json.title}</Link>
      </h2>
      <p className="font-medium line-clamp-4">{props.json.body}</p>
      <p>
        <span className="comment-amount text-gray-500">
          {/* <Link to="#!"> */}
          {props.json.commentAmount ? props.json.commentAmount : 0} comments
          {/* </Link> */}
        </span>
        <span className="like-amount text-gray-500">
          {props.json.likeAmount ? props.json.likeAmount : 0} likes
        </span>
        <span className="dislike-amount text-gray-500">
          {props.json.dislikeAmount ? props.json.dislikeAmount : 0} dislikes
        </span>
      </p>
    </div>
  );
}
