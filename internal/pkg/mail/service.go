package mail

import (
	"os"

	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/core/context"
	"github.com/Thospol/go-learning/internal/core/mail"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	Send(c *context.Context) error
}

type service struct {
	config *config.Configs
	mail   mail.Mailer
}

// NewService new service
func NewService() Service {
	return &service{
		config: config.CF,
		mail: mail.New(mail.Configuration{
			Host:     config.CF.SMTP.Host,
			Port:     config.CF.SMTP.Port,
			Username: config.CF.SMTP.Username,
			Password: config.CF.SMTP.Password,
		},
		),
	}
}

// Send send
func (s *service) Send(c *context.Context) error {
	b, err := os.ReadFile("./assets/mail.html")
	if err != nil {
		logrus.Errorf("read file error: %s", err)
		return err
	}
	err = s.mail.
		From(config.CF.SMTP.Sender, config.CF.SMTP.SenderAlias).
		To("isocare.thospol@gmail.com").
		Subject("ทดสอบการส่งอีเมล (Try to do!!!)").
		HTML(string(b)).
		Send()
	if err != nil {
		logrus.Errorf("send mail error: %s", err)
		return err
	}

	return nil
}
