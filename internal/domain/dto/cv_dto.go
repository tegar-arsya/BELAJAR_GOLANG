package dto

type   CreateCvRequest struct {
    File string `form:"file" binding:"required"` // URL atau path file CV yang diupload
}
