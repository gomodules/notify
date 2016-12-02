package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/appscode/go-notify"
	"github.com/kelseyhightower/envconfig"
)

const Uid = "twilio"

type Options struct {
	AccountSid string // TWILIO_ACCOUNT_SID
	AuthToken  string // TWILIO_AUTH_TOKEN
	From       string // TWILIO_FROM
}

type client struct {
	opt Options
	v   url.Values
	to  []string
}

var _ notify.BySMS = &client{}

func New(opt Options) *client {
	v := url.Values{}
	v.Set("From", opt.From)
	return &client{
		opt: opt,
		v:   v,
	}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process(Uid, &opt)
	if err != nil {
		return nil, err
	}
	return New(opt), nil
}

func (c *client) From(from string) {
	c.v.Set("From", from)
}

func (c *client) WithBody(body string) {
	c.v.Set("Body", body)
}

func (c *client) To(to string, cc ...string) {
	c.to = append([]string{to}, cc...)
}

func (c *client) Send() error {
	h := &http.Client{Timeout: time.Second * 10}

	for _, receiver := range c.to {
		c.v.Set("To", receiver)
		urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", c.opt.AccountSid)
		req, err := http.NewRequest("POST", urlStr, strings.NewReader(c.v.Encode()))
		if err != nil {
			return err
		}

		req.SetBasicAuth(c.opt.AccountSid, c.opt.AuthToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		_, err = h.Do(req)
		if err != nil {
			return err
		}
	}
	return nil
}
