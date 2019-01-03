package main

import (
	"bytes"
	"text/template"
)

type Message interface {
	Make(*map[string]string) (string, error)
}

type TextMessage struct {
	Pattern string
}

func (t *TextMessage) Make(m *map[string]string) (string, error) {
	var tplB bytes.Buffer
	tmp, err := template.New("").Parse(t.Pattern)
	if err != nil {
		return "", err
	}

	err = tmp.Execute(&tplB, m)

	if err != nil {
		return "", err
	}

	return tplB.String(), nil
}