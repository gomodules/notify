package smtp

import (
	"crypto/tls"

	"github.com/appscode/go-notify"
	h2t "github.com/jaytaylor/html2text"
	"github.com/kelseyhightower/envconfig"
	gomail "gopkg.in/gomail.v2"
)

type Options struct {
	Host               string // SMTP_HOST
	Port               int    // SMTP_PORT
	InsecureSkipVerify bool   // SMTP_INSECURE_SKIP_VERIFY
	Username           string // SMTP_USERNAME
	Password           string // SMTP_PASSWORD
}

type client struct {
	opt  Options
	mail *gomail.Message
	body string
	html bool
}

var _ notify.ByEmail = &client{}

func New(opt Options) *client {
	return &client{
		opt:  opt,
		mail: gomail.NewMessage(),
	}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process("smtp", &opt)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func (c *client) From(from string) {
	c.mail.SetHeader("From", from)
}

func (c *client) WithSubject(subject string) {
	c.mail.SetHeader("Subject", subject)
}
func (c *client) WithBody(body string) {
	c.body = body
}

func (c *client) To(to string, cc ...string) {
	tos := append([]string{to}, cc...)
	c.mail.SetHeader("To", tos...)
}

func (c *client) Send() error {
	if c.html {
		c.mail.SetBody("text/html", c.body)
		if t, err := h2t.FromString(c.body); err == nil {
			c.mail.AddAlternative("text/plain", t)
		}
	} else {
		c.mail.SetBody("text/plain", c.body)
	}

	var d *gomail.Dialer
	if c.opt.Username == "" && c.opt.Password == "" {
		d = gomail.NewDialer(c.opt.Host, c.opt.Port, c.opt.Username, c.opt.Password)
	} else {
		d = &gomail.Dialer{Host: c.opt.Host, Port: c.opt.Port}
	}
	if c.opt.InsecureSkipVerify {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return d.DialAndSend(c.mail)
}

func (c *client) SendHtml() error {
	c.html = true
	return c.Send()
}
