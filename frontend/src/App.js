import {Link, Navigate, Route, Routes} from 'react-router-dom';
import './App.css';
import { webSocketConnect } from "./websocket"
// import React, { useEffect } from "react";

import Layout from './components/layout/Layout';
import Feed from './pages/Feed';
import CreatePost from './pages/CreatePost';
import Messages from './pages/Messages';
import Profile from './pages/Profile';
import NotFound from './pages/NotFound';
import Login from './pages/Login';
import Register from './pages/Register';
import {useEffect, useState} from "react";

function App() {
    const [loggedIn, setLoggedIn] = useState(false)

    useEffect(() => {
        setLoggedIn(true)
    }, [])

    useEffect(() => {
      webSocketConnect("ws://localhost:4000/ws");
    }, []);

    return (
        <Layout loginState={loggedIn}>
            <Routes>
                <Route path="/" exact element={<Feed/>}/>
                <Route path="/login" element={<Login />} />
                <Route path="/logout" element={<Logout />} />
                <Route path="/register" element={<Register/>}/>
                <Route path="/create-post" element={<CreatePost/>}/>
                <Route path="/messages" element={<Messages/>}/>
                <Route path="/profile" element={<Profile/>}/>
                <Route path="*" element={<NotFound/>}/>
            </Routes>
        </Layout>
    );
  
  
}

export default App;
