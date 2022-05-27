package blog

import (
	"github.com/Thospol/go-learning/internal/repositories"
	"gorm.io/gorm"
)

// Repository repository
type Repository interface {
	Create(db *gorm.DB, i interface{}) error
	FindOneObjectByID(db *gorm.DB, id uint, i interface{}) error
}

type repository struct {
	repositories.Repository
}

// NewRepository new repository
func NewRepository() Repository {
	return &repository{
		repositories.NewRepository(),
	}
}
