package oldNotifier

import (
	"bytes"
	"text/template"

	"github.com/jinzhu/gorm"
)

type NotifyMessage interface {
	Make() (string, error)
	AddParams(*map[string]string)
}

type TextMessage struct {
	gorm.Model
	pattern string
	params  *map[string]string `gorm:"-"`
}

func NewTextMessage(pattern string) *TextMessage {
	return &TextMessage{pattern: pattern}
}

func (t *TextMessage) AddParams(m *map[string]string) {
	t.params = m
}

func (t *TextMessage) Make() (string, error) {
	var tplB bytes.Buffer
	tmp, err := template.New("").Parse(t.pattern)
	if err != nil {
		return "", err
	}

	err = tmp.Execute(&tplB, t.params)

	if err != nil {
		return "", err
	}

	return tplB.String(), nil
}
