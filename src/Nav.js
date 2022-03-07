import { BrowserRouter, Link } from "react-router-dom";
import "./nav.css";

export default function Nav() {
  return (
    <BrowserRouter>
      <header className="nav__header">
        <nav>
          <ul>
            <li className="nav__list-item">
              <Link to="/feed">Feed</Link>
            </li>
            <li className="nav__list-item">
              <Link to="/messages">Messages</Link>
            </li>
            <li className="nav__list-item">
              <Link to="/friends">Friends</Link>
            </li>
          </ul>
        </nav>
      </header>
    </BrowserRouter>
  );
}
