import { Fragment } from "react";
import { Link } from "react-router-dom";

import classes from "./MainNavigation.module.css";

export default function MainNavigation({ user }) {
  return (
    <header className={classes.header}>
      <div className={classes.logo}>real-time-forum</div>
      <nav>
        <ul>
          {/* <li>
            <Link to="/">All posts</Link>
          </li> */}
          {user ? (
            <Fragment>
              {/* <li>
                <Link to="/create-post">Add new post</Link>
              </li> */}
              <li>
                <Link to="/profile">{user.username}</Link>
              </li>
            </Fragment>
          ) : (
            <Fragment></Fragment>
          )}
        </ul>
      </nav>
    </header>
  );
}
