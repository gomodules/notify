package go_notifier

type ByEmail interface {
	From(from string) ByEmail
	WithSubject(subject string) ByEmail
	WithBody(body string) ByEmail
	To(to string, tos ...string) ByEmail
	Send() error
	SendHtml() error
}
