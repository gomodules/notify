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
	case plivo.UID:
		return plivo.Default()
	case twilio.UID:
		return twilio.Default()
	case smtp.UID:
		return smtp.Default()
	case mailgun.UID:
		return mailgun.Default()
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
	case plivo.UID:
		return plivo.Load(loader)
	case twilio.UID:
		return twilio.Load(loader)
	case smtp.UID:
		return smtp.Load(loader)
	case mailgun.UID:
		return mailgun.Load(loader)
	case slack.UID:
		return slack.Load(loader)
	case log.UID:
		return log.Load(loader)
	case webhook.UID:
		return webhook.Load(loader)
	case pushover.UID:
		return pushover.Load(loader)
	case telegram.UID:
		return telegram.Load(loader)
	case discord.UID:
		return discord.Load(loader)
	}
	return nil, fmt.Errorf("unknown notifier %s", via)
}
