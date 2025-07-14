package dto

type CreatePengalamanRequest struct {
    Title       string   `form:"title" binding:"required"`
    Description string   `form:"description" binding:"required"`
    Site        string   `form:"site" binding:"required"`
    ImageUrls   []string // diproses dari file upload
}

