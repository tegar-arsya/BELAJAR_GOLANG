package models

import "time"

type Portfolio struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    ImageUrl    string    `json:"image_url"`
    Site        string    `json:"site"` // URL atau nama situs tempat portofolio ditampilkan
    GithubUrl  string    `json:"github_url"` // URL ke repositori GitHub jika ada
    CreatedAt   time.Time `json:"created_at"`
}


