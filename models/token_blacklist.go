// models/token_blacklist.go
package models

import (
	"time"
)

// TokenBlacklist adalah struktur untuk menyimpan token yang diblacklist
type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`  // Token yang diblacklist
	ExpiresAt time.Time `json:"expires_at"` // Waktu kadaluarsa token
	CreatedAt time.Time `json:"created_at"`  // Waktu token ditambahkan ke blacklist
	UpdatedAt time.Time `json:"updated_at"`  // Waktu token terakhir diupdate
}
