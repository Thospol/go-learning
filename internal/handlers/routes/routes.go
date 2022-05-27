package routes

import (
	ctx "context"
	"fmt"

	"os"
	"os/signal"
	"time"

	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/pkg/blog"
	"github.com/Thospol/go-learning/internal/pkg/mail"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/sirupsen/logrus"
)

const (
	// MaximumSize100MB body limit 100 mb.
	MaximumSize100MB = 1024 * 1024 * 100
	// MaximumSize1MB body limit 1 mb.
	MaximumSize1MB = 1024 * 1024 * 1
)

// NewRouter new router
func NewRouter() {
	app := fiber.New(
		fiber.Config{
			IdleTimeout:    5 * time.Second,
			BodyLimit:      MaximumSize100MB,
			ReadBufferSize: MaximumSize1MB,
		},
	)

	app.Use(
		compress.New(),
		requestid.New(),
		cors.New(),
	)
	api := app.Group("/api")
	v1 := api.Group("/v1")
	if config.CF.Swagger.Enable {
		v1.Get("/swagger/*", swagger.HandlerDefault)
	}
	mailEndpoint := mail.NewEndpoint()
	mail := v1.Group("/mails")
	mail.Post("/send", mailEndpoint.Send)

	blogEndpoint := blog.NewEndpoint()
	blog := v1.Group("/blogs")
	blog.Post("/", blogEndpoint.Create)
	blog.Get("/:id", blogEndpoint.Get)

	api.Use(func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Not Found",
			"status":  404,
		})
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
		defer cancel()

		logrus.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	logrus.Infof("Start server on port: %s ...", config.CF.App.Port)
	err := app.Listen(fmt.Sprintf(":%s", config.CF.App.Port))
	if err != nil {
		logrus.Panic(err)
	}
}
