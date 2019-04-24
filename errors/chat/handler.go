package chat

import (
	"fmt"

	utilerrors "gomodules.xyz/notify/errors"
	"github.com/pkg/errors"
	"gomodules.xyz/notify"
)

// messenger sends chat via notifier in case of errors.
type messenger struct {
	notifier notify.ByChat
}

var _ utilerrors.Handler = &messenger{}

func New(notifier notify.ByChat) utilerrors.Handler {
	return &messenger{notifier}
}

func (h *messenger) Handle(err error, st errors.StackTrace) {
	body := fmt.Sprintf("%s\n%+v", err, st)
	h.notifier.WithBody(body).Send()
}
