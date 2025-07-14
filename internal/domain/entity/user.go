package entity
import "time"
type User struct {
	 ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"unique" json:"username"`
    Email     string    `json:"email"` // Pastikan ini publik
    Password  string    `json:"password"`            // Pastikan ini publik
    CreatedAt time.Time `json:"created_at"`
}

