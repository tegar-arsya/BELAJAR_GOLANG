package dto

type CreateAboutRequest struct {
    Title       string `form:"title" binding:"required"`
    Description string `form:"description" binding:"required"`
    ImageUrl    string `form:"image_url" binding:"required"` // URL atau path gambar yang diupload
}
