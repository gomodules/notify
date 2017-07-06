package notify

type ByEmail interface {
	From(from string) ByEmail
	WithSubject(subject string) ByEmail
	WithTag(tag string) ByEmail
	To(to string, cc ...string) ByEmail
	AsHtml() ByEmail
	Message
}

type BySMS interface {
	From(from string) BySMS
	To(to string, cc ...string) BySMS
	Message
}

type ByChat interface {
	To(to string, cc ...string) ByChat
	Message
}

type Message interface {
	WithBody(body string) Message
	Send() error
}
