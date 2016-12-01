package mailgun

import (
	notify "github.com/appscode/go-notify"
	"github.com/appscode/log"
	h2t "github.com/jaytaylor/html2text"
	mailgun "github.com/mailgun/mailgun-go"
)

type Options struct {
	Domain       string
	ApiKey       string
	PublicApiKey string
}

type mailer struct {
	to      []string
	from    string
	subject string
	body    string
	html    bool
	tag     string

	mg mailgun.Mailgun
}

var _ notify.ByEmail = &mailer{}

func New(opt Options) *mailer {
	return &mailer{
		mg: mailgun.NewMailgun(opt.Domain, opt.ApiKey, opt.PublicApiKey),
	}
}

func (m *mailer) From(from string) {
	m.from = from
}

func (m *mailer) WithSubject(subject string) {
	m.subject = subject
}
func (m *mailer) WithBody(body string) {
	m.body = body
}

func (m *mailer) To(to string, cc ...string) {
	m.to = append([]string{to}, cc...)
}

func (m *mailer) Send() error {
	text := m.body
	if m.html {
		if t, err := h2t.FromString(m.body); err == nil {
			text = t
		}
	}
	msg := m.mg.NewMessage(m.from, m.subject, text, m.to...)
	if m.html {
		msg.SetHtml(m.body)
	}
	if m.tag != "" {
		msg.AddTag(m.tag)
	}
	msg.SetTracking(true)
	msg.SetTrackingClicks(true)
	msg.SetTrackingOpens(true)
	response, id, err := m.mg.Send(msg)
	log.Infof("Mailgun server response[%v]: %v\n", id, response)
	if err != nil {
		log.Errorln("[Mailer] failed to send mail")
		log.V(10).Infoln("[Mailer] mail", m)
		return err
	}
	return nil
}

func (m *mailer) SendHtml() error {
	m.html = true
	return m.Send()
}
