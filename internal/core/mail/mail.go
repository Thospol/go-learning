package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// Mailer mailer
type Mailer interface {
	Subject(subject string) *mail
	From(from string, alias ...string) *mail
	To(to ...string) *mail
	Bcc(cc ...string) *mail
	Cc(cc ...string) *mail
	Text(text string) *mail
	HTML(html string) *mail
	Send() error
}

// Configuration configuration setting
type Configuration struct {
	Host     string
	Port     int
	Username string
	Password string
}

type mail struct {
	e *email.Email
	c Configuration
}

// New new
func New(config Configuration) Mailer {
	return &mail{
		e: email.NewEmail(),
		c: config,
	}
}

// Subject subject
func (mail *mail) Subject(subject string) *mail {
	mail.e.Subject = subject
	return mail
}

// From from
func (mail *mail) From(from string, alias ...string) *mail {
	v := from
	if alias != nil {
		v = fmt.Sprintf("%s <%s>", alias[0], from)
	}
	mail.e.From = v
	return mail
}

/// To to
func (mail *mail) To(to ...string) *mail {
	mail.e.To = append(mail.e.To, Uniq(to)...)
	return mail
}

// Bcc
func (mail *mail) Bcc(cc ...string) *mail {
	mail.e.Bcc = append(mail.e.Bcc, Uniq(cc)...)
	return mail
}

// Cc cc
func (mail *mail) Cc(cc ...string) *mail {
	mail.e.Cc = append(mail.e.Cc, Uniq(cc)...)
	return mail
}

// Text text
func (mail *mail) Text(text string) *mail {
	mail.e.Text = []byte(text)
	return mail
}

// HTML html
func (mail *mail) HTML(html string) *mail {
	mail.e.HTML = []byte(html)
	return mail
}

// Send send
func (mail *mail) Send() error {
	addr := fmt.Sprintf("%s:%d", mail.c.Host, mail.c.Port)
	auth := smtp.PlainAuth("", mail.c.Username, mail.c.Password, mail.c.Host)
	t := &tls.Config{InsecureSkipVerify: true, ServerName: addr}
	return mail.e.SendWithStartTLS(addr, auth, t)
}

// Uniq unique
func Uniq(emails []string) []string {
	encountered := map[string]bool{}
	uniqEmails := []string{}
	for v := range emails {
		if !encountered[emails[v]] {
			encountered[emails[v]] = true
			uniqEmails = append(uniqEmails, emails[v])
		}
	}
	return uniqEmails
}
