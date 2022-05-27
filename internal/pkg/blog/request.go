package blog

import "github.com/Thospol/go-learning/internal/models"

// CreateRequest create request
type CreateRequest struct {
	Title   string        `json:"title"`
	Message string        `json:"message"`
	Author  models.Author `json:"author"`
}
