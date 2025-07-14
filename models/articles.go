
package models
import (
    "time"
)

type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Thumbnail string    `json:"thumbnail"`
	Content   string    `json:"content"` // HTML content
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

