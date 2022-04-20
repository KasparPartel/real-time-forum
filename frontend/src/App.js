import { Route, Routes } from 'react-router-dom';
import './App.css';

import Layout from './components/layout/Layout';
import Feed from './pages/Feed';
import CreatePost from './pages/CreatePost';
import Messages from './pages/Messages';
import Profile from './pages/Profile';
import NotFound from './pages/NotFound';
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" exact element={<Feed />} />
        <Route path="/create-post" element={<CreatePost />} />
        <Route path="/messages" element={<Messages />} />
        <Route path="/profile" element={<Profile />} />
        {/* <Route path="/myposts" element={<MyPosts />} /> */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Layout>
    
  );
}

export default App;
