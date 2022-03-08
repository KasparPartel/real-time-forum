import { BrowserRouter, Routes, Route } from "react-router-dom";

import Nav from "./Nav";
import Feed from "./Feed";

export default function App() {
  return (
    <div className="flex">
      <Nav />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Feed />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}
