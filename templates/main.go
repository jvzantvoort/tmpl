package messages

import (
	"embed"
	"fmt"
	"io/fs"
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

func ListTemplates(lang string) ([]string, error) {
	retv := []string{}
	files, err := fs.ReadDir(Content, lang)
	if err != nil {
		log.Errorf("Error listing embedded targets: %s", err)
	return retv, err
	}
	for _, item := range files {
		retv = append(retv, item.Name())
	}
	return retv, err
}
