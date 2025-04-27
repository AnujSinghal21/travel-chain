// App.jsx
import { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";
import "bootstrap-icons/font/bootstrap-icons.css";
import axios from "axios";
import Home from "./pages/Home";
import Navbar from "./Navbar";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Profile from "./pages/Profile";
import Browse from "./pages/Browse";
import MyBookings from "./pages/MyBookings";
// import About from "./pages/About";
// import Services from "./pages/Services";
// import Contact from "./pages/Contact";
// import Dashboard from "./pages/Dashboard";
// import Profile from "./pages/Profile";
import NotFound from "./pages/NotFound";

function App() {
  const [token, setToken] = useState(localStorage.getItem('jwt'));
  const [user, setUser] = useState(JSON.parse(localStorage.getItem('user')));
  const [tickets, setTickets] = useState([]);
  useEffect(() => {
    if (token){
      axios.get('/api/tickets', {
        headers: {
          'Authorization': token
        }
      })
      .then(response => {
        setTickets(response.data);
        console.log(response.data);
      })
      .catch(error => {
        console.error('Error fetching tickets:', error);
      });
    }
  }, [token]);

    return (
    <Router>
      <Navbar user={user}/>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login setToken={setToken} setUser={setUser}/>} />
        <Route path="/register" element={<Register />} />
        <Route path="/profile" element={<Profile user={user} />} />
        <Route path="/browse" element={<Browse token={token} tickets={tickets}/>} />
        <Route path="/mybookings" element={<MyBookings user={user} token={token} tickets={tickets}/>} />
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
