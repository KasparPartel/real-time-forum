import {Route, Routes} from 'react-router-dom';
import {useEffect, useState} from "react";

import Layout from './components/layout/Layout';
import Feed from './pages/Feed';
import CreatePost from './pages/CreatePost';
import Messages from './pages/Messages';
import Profile from './pages/Profile';
import NotFound from './pages/NotFound';
import Login from './pages/Login';
import Register from './pages/Register';

import './App.css';
import {useCookies} from "react-cookie";

function App() {
    const [cookies, setCookie, removeCookie] = useCookies(['session_token']);
    const [user, setUser] = useState({})

    useEffect(() => {
        cookies["session_token"] ? getUser() : setUser(null)
    }, [cookies])

    const getUser = async () => {
        const res = await fetch("http://localhost:4000/v1/api/user/me", {
            method: "GET",
            credentials: "include"
        })

        if (!res.ok) {
            const msg = `User not authorized: ${res.status}`
            setUser(null)
            console.log(msg)
            return
        }

        const user = await res.json()
        setUser(user)
        console.log(user)
    }

    const Logout = () => {
        removeCookie("session_token")
    }

    return (
        <Layout user={user}>
            <Routes>
                <Route path="/" exact element={<Feed/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/logout" element={<Logout/>}/>
                <Route path="/register" element={<Register/>}/>
                <Route path="/create-post" element={<CreatePost/>}/>
                <Route path="/messages" element={<Messages/>}/>
                <Route path="/profile" element={<Profile user={user}/>}/>
                <Route path="*" element={<NotFound/>}/>
            </Routes>
        </Layout>
    );
}

export default App;
