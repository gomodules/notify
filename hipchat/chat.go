package hipchat

import (
	"github.com/appscode/go-notify"
	"github.com/kelseyhightower/envconfig"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

type Options struct {
	AuthToken string // HIPCHAT_ACCOUNT_TOKEN
}

type client struct {
	opt  Options
	to   string
	body string
}

var _ notify.ByChat = &client{}

func New(opt Options) *client {
	return &client{opt: opt}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process("hipchat", &opt)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func (c *client) WithBody(body string) {
	c.body = body
}

func (c *client) To(to string) {
	c.to = to
}

func (c *client) Send() error {
	h := hipchat.NewClient(c.opt.AuthToken)

	_, err := h.Room.Notification(c.to, &hipchat.NotificationRequest{Message: c.body})
	return err
}
