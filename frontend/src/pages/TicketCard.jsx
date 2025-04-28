function TicketCard({ ticket, option, token }) {
    const statusColor = ticket.status === "Available" ? "success" : "secondary";
    const handleOptionChosen = () => {
        if (option == "Delete"){
            const request = {
                "tids": [ticket.tid],
            }
            fetch("/api/ticket/delete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: token,
                },
                body: JSON.stringify(request),
            })
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.json();
            }
            )
            .then((data) => {
                console.log(data);
            }
            )
        }else if (option == "Book"){
            const request = {
                "tids": [ticket.tid],
            }
            fetch("/api/ticket/book", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: token,
                },
                body: JSON.stringify(request),
            })
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.json();
            }
            )
            .then((data) => {
                console.log(data);
            }
            )
        }else if (option == "Cancel"){
            const request = {
                "tids": [ticket.tid],
            }
            fetch("/api/ticket/cancel", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: token,
                },
                body: JSON.stringify(request),
            })
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.json();
            }
            )
            .then((data) => {
                console.log(data);
            }
            )
        }
    }
  
    return (
      <div className="card shadow-sm rounded-4 border-success mb-4" style={{ background: "#f9fff9" }}>
        <div className="card-body p-4" style={{ position: "relative" }}>
          {/* Top - Service and Transport Type */}
          <div className="d-flex justify-content-between align-items-center mb-3">
            <h5 className="mb-0 text-success">
              {ticket.service_name} <span className="text-muted small fs-6">({ticket.service_provider})</span>
            </h5>
            <h6 className="text-muted" style={{marginRight: "250px"}}>
                {new Date(ticket.start_time).toLocaleDateString([], { year: 'numeric', month: 'long', day: 'numeric' })}
                &nbsp;&minus;&nbsp;
                {new Date(new Date(ticket.start_time).getTime() + ticket.duration * 60000).toLocaleDateString([], { year: 'numeric', month: 'long', day: 'numeric' })}
            </h6>
            <h5 className="badge large bg-success">         
                {ticket.transport_type === "Bus" && <i className="bi bi-bus-front-fill me-2"></i>}
                {ticket.transport_type === "Train" && <i className="bi bi-train-front-fill me-2"></i>}
                {ticket.transport_type === "Flight" && <i className="bi bi-airplane-fill me-2"></i>}
                {ticket.transport_type === "Ship" && <i className="bi bi-ship-fill me-2"></i>}
                {ticket.transport_type}
            </h5>
          </div>
  
          {/* Route */}
          <div className="d-flex justify-content-between align-items-center mb-4">
            <div className="text-center">
              <h6 className="mb-0">{ticket.source}</h6>
              <small className="text-muted">Source</small>
            </div>
            <div className="d-flex align-items-center justify-content-center gap-4">
                {/* Start Time */}
                <div className="text-center">
                    <h6 className="mb-0">
                    {new Date(ticket.start_time).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                    </h6>
                    <small className="text-muted">Start Time</small>
                </div>

                {/* Arrow Icon */}
                <i className="bi bi-arrow-right-circle-fill fs-2 text-success"></i>

                {/* End Time */}
                <div className="text-center">
                    <h6 className="mb-0">
                    {new Date(new Date(ticket.start_time).getTime() + ticket.duration * 60000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                    </h6>
                    <small className="text-muted">End Time</small>
                </div>
            </div>
            <div className="text-center">
              <h6 className="mb-0">{ticket.destination}</h6>
              <small className="text-muted">Destination</small>
            </div>
          </div>
  
          {/* Details */}
          <div className="row text-center mb-3">
            <div className="col">
              <strong>Seat No</strong>
              <div>{ticket.seat}</div>
            </div>
            <div className="col">
              <strong>Price</strong>
              <div>₹{ticket.price}</div>
            </div>
            <div className="col">
              <strong>Status</strong>
              <div>
                <span className={`badge bg-${statusColor}`}>{ticket.status}  {option=="Delete"? ticket.passenger: ""}</span>
              </div>
            </div>
          </div>
  
          {/* Book Now Button */}
          <button
            className="btn btn-outline-success w-100"
            onClick={handleOptionChosen}
          >
            {option}
          </button>
  
          {/* Decorative Divider (like plane ticket punch effect) */}
          <div
            style={{
              position: "absolute",
              top: "50%",
              left: "-10px",
              width: "20px",
              height: "20px",
              background: "white",
              borderRadius: "50%",
              boxShadow: "0 0 0 1px #dee2e6",
            }}
          ></div>
          <div
            style={{
              position: "absolute",
              top: "50%",
              right: "-10px",
              width: "20px",
              height: "20px",
              background: "white",
              borderRadius: "50%",
              boxShadow: "0 0 0 1px #dee2e6",
            }}
          ></div>
        </div>
      </div>
    );
  }
  
  export default TicketCard;
  