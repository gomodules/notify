package notify17

import (
	"bytes"
	"errors"
	"fmt"
	"gomodules.xyz/envconfig"
	"gomodules.xyz/notify"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Notify17APIURL = "https://hook.notify17.net/api/raw"
	UID            = "notify17"
)

// Options allows full configuration of the message sent to the Notify17 API
type Options struct {
	Token   string `envconfig:"TOKEN" required:"true"`
	Message string `envconfig:"MESSAGE"`

	// Optional params
	Title string `envconfig:"TITLE"`
}

type client struct {
	opt Options
}

var _ notify.ByPush = &client{}

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

func (c client) WithBody(body string) notify.ByPush {
	c.opt.Message = body
	return &c
}

func (c client) To(to ...string) notify.ByPush {
	return &c
}

func (c *client) Send() error {
	if c.opt.Token == "" {
		return errors.New("Missing token")
	}

	if c.opt.Message == "" {
		return errors.New("Missing message")
	}

	if c.opt.Title == "" {
		c.opt.Title = "New notification"
	}

	msg := url.Values{}
	msg.Set("title", c.opt.Title)
	msg.Set("content", c.opt.Message)
	buf := bytes.NewBufferString(msg.Encode())

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", Notify17APIURL, c.opt.Token), buf)
	if err != nil {
		return err
	}

	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	return nil
}
