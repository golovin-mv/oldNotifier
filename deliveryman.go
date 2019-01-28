package oldNotifier

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Deliveryman interface {
	Deliver(message NotifyMessage) error
}

type TelegramDelivery struct {
	bot         *tgbotapi.BotAPI
	Destination Destination
}

func NewTelegramDelivery(destination Destination) *TelegramDelivery {
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	isDebug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		isDebug = false
	}

	bot.Debug = isDebug

	if err != nil {
		log.Panic(err)
	}

	return &TelegramDelivery{bot: bot, Destination: destination}
}

func (t *TelegramDelivery) sendMessageToChat(chatId int64, message NotifyMessage) error {
	msgText, err := message.Make()

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatId, msgText)

	_, err = t.bot.Send(msg)

	return err

}

func (t *TelegramDelivery) Deliver(message NotifyMessage) error {
	var errorsArr []string

	for _, val := range t.Destination.GetAddressee() {
		chatId, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			errorsArr = append(errorsArr, err.Error())
			continue
		}
		if err := t.sendMessageToChat(chatId, message); err != nil {
			errorsArr = append(errorsArr, err.Error())
		}
	}

	if len(errorsArr) == 0 {
		return nil
	}

	return errors.New(strings.Join(errorsArr, ";"))
}
