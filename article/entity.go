package article

import "time"

type Posts struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"size:200"`
	Content   string
	Category  string `gorm:"size:200"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string `gorm:"size:200"`
}
