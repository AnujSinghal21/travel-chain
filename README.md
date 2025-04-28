#Travel-Chain
Travel-Chain is an Hyperledger based online ticket booking system, supported for multiple users and service providers, just liek current travel booking sites. This is the course assignment for CS731, under Prog. Angshuman Karmamkar.

##Main Features and Ideas
The code structure is divided into broadly 3 parts : 
1. hyperledger (contains the chaincodes and main blockchain structures)
2. backend (the api calls, that connects with the blockchain fabric)
3. frontend (a functional ui design)
4. Our structure contains 2 organizations, each having 1 peer node.

## Report Link
https://docs.google.com/document/d/1nWAvZShAmZd0bYByG0k0o5g4vE3dVqUYQnUlfZCAZYU/edit?usp=sharing

### chaincodes
- Functionalities implemented for interacting with teh blockchain, and integrated with the backend
- Golang is used for the chaincode development
- Chaincodes are properly structured in multiple files for modularity (user related functions and ticket related functiosn are separated)
- The main structure stored in the ledger is the bookings or the tickets, along with users or providers. (A bit costly, but simplifies some good functional features)

### backend
- Golang is used for backend development, fabrikSDK is used for integration with blockchain fabric.
- Routes for api calls are propery defined, separate backend controller file sofr user and ticket related functioanlities.
- Implemented the initialization of fabric as a middleware, constructed a .yaml file for connection with hyperledger entities.

### frontend
- React.js and bootstrap is used for frontend development, a user friendly UI, with multiple features
- Complex features such as connected journey search are implemented with the help of the frontend and chaincode interactions, justifying the costly stores on the ledger.
- Primary targets are desktop users, fats and responsive UI

## Code Running
### Prerequisites
- Set up hyperledger fabric 
- Set up GO environment
- Set up docker
- Set up node.js
### setting up the network and deploying the chaincodes
- ` ./network.sh up createChannel -c mychannel -ca -s couchdb ` 
- ` /network.sh deployCC -ccn mycc -ccp ../chaincode/ticketbooking -ccl go -ccv 1.0 -c mychannel `
### backend
- ` go mod tidy `
- ` go run . `
### frontend
- ` npm install `
- ` npm run dev `


