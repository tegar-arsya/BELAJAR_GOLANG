// models/news.go
package models

import "time"

type News struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title"`
    AuthorName string   `json:"author_name"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
