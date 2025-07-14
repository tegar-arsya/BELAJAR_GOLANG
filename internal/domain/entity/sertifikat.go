package entity
import "time"
type Sertifikat struct {
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Site        string `json:"site"`
    ImageUrl    string `json:"image_url"`
    CreatedAt   time.Time `json:"created_at"`
}
