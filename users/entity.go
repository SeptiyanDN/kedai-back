package users

import "time"

type User struct {
	ID         int
	Uuid       string
	Username   string
	Email      string
	Password   string
	Token      string
	Created_at time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Updated_at time.Time
}
