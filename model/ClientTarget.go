package model

// ClientTarget is a join table
type ClientTarget struct {
	ClientID uint `gorm:"primary_key" sql:"not null"`
	Client   *Client
	TargetID uint `gorm:"primary_key" sql:"not null"`
	Target   *Target
}
