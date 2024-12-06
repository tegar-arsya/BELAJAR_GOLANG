// models/user.go
package models

import "time"
type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"unique" json:"username"`
    Email     string    `json:"email"` // Pastikan ini publik
    Password  string    `json:"password"`            // Pastikan ini publik
    Role      string    `gorm:"default:user" json:"role"` // Default role adalah "user"
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
