#!/bin/bash

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

ORDERER_TLS_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
ORG1_PEER_TLS=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
ORG2_PEER_TLS=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

INVOKE() {
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_TLS_CA \
        -C mychannel -n mycc \
        --peerAddresses localhost:7051 --tlsRootCertFiles $ORG1_PEER_TLS \
        --peerAddresses localhost:9051 --tlsRootCertFiles $ORG2_PEER_TLS \
        -c "$1"
    sleep 2
}

QUERY() {
    CORE_PEER_TLS_ENABLED=false peer chaincode query -C mychannel -n mycc -c "$1"
    sleep 2
}

# --- Users ---
INVOKE '{"function":"createUser","Args":["alex@journeys.io","Alex","user","alexpwd","1300"]}'
INVOKE '{"function":"createUser","Args":["bella@skypass.com","Bella","provider","bellapwd","5800"]}'
INVOKE '{"function":"createUser","Args":["cameron@travelsphere.com","Cameron","user","cameronpwd","950"]}'
INVOKE '{"function":"createUser","Args":["danielle@airhop.com","Danielle","provider","daniellepwd","4200"]}'
INVOKE '{"function":"createUser","Args":["elliot@roadconnect.io","Elliot","provider","elliotpwd","4700"]}'
INVOKE '{"function":"createUser","Args":["fiona@pathfinders.net","Fiona","user","fionapwd","1600"]}'
INVOKE '{"function":"createUser","Args":["gavin@trailmix.com","Gavin","user","gavinpwd","1900"]}'
INVOKE '{"function":"createUser","Args":["harper@skyexpress.com","Harper","provider","harperpwd","3700"]}'
INVOKE '{"function":"createUser","Args":["isla@journeyhub.com","Isla","user","islapwd","2100"]}'
INVOKE '{"function":"createUser","Args":["jackson@wanderers.org","Jackson","provider","jacksonpwd","4300"]}'

# --- Update Users ---
INVOKE '{"function":"updateUser","Args":["cameron@travelsphere.com","cameron@travelsphere.com","user","1000"]}'
INVOKE '{"function":"updateUser","Args":["bella@skypass.com","bella@skypass.com","provider","6000"]}'

# --- Create Tickets ---
INVOKE '{"function":"createTicket","Args":["3001","SVC201","1","Bus CityAlpha to CityBeta","bella@skypass.com","40.00","2025-06-12T08:30:00Z","100","CityAlpha","CityBeta","bus"]}'
INVOKE '{"function":"createTicket","Args":["3002","SVC202","1","Flight Tokyo to Seoul","danielle@airhop.com","220.00","2025-06-17T13:00:00Z","270","Tokyo","Seoul","flight"]}'
INVOKE '{"function":"createTicket","Args":["3003","SVC203","1","Shuttle CenterPark to MainStation","elliot@roadconnect.io","18.00","2025-06-19T07:30:00Z","50","CenterPark","MainStation","shuttle"]}'
INVOKE '{"function":"createTicket","Args":["3004","SVC204","1","Bus Suburbia to Downtown","harper@skyexpress.com","38.00","2025-06-21T10:00:00Z","130","Suburbia","Downtown","bus"]}'
INVOKE '{"function":"createTicket","Args":["3005","SVC205","1","Train CentralHub to RiverSide","jackson@wanderers.org","85.00","2025-06-23T06:30:00Z","190","CentralHub","RiverSide","train"]}'
INVOKE '{"function":"createTicket","Args":["3006","SVC206","1","Cruise BayHarbor to CrystalCove","danielle@airhop.com","500.00","2025-07-04T11:00:00Z","480","BayHarbor","CrystalCove","cruise"]}'
INVOKE '{"function":"createTicket","Args":["3007","SVC207","1","Flight London to Dublin","danielle@airhop.com","190.00","2025-07-08T15:00:00Z","290","London","Dublin","flight"]}'
INVOKE '{"function":"createTicket","Args":["3008","SVC208","1","Bus NorthEnd to SouthVille","bella@skypass.com","28.00","2025-07-10T09:00:00Z","85","NorthEnd","SouthVille","bus"]}'
INVOKE '{"function":"createTicket","Args":["3009","SVC209","1","Train MountainView to ValleyTown","jackson@wanderers.org","92.00","2025-07-13T08:00:00Z","210","MountainView","ValleyTown","train"]}'

# --- Book Tickets ---
INVOKE '{"function":"bookTicket","Args":["alex@journeys.io","3001"]}'
INVOKE '{"function":"bookTicket","Args":["fiona@pathfinders.net","3002"]}'
INVOKE '{"function":"bookTicket","Args":["gavin@trailmix.com","3003"]}'
INVOKE '{"function":"bookTicket","Args":["isla@journeyhub.com","3005"]}'
INVOKE '{"function":"bookTicket","Args":["gavin@trailmix.com","3006"]}'
INVOKE '{"function":"bookTicket","Args":["alex@journeys.io","3007"]}'

# --- Cancel Tickets ---
INVOKE '{"function":"cancelTicket","Args":["alex@journeys.io","3001"]}'
INVOKE '{"function":"cancelTicket","Args":["gavin@trailmix.com","3006"]}'

# --- Cancel Listing (fixed 2 args) ---
INVOKE '{"function":"cancelListing","Args":["elliot@roadconnect.io","SVC203"]}'

# --- Delete User ---
INVOKE '{"function":"deleteUser","Args":["fiona@pathfinders.net","fiona@pathfinders.net"]}'

# --- Update Ticket Price (fixed 3 args) ---
INVOKE '{"function":"updateTicketPrice","Args":["3005","jackson@wanderers.org","95.00"]}'

# --- Delete Ticket (fixed 2 args) ---
INVOKE '{"function":"deleteTicket","Args":["3005","jackson@wanderers.org"]}'

# --- Extra Bookings & Cancels ---
INVOKE '{"function":"bookTicket","Args":["isla@journeyhub.com","3004"]}'
INVOKE '{"function":"cancelTicket","Args":["isla@journeyhub.com","3004"]}'


echo "=== Finished TEST ==="
