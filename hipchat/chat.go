package hipchat

import (
	"github.com/appscode/go-notify"
	"github.com/kelseyhightower/envconfig"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

const Uid = "hipchat"

type Options struct {
	AuthToken string // HIPCHAT_AUTH_TOKEN
}

type client struct {
	opt  Options
	to   []string
	body string
}

var _ notify.ByChat = &client{}

func New(opt Options) *client {
	return &client{opt: opt}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process(Uid, &opt)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func (c *client) WithBody(body string) {
	c.body = body
}

func (c *client) To(to string, cc ...string) {
	c.to = append([]string{to}, cc...)
}

func (c *client) Send() error {
	h := hipchat.NewClient(c.opt.AuthToken)

	for _, room := range c.to {
		_, err := h.Room.Notification(room, &hipchat.NotificationRequest{Message: c.body})
		if err != nil {
			return err
		}
	}
	return nil
}
