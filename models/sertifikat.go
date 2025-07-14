package models

import "time"

type Sertifikat struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    ImageUrl    string    `json:"image_url"`
    Site        string    `json:"site"` // URL atau nama situs tempat sertifikat diterbitkan
    CreatedAt   time.Time `json:"created_at"`
}


