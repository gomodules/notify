package tiniyo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gomodules.xyz/notify"
	"net/http"
	"time"
)

func (c vClient) UID() string {
	return UID
}

func (c vClient) WithBody(body string) notify.ByVoice {
	c.body = body
	return &c
}

func (c vClient) From(from string) notify.ByVoice {
	c.opt.From = from
	return &c
}

func (c vClient) To(to string, cc ...string) notify.ByVoice {
	c.opt.To = append([]string{to}, cc...)
	return &c
}
func (c *vClient) Send() error {
	if len(c.opt.To) == 0 {
		return errors.New("Missing to")
	}

	hc := &http.Client{Timeout: time.Second * 10}

	url := fmt.Sprintf(urlTemplate, c.opt.AuthID, "Call")
	params := struct {
		Src  string `json:"to,omitempty"`
		Dst  string `json:"from,omitempty"`
		Text string `json:"speak,omitempty"`
	}{
		c.opt.From,
		"",
		c.body,
	}
	for _, dst := range c.opt.To {
		params.Dst = dst
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(params); err != nil {
			return err
		}

		req, err := http.NewRequest("POST", url, buf)
		if err != nil {
			return err
		}

		req.SetBasicAuth(c.opt.AuthID, c.opt.AuthToken)
		req.Header.Add("Content-Type", "application/json")

		resp, err := hc.Do(req)
		if err != nil {
			return err
		}

		respBody := struct {
			Error string `json:"error"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			return err
		}
		if respBody.Error != "" {
			return errors.New(respBody.Error)
		}

		resp.Body.Close()
	}
	return nil
}
