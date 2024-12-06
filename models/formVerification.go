package models

import "time"

type FormVerification struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    FormID      uint      `json:"form_id"`
    VerifiedBy  string    `json:"verified_by"`
    VerifiedAt  time.Time `json:"verified_at"`
}