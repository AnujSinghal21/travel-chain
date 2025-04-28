import { useState } from "react";
import { useNavigate } from "react-router-dom";

function Login({ setToken, setUser }) {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!email || !password) {
      alert("Email and password can not be blank.");
      return;
    }

    const formData = {email, password};
    
    const res = await fetch("/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });

    const data = await res.json();
    if (res.ok && data.token) {
      localStorage.setItem("jwt", data.token);
      localStorage.setItem("user", JSON.stringify(data.user));
      console.log(data);
      setToken(data.token);
      setUser(data.user);
      // console.log(data.token);
      // console.log(data.user)
      navigate("/profile");
    } else {
      console.log(res.details);
      alert("Authentication failed");
    }
  };

  return (
    <div className="container py-5" style={{ maxWidth: "500px" }}>
      <h2 className="text-success mb-4">Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="form-label">Email</label>
          <input
            type="email"
            className="form-control"
            value={email}
            onChange={handleEmailChange}
            placeholder="Enter your email"
            required
          />
        </div>
        <div className="mb-4">
          <label className="form-label">Password</label>
          <input
            type="password"
            value={password}
            onChange={handlePasswordChange}
            className="form-control"
            placeholder="Enter your password"
            required
          />
        </div>
        <button type="submit" className="btn btn-success w-100">
          <i className="bi bi-box-arrow-in-right me-2"></i>Login
        </button>
      </form>
    </div>
  );
}

export default Login;
