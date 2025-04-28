import { useEffect, useState } from "react";
import Filter from "./Filters"
import TicketCard from "./TicketCard";

function Browse({token, tickets}) {
  const [filteredTickets, setFilteredTickets] = useState([]);

  return (
    <div className="container py-4">
      {/* Toolbar Placeholder */}
      <Filter filtered={filteredTickets} setFiltered={setFilteredTickets} tickets={tickets}/>
 
      {/* Ticket Cards */}
      <div className="row g-4">
        {filteredTickets.map((ticket) => (
            <TicketCard ticket={ticket} key={ticket.tid} option={"Book"} token={token}/>
        ))}
      </div>
    </div>
  );
}

export default Browse;
