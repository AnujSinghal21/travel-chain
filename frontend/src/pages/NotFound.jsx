// pages/NotFound.jsx
import { Link } from "react-router-dom";

function NotFound() {
  return (
    <div className="container text-center py-5">
      <h1 className="display-4 text-success">404</h1>
      <p className="lead">Page not found</p>
      <Link to="/" className="btn btn-success">
        <i className="bi bi-house-door-fill me-2"></i>
        Go Home
      </Link>
    </div>
  );
}

export default NotFound;
