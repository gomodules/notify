package eventer

import (
	"strings"

	notify "gomodules.xyz/notify"
	"gomodules.xyz/notify/unified"
)

type (
	//EventForwarder hold to custom loadFunc
	EventForwarder struct{}
)

//This variable can be load from config
var (
	sendTo = ""
	sendCC = []string{}
)

//Notify send notify
func (f *EventForwarder) Notify(emailSub, chatSub, body, notifyBy string) error {

	notifier, err := unified.NotifyVia(strings.ToLower(notifyBy))

	if err != nil {
		return err
	}
	switch n := notifier.(type) {

	case notify.ByEmail:
		return n.
			To(sendTo, sendCC...).
			WithSubject(emailSub).
			WithBody(body).
			WithNoTracking().
			Send()
	case notify.BySMS:
		return n.To(sendTo, sendCC...).
			WithBody(emailSub).
			Send()
	case notify.ByChat:
		return n.To(sendTo, sendCC...).
			WithBody(chatSub).
			Send()
	case notify.ByPush:
		return n.To(sendCC...).
			WithBody(chatSub).
			Send()
	}
	return nil
}
