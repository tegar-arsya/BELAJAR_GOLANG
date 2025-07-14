package dto

type CreatePortofolioRequest struct {
    Title       string   `form:"title" binding:"required"`
    Description string   `form:"description" binding:"required"`
    Site        string   `form:"site" binding:"required"`
    GithubUrl   string   `form:"github_url" binding:"required"`
    ImageUrls   []string // diproses dari file upload
}
