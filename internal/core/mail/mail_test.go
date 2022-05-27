package mail

import (
	"os"
	"testing"

	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/stretchr/testify/require"
)

func init() {
	// Init configuration
	err := config.InitConfig("../../../configs")
	if err != nil {
		panic(err)
	}
}

func TestMailWithHTML(t *testing.T) {
	c := Configuration{
		Host:     config.CF.SMTP.Host,
		Port:     config.CF.SMTP.Port,
		Username: config.CF.SMTP.Username,
		Password: config.CF.SMTP.Password,
	}
	b, _ := os.ReadFile("mail.html")
	mail := New(c)
	err := mail.
		From(config.CF.SMTP.Sender, config.CF.SMTP.SenderAlias).
		To("isocare.thospol@gmail.com").
		Subject("ทดสอบการส่งอีเมล").
		HTML(string(b)).
		Send()
	require.NoError(t, err)
}
