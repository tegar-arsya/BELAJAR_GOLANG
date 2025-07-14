package entity
import "time"
type Cv struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    File      string `json:"file"` // URL atau path file CV yang diupload
    CreatedAt time.Time `json:"created_at"` // Tanggal pembuatan CV dalam format string
}
