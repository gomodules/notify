package unified

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gomodules.xyz/envconfig"
	"gomodules.xyz/notify/discord"
	"gomodules.xyz/notify/log"
	"gomodules.xyz/notify/mailgun"
	"gomodules.xyz/notify/mattermost"
	"gomodules.xyz/notify/plivo"
	"gomodules.xyz/notify/pushover"
	"gomodules.xyz/notify/slack"
	"gomodules.xyz/notify/smtp"
	"gomodules.xyz/notify/telegram"
	"gomodules.xyz/notify/twilio"
	"gomodules.xyz/notify/webhook"
)

const (
	NotifyVia = "NOTIFY_VIA"
)

func Default() (interface{}, error) {
	via, ok := os.LookupEnv(NotifyVia)
	if !ok {
		return nil, errors.New(`"NOTIFY_VIA" is not set.`)
	}
	return DefaultVia(via)
}

func DefaultVia(via string) (interface{}, error) {
	switch strings.ToLower(via) {
	case discord.UID:
		return discord.Default()
	case log.UID:
		return log.Default()
	case mailgun.UID:
		return mailgun.Default()
	case mattermost.UID:
		return mattermost.Default()
	case plivo.UID:
		return plivo.Default()
	case pushover.UID:
		return pushover.Default()
	case slack.UID:
		return slack.Default()
	case smtp.UID:
		return smtp.Default()
	case telegram.UID:
		return telegram.Default()
	case twilio.UID:
		return twilio.Default()
	case webhook.UID:
		return webhook.Default()
	}
	return nil, fmt.Errorf("unknown notifier %s", via)
}

func Load(loader envconfig.LoaderFunc) (interface{}, error) {
	via, ok := loader(NotifyVia)
	if !ok {
		return nil, errors.New(`"NOTIFY_VIA" is not set.`)
	}
	return LoadVia(via, loader)
}

func LoadVia(via string, loader envconfig.LoaderFunc) (interface{}, error) {
	switch strings.ToLower(via) {
	case discord.UID:
		return discord.Load(loader)
	case log.UID:
		return log.Load(loader)
	case mailgun.UID:
		return mailgun.Load(loader)
	case mattermost.UID:
		return mattermost.Load(loader)
	case plivo.UID:
		return plivo.Load(loader)
	case pushover.UID:
		return pushover.Load(loader)
	case twilio.UID:
		return twilio.Load(loader)
	case slack.UID:
		return slack.Load(loader)
	case smtp.UID:
		return smtp.Load(loader)
	case telegram.UID:
		return telegram.Load(loader)
	case webhook.UID:
		return webhook.Load(loader)
	}
	return nil, fmt.Errorf("unknown notifier %s", via)
}
