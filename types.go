package notify

type ByEmail interface {
	UID() string
	From(from string) ByEmail
	WithSubject(subject string) ByEmail
	WithBody(body string) ByEmail
	WithTag(tag string) ByEmail
	WithNoTracking() ByEmail
	To(to string, cc ...string) ByEmail
	Send() error
	SendHtml() error
}

type BySMS interface {
	UID() string
	From(from string) BySMS
	WithBody(body string) BySMS
	To(to string, cc ...string) BySMS
	Send() error
}
type ByVoice interface {
	UID() string
	From(from string) ByVoice
	WithBody(body string) ByVoice
	To(to string, cc ...string) ByVoice
	Send() error
}
type ByChat interface {
	UID() string
	WithBody(body string) ByChat
	To(to string, cc ...string) ByChat
	Send() error
}

type ByPush interface {
	UID() string
	WithBody(body string) ByPush
	To(to ...string) ByPush
	Send() error
}
