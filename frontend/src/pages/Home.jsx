// pages/Home.jsx
import React from "react";
function Home() {
    return (
    <div className="container py-5">
        <h1 className="text-success">Welcome to Travel-Chain. A blockchain based ticket booking system.</h1>
        <p className="lead">
            Travel-Chain is a decentralized ticket booking system that leverages blockchain technology to provide a secure and transparent platform for users to book tickets for various events and services.
            With Travel-Chain, users can enjoy a seamless booking experience while ensuring the authenticity and integrity of their transactions.
            Our platform utilizes smart contracts to automate the booking process, eliminating the need for intermediaries and reducing costs.
            By using blockchain, we ensure that all transactions are recorded on a public ledger, providing transparency and security for both users and service providers.
            Join us in revolutionizing the ticket booking industry with Travel-Chain.
        </p>
        <h2 className="text-success">Features</h2>
        <ul className="lead">
            <li>Decentralized booking system</li>
            <li>Secure and transparent transactions</li>
            <li>Smart contracts for automated booking</li>
            <li>User-friendly interface</li>
        </ul>
        <h2 className="text-success">Acknowledgements</h2>
        <p className="lead">
            This project is done as part of CS731 - Introduction to Blockchain Technology Course at IIT Kanpur. <br></br>
            We would like to thank our professor, Dr. Angshuman Karmakar, for his guidance and support throughout the course. <br></br>
            We woulld also like to express our heartfelt gratitude to all our TAs for their valuable feedback and continuous help. <br></br> 
        </p>
        <h2 className="text-success">Developers</h2>
        <div>
            The complete project code along with docs is hosted on Github at <a href="https://github.com/AnujSinghal21/travel-chain"> https://github.com/AnujSinghal21/travel-chain</a>.
            <div className="card" style={{ width: "18rem", margin: "1rem auto", border: "1px solid #28a745" }}>
                <div className="card-body">
                    <h5 className="card-title text-success">Team Members</h5>
                    <p className="card-text">
                        <strong>Name:</strong> Anuj<br />
                        <strong>Email:</strong> anuj21@iitk.ac.in<br />
                        <strong>Roll No:</strong> 210166<br />
                    </p>
                    <hr />
                    <p className="card-text">
                        <strong>Name:</strong> Goutam Das<br />
                        <strong>Email:</strong> goutamd21@iitk.ac.in<br />
                        <strong>Roll No:</strong> 210394<br />
                    </p>
                </div>
            </div>
        </div>
        

    </div>
    );
}
  
export default Home;
  