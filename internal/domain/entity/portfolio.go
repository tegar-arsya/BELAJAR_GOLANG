package entity
import "time"
type Portfolio struct {
    ID          uint     `json:"id"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Site        string   `json:"site"`
    GithubUrl   string   `json:"github_url"`
    ImageUrl   string   `json:"image_url"` // URL gambar yang diupload
    CreatedAt   time.Time `json:"created_at"`
}
