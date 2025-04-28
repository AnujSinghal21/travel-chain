
function Profile({ user }) {
    if (!user){
        window.location.href = "/";
        return null;
    }

  return (
    <div className="container py-5" style={{ maxWidth: "600px" }}>
      <div className="card shadow rounded-4">
        <div className="card-body">
          <h3 className="card-title text-success mb-4">
            <i className="bi bi-person-circle me-2"></i>Profile
          </h3>
          <ul className="list-group list-group-flush">
            <li className="list-group-item">
              <strong>Name:</strong> {user.name}
            </li>
            <li className="list-group-item">
              <strong>Email:</strong> {user.email}
            </li>
            <li className="list-group-item">
              <strong>Age:</strong> {user.age}
            </li>
            <li className="list-group-item">
              <strong>Phone:</strong> {user.phone}
            </li>
            <li className="list-group-item">
              <strong>Role:</strong> {user.role}
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}

export default Profile;
