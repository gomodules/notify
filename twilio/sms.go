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

type Options struct {
	AccountSid string // TWILIO_ACCOUNT_SID
	AuthToken  string // TWILIO_ACCOUNT_TOKEN
}

type client struct {
	opt Options
	v   url.Values
}

var _ notify.BySMS = &client{}

func New(opt Options) *client {
	return &client{
		opt: opt,
		v:   url.Values{},
	}
}

func Default() (*client, error) {
	var opt Options
	err := envconfig.Process("twilio", &opt)
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

func (c *client) To(to string) {
	c.v.Set("To", to)
}

func (c *client) Send() error {
	h := &http.Client{Timeout: time.Second * 10}

	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", c.opt.AccountSid)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(c.v.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.opt.AccountSid, c.opt.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = h.Do(req)
	return err
}
