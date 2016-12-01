package smtp

import (
	"github.com/appscode/go-notify"
	h2t "github.com/jaytaylor/html2text"
	gomail "gopkg.in/gomail.v2"
)

type Options struct {
	Host               string
	Port               int
	InsecureSkipVerify bool
	Username, Password string
}

type mailer struct {
	opt  Options
	mail *gomail.Message
	body string
	html bool
}

var _ notify.ByEmail = &mailer{}

func New(opt Options) *mailer {
	return &mailer{
		opt:  opt,
		mail: gomail.NewMessage(),
	}
}

func (m *mailer) From(from string) {
	m.mail.SetHeader("From", from)
}

func (m *mailer) WithSubject(subject string) {
	m.mail.SetHeader("Subject", subject)
}
func (m *mailer) WithBody(body string) {
	m.body = body
}

func (m *mailer) To(to string, cc ...string) {
	tos := append([]string{to}, cc...)
	m.mail.SetHeader("To", tos...)
}

func (m *mailer) Send() error {
	if m.html {
		m.mail.SetBody("text/html", m.body)
		if t, err := h2t.FromString(m.body); err == nil {
			m.mail.AddAlternative("text/plain", t)
		}
	} else {
		m.mail.SetBody("text/plain", m.body)
	}

	if m.opt.Username == "" && m.opt.Password == "" {
		d := gomail.NewDialer(m.opt.Host, m.opt.Port, m.opt.Username, m.opt.Password)
		return d.DialAndSend(m.mail)
	} else {
		d := gomail.Dialer{Host: m.opt.Host, Port: m.opt.Port}
		return d.DialAndSend(m.mail)
	}
	return nil
}

func (m *mailer) SendHtml() error {
	m.html = true
	return m.Send()
}
