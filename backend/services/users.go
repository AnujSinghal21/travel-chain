package services

import "github.com/anujSinghal21/travel-chain/api-interface/models"

func GetUserDetails(certificate string) (models.UserStruct, error) {
	// TODO: Implement the logic to fetch user details based on the certificate
	// This is a placeholder implementation
	userDetails, err := verifyCertificateAndFetchDetails(certificate)
	if err != nil {
		// Handle error (e.g., log it, return an empty struct, etc.)
		return models.UserStruct{}, err
	}
	// Return the user details
	return userDetails, nil
}

func verifyCertificateAndFetchDetails(certificate string)  (models.UserStruct, error) {
	// TODO: Implement the logic to verify the certificate and fetch user details
	// This is a placeholder implementation
	// In a real implementation, you would verify the certificate and fetch user details from blockchain
	userDetails := models.UserStruct{
		Name:      "John Doe",
		Certificate: certificate,
		Email: "johndoe@gmail.com",
		Phone: "1234567890",
		Id: "user123",
		Balance: 100,
	}
	// Return the user details
	return userDetails, nil
}