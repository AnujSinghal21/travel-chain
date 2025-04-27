import React, { useState, useEffect } from 'react';
import TicketCard from './TicketCard';
import { useNavigate } from 'react-router-dom';

const MyBookings = ({user, token, tickets}) => {
  const [myTickets, setMyTickets] = useState([]);
  const [bookingDialog, setBookingDialog] = useState(false);
  useEffect(() => {
    const t = tickets.filter((t) => {
        if (user.role == "Customer"){
            return t.passenger === user.email
        }else{
            return t.service_provider === user.email
        }
      });
    setMyTickets(t);    
  }, [tickets]);
  const handleCreateService = (e) => {
    e.preventDefault();
    const formData = {
        serviceName: document.getElementById('serviceName').value,
        departure: document.getElementById('departure').value,
        destination: document.getElementById('destination').value,
        departureTime: document.getElementById('departureTime').value,
        duration: parseInt(document.getElementById('duration').value, 10),
        price: parseFloat(document.getElementById('price').value),
        capacity: parseInt(document.getElementById('capacity').value, 10),
        transportType: document.getElementById('transport_type').value,
    };
    let tickets = []
    let service_id = "service" + (Math.floor(Math.random() * 1e9) + 1);
    for (let i = 0; i < formData.capacity; i++){
        tickets.push({
            "tid": Math.floor(Math.random() * 1e9) + 1,
            "passenger": "",
            "service_id": service_id,
            "service_provider": user.email,
            "service_name": formData.serviceName,
            "status": "Available",
            "price": formData.price,
            "start_time": formData.departureTime,
            "duration": formData.duration,
            "source": formData.departure,
            "destination": formData.destination,
            "seat": i+1,
            "transport_type": formData.transportType
        })
    }

    fetch('/api/ticket/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: token,
        },
        body: JSON.stringify(tickets),
    })
        .then(response => {
            if (!response.ok) {
                console.log(response);
                throw new Error('Failed to create service');
            }
            return response.json();
        })
        .then(data => {
            alert('Service created successfully!');
            setBookingDialog(false);
        })
        .catch(error => {
            console.error('Error:', error);
        });
  };

  if (!user) return null;

  return (
    <div className="container my-4">
        {
            bookingDialog? (
                <div className="text-center mb-4">
                    <form className='form border border-2 rounded-4 p-4 shadow-sm'>
                    <div className="mb-3">
                        <label htmlFor="serviceName" className="form-label">Service Name</label>
                        <input type="text" className="form-control" id="serviceName" placeholder="Enter service name" required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="departure" className="form-label">Departure Location</label>
                        <input type="text" className="form-control" id="departure" placeholder="Enter Source Location" required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="destination" className="form-label">Destination</label>
                        <input type="text" className="form-control" id="destination" placeholder="Enter destination" required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="departureTime" className="form-label">Departure Time</label>
                        <input type="datetime-local" className="form-control" id="departureTime" required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="departureTime" className="form-label">Journey Duration (in minutes) </label>  
                        <input type="number" className="form-control" id="duration" placeholder='240' required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="price" className="form-label">Price</label>
                        <input type="number" className="form-control" id="price" placeholder="Enter price" required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="capacity" className="form-label">Capacity</label>
                        <input type="number" className="form-control" id="capacity" placeholder="Enter no. of seats" required />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="modeOfTransport" className="form-label">Mode of Transport</label>
                        <select className="form-select" id="transport_type" required>
                            <option value="">Select mode of transport</option>
                            <option value="Bus">Bus</option>
                            <option value="Train">Train</option>
                            <option value="Flight">Flight</option>
                        </select>
                    </div>
                    <button type="submit" className="btn btn-success" onClick={handleCreateService}>Submit</button>
                    </form>
                </div>
            ) : (<></>)
        }
      {user.role === 'Transport Provider' && (
        <div className="text-center mb-4">
            <button className="btn btn-success" onClick={() => setBookingDialog(!bookingDialog)}>
                {bookingDialog? "Close": "Create New Service"}
            </button>
        </div>
      )}
      <h3> Your {user.role == "Customer"? "Tickets": "Services"}</h3>
      <div className="row g-4">
        {myTickets.length > 0 ? (
          myTickets.map(ticket => (
            <TicketCard ticket={ticket} key={ticket.tid} option={user.role == "Customer"? "Cancel": "Delete"} token={token}/>
          ))
        ) : (
          <p className="text-muted">No tickets found.</p>
        )}
      </div>
    </div>
  );
};

export default MyBookings;
