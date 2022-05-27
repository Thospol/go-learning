package mail

import (
	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	Send(c *fiber.Ctx) error
}

type endpoint struct {
	config  *config.Configs
	result  *config.ReturnResult
	service Service
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config:  config.CF,
		result:  config.RR,
		service: NewService(),
	}
}

// Send send
// @Tags Mail
// @Summary Send
// @Description Send
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param Level header string true "(hr_officer, hr_manager, recruite, admin...)" default(admin)
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Security ApiKeyAuth
// @Router /mails/send [post]
func (ep *endpoint) Send(c *fiber.Ctx) error {
	return handlers.ResponseSuccessWithoutRequest(c, ep.service.Send)
}
