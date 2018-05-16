package stride

import (
	"errors"

	"bitbucket.org/atlassian/go-stride/pkg/stride"
	"github.com/appscode/envconfig"
	"github.com/appscode/go-notify"
)

const (
	UID = "stride"
)

type Options struct {
	CloudID      string   `envconfig:"CLOUD_ID" required:"true"`
	ClientID     string   `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string   `envconfig:"CLIENT_SECRET" required:"true"`
	To           []string `envconfig:"TO"`
}

type client struct {
	opt  Options
	body string
}

var _ notify.ByChat = &client{}

func New(opt Options) *client {
	return &client{opt: opt}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process(UID, &opt)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func Load(loader envconfig.LoaderFunc) (*client, error) {
	var opt Options
	err := envconfig.Load(UID, &opt, loader)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func (c client) UID() string {
	return UID
}

func (c client) WithBody(body string) notify.ByChat {
	c.body = body
	return &c
}

func (c client) To(to string, cc ...string) notify.ByChat {
	c.opt.To = append([]string{to}, cc...)
	return &c
}

func (c *client) Send() error {
	if len(c.opt.To) == 0 {
		return errors.New("missing to")
	}

	s := stride.New(c.opt.ClientID, c.opt.ClientSecret)

	for _, to := range c.opt.To {
		conversation, err := s.GetConversationByName(c.opt.CloudID, to)
		if err != nil {
			return err
		}
		if err := stride.SendText(s, conversation.CloudID, conversation.ID, c.body); err != nil {
			return err
		}
	}
	return nil
}
