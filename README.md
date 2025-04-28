# travel-chain
CS731 course project.

## Chaincode (Go)

## Backend (Go)
API routes that will be required
- `/createUser` [POST]
    given details of users, signup the user [ call to CA to get the credentials which user can download and use later for signing ]
- `/login` [POST]
    finds the user details
- `/getUserDetails` [GET]
    gets the user details {name, email, phone, id, certificates, balance}
- `/getAvailableBookings` [GET]
    RBAC -> ['customer': get all available bookings details, all filters applicable in frontend]
- `/showMyBookings` [GET]
    RBAC -> ['customer': get all his bookings, 'serviceProvider': show all bookings he has offered]
- `/bookTicket` [POST]
    'customer': books the given tickets
- `/createSeat` [POST]
    'serviceProvider': create a new seat
- `/cancelBooking` [POST]
    RBAC -> ['customer': unbook his ticket (refund given with some deduction), 'serviceProvider': remove the ticket  (full refund)]
- `/addMoney` [POST]
    Increment balance of user

## Frontend (React)

