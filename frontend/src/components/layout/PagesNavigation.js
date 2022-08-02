import { Link } from "react-router-dom";
import { Fragment } from "react";

import classes from "./PagesNavigation.module.css";

export default function PagesNavigation({ user }) {
  return (
    <nav className={classes.pages}>
      <ul>
        <li>
          <Link to="/">Feed</Link>
        </li>
        {user ? (
          <Fragment>
            <li>
              <Link to="/create-post">Create new post</Link>
            </li>
            <li>
              <Link to="/profile">User profile</Link>
            </li>
            <li>
              <Link to="/logout">Logout</Link>
            </li>
          </Fragment>
        ) : (
          <Fragment>
            <li>
              <Link to="/login">Log in</Link>
            </li>
            <li>
              <Link to="/register">Register</Link>
            </li>
          </Fragment>
        )}
      </ul>
    </nav>
  );
}
