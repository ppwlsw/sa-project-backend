package entities

type User struct {
	ID           int    `json:"id"`
	CredentialID string `json:"credential_id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Status       string `json:"status"`
	Role         int    `json:"role"`
	// Tier entities.Tier
}
