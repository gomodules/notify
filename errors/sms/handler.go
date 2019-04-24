package sms

import (
	utilerrors "gomodules.xyz/notify/errors"
	"github.com/pkg/errors"
	"gomodules.xyz/notify"
)

// messenger sends chat via notifier in case of errors.
type messenger struct {
	notifier notify.BySMS
}

var _ utilerrors.Handler = &messenger{}

func New(notifier notify.BySMS) utilerrors.Handler {
	return &messenger{notifier}
}

func (h *messenger) Handle(err error, st errors.StackTrace) {
	h.notifier.WithBody(err.Error()).Send()
}
