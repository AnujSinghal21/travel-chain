// App.jsx
import { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";
import "bootstrap-icons/font/bootstrap-icons.css";
import * as jwt_decode from 'jwt-decode';
import Home from "./pages/Home";
import Navbar from "./Navbar";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Profile from "./pages/Profile";
// import About from "./pages/About";
// import Services from "./pages/Services";
// import Contact from "./pages/Contact";
// import Dashboard from "./pages/Dashboard";
// import Profile from "./pages/Profile";
import NotFound from "./pages/NotFound";

function App() {
  const [token, setToken] = useState(localStorage.getItem('jwt'));
  const [user, setUser] = useState(JSON.parse(localStorage.getItem('user')));

  return (
    <Router>
      <Navbar user={user}/>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login setToken={setToken} setUser={setUser}/>} />
        <Route path="/register" element={<Register />} />
        <Route path="/profile" element={<Profile user={user} />} />
        {/* <Route path="/register" element={<About />} />
        <Route path="/login" element={<Services />} />
        <Route path="/mybookings" element={<Contact />} />
        <Route path="/browse" element={<Dashboard />} />
        <Route path="/profile" element={<Profile />} /> */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  );
}

export default App;
