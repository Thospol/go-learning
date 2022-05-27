package blog

import (
	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/core/context"
	"github.com/Thospol/go-learning/internal/models"
	"github.com/Thospol/go-learning/internal/request"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	Create(c *context.Context, request *CreateRequest) (*models.Blog, error)
	Get(c *context.Context, request *request.GetOne) (*models.Blog, error)
}

type service struct {
	config     *config.Configs
	result     *config.ReturnResult
	repository Repository
}

// NewService new service
func NewService() Service {
	return &service{
		config:     config.CF,
		result:     config.RR,
		repository: NewRepository(),
	}
}

// Create create
func (s *service) Create(c *context.Context, request *CreateRequest) (*models.Blog, error) {
	blog := &models.Blog{}
	_ = copier.Copy(blog, request)
	err := s.repository.Create(c.GetDatabase(), blog)
	if err != nil {
		logrus.Errorf("create error: %s", err)
		return nil, err
	}

	return blog, nil
}

// Get get
func (s *service) Get(c *context.Context, request *request.GetOne) (*models.Blog, error) {
	blog := &models.Blog{}
	err := s.repository.FindOneObjectByID(c.GetDatabase(), request.ID, blog)
	if err != nil {
		logrus.Errorf("find error: %s", err)
		return nil, err
	}

	return blog, nil
}
