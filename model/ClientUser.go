package model

// ClientUser is a join table
type ClientUser struct {
	ClientID uint `gorm:"primary_key" sql:"not null"`
	Client   *Client
	UserID   uint `gorm:"primary_key" sql:"not null"`
	User     *User
	Role     ClientRole `sql:"type:client_role;not null"`
}
