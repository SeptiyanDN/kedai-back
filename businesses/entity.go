package businesses

import "time"

type Business struct {
	ID            int
	Business_name string
	Created_at    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Updated_at    time.Time
}
