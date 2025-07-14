package models

import "time"

type Cv struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    File       string    `json:"file"`
    CreatedAt   time.Time `json:"created_at"`
}


