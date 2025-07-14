package entity

import "time"

type Pengalaman struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    ImageUrl    string    `json:"image_url"` // URL gambar yang diupload
    CreatedAt   time.Time `json:"created_at"`
}
