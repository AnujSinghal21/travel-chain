// components/Navbar.jsx
import { Link, useLocation } from "react-router-dom";

function Navbar() {
  const { pathname } = useLocation();

  const commonLinks = (
    <>
      <Link className="nav-link text-white" to="/">Home</Link>
      <Link className="nav-link text-white" to="/about">About</Link>
      <Link className="nav-link text-white" to="/services">Services</Link>
      <Link className="nav-link text-white" to="/contact">Contact</Link>
    </>
  );

  const authLinks = (
    <>
      <Link className="nav-link text-white" to="/dashboard">Dashboard</Link>
      <Link className="nav-link text-white" to="/profile">Profile</Link>
    </>
  );

  const getLinks = () => {
    if (pathname.startsWith("/dashboard") || pathname.startsWith("/profile")) {
      return authLinks;
    }
    return commonLinks;
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-success px-3">
      <Link className="navbar-brand text-white" to="/">GreenSite</Link>
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
