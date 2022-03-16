package slack

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"gomodules.xyz/envconfig"
	"gomodules.xyz/notify"
)

const UID = "slack"

type Options struct {
	AuthToken string   `envconfig:"AUTH_TOKEN" required:"true"`
	Channel   []string `envconfig:"CHANNEL"`
	ProxyHost string   `envconfig:"PROXY_HOST"`
	ProxyPort int      `envconfig:"PROXY_PORT"`
}

type client struct {
	opt       Options
	body      string
	msg       slack.MsgOption
	clientOpt *http.Client
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

func (c client) WithMsgOption(msgOption slack.MsgOption) notify.ByChat {
	c.msg = msgOption
	return &c
}

func (c client) WithProxy(proxyHost string, proxyPort int) notify.ByChat {
	proxyURL, err := url.Parse(fmt.Sprintf("%v:%v", proxyHost, proxyPort))
	if err != nil {
		log.Error().Err(err).Msg("failed to parse URL")
	}
	c.clientOpt = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	return &c
}

func (c client) To(to string, cc ...string) notify.ByChat {
	c.opt.Channel = append([]string{to}, cc...)
	return &c
}

func (c *client) Send() error {
	if len(c.opt.Channel) == 0 {
		return errors.New("missing to")
	}

	s := slack.New(c.opt.AuthToken)
	for _, channel := range c.opt.Channel {
		if _, _, err := s.PostMessageContext(
			context.TODO(),
			channel,
			slack.MsgOptionText(c.body, false)); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) SendAsBlockMessage() error {
	if len(c.opt.Channel) == 0 {
		return errors.New("missing to")
	}
	s := slack.New(c.opt.AuthToken, slack.OptionHTTPClient(c.clientOpt))

	for _, channel := range c.opt.Channel {
		if _, _, err := s.PostMessageContext(
			context.TODO(),
			channel,
			c.msg); err != nil {
			return err
		}
	}
	return nil
}
