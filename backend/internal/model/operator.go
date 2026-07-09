package model

type Operator struct {
	ID           int       `json:"id"`
	Callsign     string    `json:"callsign" binding:"required,min=3,max=50"`
	PasswordHash string    `json:"-" `
}
