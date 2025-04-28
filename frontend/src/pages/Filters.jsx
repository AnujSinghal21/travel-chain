import { useState, useEffect } from "react";

function Filter({ filtered, setFiltered, tickets }) {
    const [filters, setFilters] = useState({
      source: "",
      destination: "",
      date: "",
      allowConnected: false,
      transportType: "All",
      priceHigh: 100000,
      sortBy: "price" // "duration", "Start Time"
    });
    useEffect(() => {
        const f = tickets.filter((t) => {
            let ans = true;
            if (t.status != "Available"){
                ans = false;
            }
            if (filters.source != "" && t.source != filters.source){
                ans = false;
            }
            if (filters.destination != "" && t.destination != filters.destination){
                ans = false;
            }
            if (filters.transportType != "All" && t.transport_type != filters.transportType){
                ans = false;
            }
            if (filters.date != "" && new Date(t.start_time).toDateString() != new Date(filters.date).toDateString()) {
                ans = false;
            }
            if (t.price > filters.priceHigh){
                ans = false;
            }
            return ans; 
        });
        const uniqueServices = new Map();
        f.forEach((ticket) => {
            if (!uniqueServices.has(ticket.service_id)) {
                uniqueServices.set(ticket.service_id, ticket);
            }
        });
        const f2 = Array.from(uniqueServices.values());

        f2.sort((t1, t2) => {
            if (filters.sortBy == "price"){
                return t1.price < t2.price
            } else if (filters.sortBy == "start_time") {
                return new Date(t1.start_time) - new Date(t2.start_time);
            }else{
                return t1.duration < t2.duration
            }
        });
        setFiltered(f2)
    }, [tickets, filters]);
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFilters((prev) => ({ ...prev, [name]: value }));
      };
    return (
        <div className="row g-3 mt-3 mb-3">
        <div className="col-md-2 text-center">
            <label className="form-label">
                <i className="bi bi-calendar-event"></i>
            </label>
            <input
                type="date"
                name="date"
                className="form-control"
                value={filters.date || ""}
                onChange={handleChange}
            />
        </div>
        <div className="col-md-2 text-center">
            <label className="form-label">
            <i className="bi bi-geo-alt-fill"></i>
            </label>
            <input
            type="text"
            name="source"
            className="form-control"
            placeholder="Source Location"
            value={filters.source || ""}
            onChange={handleChange}
            />
        </div>
        <div className="col-md-2 text-center">
            <label className="form-label">
            <i className="bi bi-geo-alt"></i>
            </label>
            <input
            type="text"
            name="destination"
            className="form-control"
            placeholder="Destination"
            value={filters.destination || ""}
            onChange={handleChange}
            />
        </div>
        <div className="col-md-2 text-center">
            <label className="form-label">
            <i className="bi bi-bus-front-fill"></i>
            </label>
            <select
            name="transportType"
            className="form-select"
            value={filters.transportType || ""}
            onChange={handleChange}
            >
            <option value="">All</option>
            <option value="Bus">Bus</option>
            <option value="Train">Train</option>
            <option value="Flight">Flight</option>
            </select>
        </div>

        <div className="col-md-2 text-center">
            <label className="form-label">
            Max Price ₹
            </label>
            <input
            type="number"
            name="priceHigh"
            className="form-control"
            placeholder="Max Price"
            value={filters.priceHigh || ""}
            onChange={handleChange}
            />
        </div>
        <div className="col-md-2 text-center">
            <label className="form-label">
            <i className="bi bi-sort-alpha-down"></i>
            </label>
            <select
            name="sortBy"
            className="form-select"
            value={filters.sortBy || ""}
            onChange={handleChange}
            >
            <option value="price">Price</option>
            <option value="duration">Duration</option>
            <option value="start_time">Start time</option>
            </select>
        </div>
        <div className="col-12"></div>
            <hr className="my-4" style={{ borderTop: "5px solid rgb(67, 139, 62)", opacity: 0.7 }} />
        </div>
    );

  }
  
  export default Filter;
  