package blog

import (
	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/handlers"
	"github.com/Thospol/go-learning/internal/request"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
}

type endpoint struct {
	config  *config.Configs
	result  *config.ReturnResult
	service Service
}

// NewEndpoint new migrate endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config:  config.CF,
		result:  config.RR,
		service: NewService(),
	}
}

// Create create
// @Tags Blog
// @Summary Create
// @Description Create
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body CreateRequest true "request body"
// @Success 200 {object} models.Blog
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Security ApiKeyAuth
// @Router /blogs [post]
func (ep *endpoint) Create(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.Create, &CreateRequest{})
}

// Get get
// @Tags Blog
// @Summary Get
// @Description Get
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param id path string true "id"
// @Success 200 {object} models.Blog
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /blogs/{id} [get]
func (ep *endpoint) Get(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.Get, &request.GetOne{})
}
