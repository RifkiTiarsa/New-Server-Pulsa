package entity

import "time"

type Member struct {
	ID        int       `json:"id"`
	MemberID  string    `json:"member_id,omitempty"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Balance   float64   `json:"balance,omitempty"`
	Pin       string    `json:"pin,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
