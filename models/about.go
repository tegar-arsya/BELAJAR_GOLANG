package models

import "time"


type About struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    ImageUrl    string    `json:"image_url"`
    CreatedAt   time.Time `json:"created_at"`
}
