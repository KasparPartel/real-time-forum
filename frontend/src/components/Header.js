import React from "react";
import { useLocation } from "react-router-dom";

export default function Header() {
  const location = useLocation();

  function chooseTitle() {
    switch (location.pathname) {
      case "/":
        return "Post Feed";
      case "/create":
        return "Create a Post";
      case "/messages":
        return "Messages";
      case "/friends":
        return "Friends";
      case "/profile":
        return "Profile";
      default:
        return "";
    }
  }

  return <h1 className="w-full py-3 text-5xl underline">{chooseTitle()}</h1>;
}
