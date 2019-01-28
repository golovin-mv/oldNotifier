package oldNotifier

type Notice interface {
	Notify(params *map[string]string) error
}

type TextNotice struct {
	deliveryman Deliveryman
	message     NotifyMessage
	NoticeType  string
}

func (t *TextNotice) Notify(params *map[string]string) error {
	return t.deliveryman.Deliver(t.message)
}
