import { BrowserRouter, Link } from "react-router-dom";
import { useState } from "react";

import { HiPlus } from "react-icons/hi";
import { BiEnvelope, BiLayer, BiLogOut, BiUser } from "react-icons/bi";
import { MdOutlinePeopleAlt } from "react-icons/md";

export default function Nav() {
  const [active, setActive] = useState(false);

  return (
    <BrowserRouter>
      {active ? (
        <nav className="flex flex-col gap-20 sticky top-0 h-screen pl-4 pr-5 py-6 bg-red-100">
          <HiPlus
            size="1.6em"
            className="self-end transition-transform duration-500 rotate-45 cursor-pointer"
            onClick={() => setActive(false)}
          />
          <ul className="flex flex-col gap-3">
            <li className="nav__list-item">
              <Link to="/">Feed</Link>
            </li>
            <li className="nav__list-item">
              <Link to="/messages">Messages</Link>
            </li>
            <li className="nav__list-item">
              <Link to="/friends">Friends</Link>
            </li>
            <li className="nav__list-item">
              <Link to="/profile">Profile</Link>
            </li>
          </ul>

          <Link to="/logout" className="mt-auto">
            Logout
          </Link>
        </nav>
      ) : (
        <nav className="flex flex-col items-center gap-20 sticky top-0 h-screen px-4 py-6 bg-red-100">
          <HiPlus
            size="1.6em"
            className="cursor-pointer"
            onClick={() => setActive(true)}
          />
          <ul className="flex flex-col items-center gap-10">
            <li className="nav__list-item">
              <Link to="/">
                <BiLayer size="1.4em" />
              </Link>
            </li>
            <li className="nav__list-item">
              <Link to="/messages">
                <BiEnvelope size="1.4em" />
              </Link>
            </li>
            <li className="nav__list-item">
              <Link to="/friends">
                <MdOutlinePeopleAlt size="1.4em" />
              </Link>
            </li>
            <li className="nav__list-item">
              <Link to="/profile">
                <BiUser size="1.4em" />
              </Link>
            </li>
          </ul>

          <Link to="/logout" className="mt-auto">
            <BiLogOut size="1.4em" />
          </Link>
        </nav>
      )}
    </BrowserRouter>
  );
}
