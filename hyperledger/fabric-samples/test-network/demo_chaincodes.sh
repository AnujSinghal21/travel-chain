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
INVOKE '{"function":"createUser","Args":["mia@voyage.com","Mia","user","miapwd","1200"]}'
INVOKE '{"function":"createUser","Args":["noah@routes.net","Noah","provider","noahpwd","6000"]}'
INVOKE '{"function":"createUser","Args":["oliver@travelo.org","Oliver","user","oliverpwd","800"]}'
INVOKE '{"function":"createUser","Args":["piper@flyhigh.com","Piper","provider","piperpwd","4000"]}'
INVOKE '{"function":"createUser","Args":["quinn@speedshuttle.io","Quinn","provider","quinnpwd","5000"]}'
INVOKE '{"function":"createUser","Args":["riley@transithub.com","Riley","user","rileypwd","1500"]}'
INVOKE '{"function":"createUser","Args":["sam@adventureco.com","Sam","user","sampwd","1800"]}'
INVOKE '{"function":"createUser","Args":["taylor@fastbuses.com","Taylor","provider","taylorpwd","3500"]}'
INVOKE '{"function":"createUser","Args":["uma@voyagers.net","Uma","user","umapwd","2200"]}'
INVOKE '{"function":"createUser","Args":["victor@nomadshub.com","Victor","provider","victorpwd","4200"]}'

# --- Update Users ---
INVOKE '{"function":"updateUser","Args":["oliver@travelo.org","oliver@travelo.org","user","850"]}'
INVOKE '{"function":"updateUser","Args":["noah@routes.net","noah@routes.net","provider","6200"]}'

# --- Create Tickets ---
INVOKE '{"function":"createTicket","Args":["2001","SVC101","1","Bus MetroCity to OceanView","noah@routes.net","35.00","2025-06-11T08:00:00Z","110","MetroCity","OceanView","bus"]}'
INVOKE '{"function":"createTicket","Args":["2002","SVC102","1","Flight Paris to Rome","piper@flyhigh.com","180.00","2025-06-15T14:00:00Z","300","Paris","Rome","flight"]}'
INVOKE '{"function":"createTicket","Args":["2003","SVC103","1","Shuttle Uptown to Midtown","quinn@speedshuttle.io","20.00","2025-06-18T09:00:00Z","60","Uptown","Midtown","shuttle"]}'
INVOKE '{"function":"createTicket","Args":["2004","SVC104","1","Bus EastEnd to WestSide","taylor@fastbuses.com","45.00","2025-06-20T10:30:00Z","140","EastEnd","WestSide","bus"]}'
INVOKE '{"function":"createTicket","Args":["2005","SVC105","1","Train GrandCentral to EastBay","victor@nomadshub.com","90.00","2025-06-22T07:00:00Z","200","GrandCentral","EastBay","train"]}'
INVOKE '{"function":"createTicket","Args":["2006","SVC106","1","Cruise PortCity to SunnyIsle","piper@flyhigh.com","550.00","2025-07-03T12:00:00Z","600","PortCity","SunnyIsle","cruise"]}'

# --- Book Tickets ---
INVOKE '{"function":"bookTicket","Args":["mia@voyage.com","2001"]}'
INVOKE '{"function":"bookTicket","Args":["riley@transithub.com","2002"]}'
INVOKE '{"function":"bookTicket","Args":["sam@adventureco.com","2003"]}'
INVOKE '{"function":"bookTicket","Args":["uma@voyagers.net","2005"]}'
INVOKE '{"function":"bookTicket","Args":["sam@adventureco.com","2006"]}'

# --- Cancel Tickets ---
INVOKE '{"function":"cancelTicket","Args":["mia@voyage.com","2001"]}'
INVOKE '{"function":"cancelTicket","Args":["sam@adventureco.com","2006"]}'

# --- Cancel Listing (fixed 2 args) ---
INVOKE '{"function":"cancelListing","Args":["quinn@speedshuttle.io","SVC103"]}'

# --- Delete User (expect failure if user still booked tickets) ---
INVOKE '{"function":"deleteUser","Args":["riley@transithub.com","riley@transithub.com"]}'

# --- Update Ticket Price (fixed 3 args) ---
INVOKE '{"function":"updateTicketPrice","Args":["2005","victor@nomadshub.com","100.00"]}'

# --- Delete Ticket (fixed 2 args) ---
INVOKE '{"function":"deleteTicket","Args":["2005","victor@nomadshub.com"]}'

# --- Extra Bookings & Cancels ---
INVOKE '{"function":"bookTicket","Args":["uma@voyagers.net","2004"]}'
INVOKE '{"function":"cancelTicket","Args":["uma@voyagers.net","2004"]}'

echo "=== Finished for DEMO ==="
