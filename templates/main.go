package messages

import (
	"embed"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Content missing godoc.
//
//go:embed bash/*
var Content embed.FS

func GetTemplate(lang, name string) string {
	filename := fmt.Sprintf("%s/%s", lang, name)
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}
