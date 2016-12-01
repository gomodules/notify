package notify

type ByEmail interface {
	From(from string)
	WithSubject(subject string)
	WithBody(body string)
	To(to string, cc ...string)
	Send() error
	SendHtml() error
}

type BySMS interface {
	From(from string)
	WithBody(body string)
	To(to string, cc ...string)
	Send() error
}

type ByChat interface {
	WithBody(body string)
	To(to string, cc ...string)
	Send() error
}
