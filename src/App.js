import { Routes, Route } from "react-router-dom";

import { Feed, Messages, Header, NavColumn, FriendsColumn } from "./components";

export default function App() {
  return (
    <div className="flex">
      <NavColumn />
      <div className="grid px-10 basis-full min-h-screen bg-purple-100">
        <Header />
        <Routes>
          <Route path="/" element={<Feed />} />
          <Route path="/messages" element={<Messages />} />
        </Routes>
      </div>
      <FriendsColumn />
    </div>
  );
}
