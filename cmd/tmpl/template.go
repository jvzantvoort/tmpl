package main

import (
	"bytes"
	"os/user"
	"text/template"
	"time"
	tmpl "github.com/jvzantvoort/tmpl/templates"
)

type Template struct {
	Author  string
	Date    string
	Homedir string
	Lang    string
	Login   string
	Name    string
	Time    string
}

func (t Template) Template(path string) string {
	data := tmpl.GetTemplate(t.Lang, path)

	return t.Parse(data)
}

// buildConfig contruct the text from the template definition and arguments.
func (t Template) Parse(templatestring string) string {

	tmpl, err := template.New("prompt").Parse(templatestring)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func NewTemplate(lang string) *Template {
	retv := &Template{}
	retv.Lang = lang
	retv.Name = lang
	now := time.Now()

	retv.Date = now.Format("2006-01-02")
	retv.Time = now.Format("15:04:05")

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	retv.Homedir = usr.HomeDir
	retv.Author = usr.Name
	retv.Login = usr.Username
	return retv
}
