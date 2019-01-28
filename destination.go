package oldNotifier

type Destination interface {
	AddAddressee(addressee string)
	GetAddressee() []string
	ToDestination() interface{}
}

type TelegramDestination struct {
	addressee []string
}

func (t *TelegramDestination) AddAddressee(addressee string) {
	t.addressee = append(t.addressee, addressee)
}

func (t *TelegramDestination) GetAddressee() []string {
	return t.addressee
}

func (*TelegramDestination) ToDestination() interface{} {
	panic("implement me")
}
