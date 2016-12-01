package notify

type ByEmail interface {
	From(from string)
	WithSubject(subject string)
	WithBody(body string)
	To(to string, tos ...string)
	Send() error
	SendHtml() error
}
