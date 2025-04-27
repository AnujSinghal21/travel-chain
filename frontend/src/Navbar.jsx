// components/Navbar.jsx
import { Link, useLocation } from "react-router-dom";

function Navbar({user}) {
  const { pathname } = useLocation();

  const getLinks = () => {
    if (user){
        return (
            <>
          <Link className="nav-link text-white mx-2" to="/browse">
              <i className="bi bi-search"></i> Browse
          </Link>
          <Link className="nav-link text-white mx-2" to="/mybookings">
              <i className="bi bi-calendar-check"></i> My {user.role == "Customer" ? " bookings": " services"}
          </Link>
          
          <Link className="nav-link text-white mx-2" to="/profile">
              <i className="bi bi-person"></i> Profile
          </Link>
          <Link
            className="nav-link text-white mx-2"
            to="/"
            onClick={() => {
              localStorage.removeItem("token");
              localStorage.removeItem("user");
              window.location.href = "/";
            }}
          >
              <i className="bi bi-box-arrow-right"></i> Logout
          </Link>
            </>
        )
    }else{
        return (
            <>
                <Link className="nav-link text-white" to="/login">Login</Link>
                <Link className="nav-link text-white" to="/register">Register</Link>
            </>
        )
        
    }
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-success px-3">
      <Link className="navbar-brand text-white" to="/">Travel-Chain</Link>
      <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
        <span className="navbar-toggler-icon" />
      </button>
      <div className="collapse navbar-collapse" id="navbarNav">
        <div className="navbar-nav ms-auto">
          {getLinks()}
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
