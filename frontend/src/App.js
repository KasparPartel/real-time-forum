import React from "react";
import { Routes, Route } from "react-router-dom";

import {
  Feed,
  Messages,
  Header,
  NavColumn,
  FriendsColumn,
  CreatePost,
} from "./components";

export default function App() {
  return (
    <div className="flex">
      <NavColumn />
      <div className="px-10 basis-full min-h-screen bg-blue-200">
        <Header />
        <Routes>
          <Route path="/create" element={<CreatePost />} />
          <Route path="/" element={<Feed />} />
          <Route path="/messages" element={<Messages />} />
        </Routes>
      </div>
      <FriendsColumn />
    </div>
  );
}
