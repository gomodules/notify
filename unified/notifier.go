package unified

import (
	"fmt"
	"strings"

	"github.com/2tgroup/go-notify/discord"
	"github.com/2tgroup/go-notify/hipchat"
	"github.com/2tgroup/go-notify/log"
	"github.com/2tgroup/go-notify/mailgun"
	"github.com/2tgroup/go-notify/plivo"
	"github.com/2tgroup/go-notify/pushover"
	"github.com/2tgroup/go-notify/slack"
	"github.com/2tgroup/go-notify/smtp"
	"github.com/2tgroup/go-notify/stride"
	"github.com/2tgroup/go-notify/telegram"
	"github.com/2tgroup/go-notify/twilio"
	"github.com/2tgroup/go-notify/webhook"
)

//NotifyVia we push notify via what
func NotifyVia(via string) (interface{}, error) {
	switch strings.ToLower(via) {
	case plivo.UID:
		return plivo.Default()
	case twilio.UID:
		return twilio.Default()
	case smtp.UID:
		return smtp.Default()
	case mailgun.UID:
		return mailgun.Default()
	case hipchat.UID:
		return hipchat.Default()
	case slack.UID:
		return slack.Default()
	case log.UID:
		return log.Default()
	case webhook.UID:
		return webhook.Default()
	case pushover.UID:
		return pushover.Default()
	case telegram.UID:
		return telegram.Default()
	case discord.UID:
		return discord.Default()
	case stride.UID:
		return stride.Default()
	}
	return nil, fmt.Errorf("unknown notifier %s", via)
}
