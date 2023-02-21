package categories

import (
	"kedaiprogrammer/businesses"
	"time"
)

type Category struct {
	ID            int
	Category_name string
	Business_id   int
	Business      businesses.Business
	Created_at    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Updated_at    time.Time
}
