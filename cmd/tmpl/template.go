package main

import (
	"fmt"
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
	Company string
	License string
	MailAddress string
	Time    string
	AuthorStr string
}

func (t Template) Template(path string) string {
	data := tmpl.GetTemplate(t.Lang, path)

	return t.Parse(data)
}

// buildConfig contruct the text from the template definition and arguments.
func (t *Template) Parse(templatestring string) string {
	t.AuthorStr = fmt.Sprintf("%s (%s), %s", t.Login, t.Author, t.MailAddress)

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
	if len(Name) != 0 {
		retv.Author = Author
	}

	if len(Company) != 0 {
		retv.Company = Company
	} else {
		retv.Company = "None"
	}

	if len(License) != 0 {
		retv.License = License
	} else {
		retv.License = "None"
	}

	if len(MailAddress) != 0 {
		retv.MailAddress = MailAddress
	} else {
		retv.MailAddress = "None"
	}

	retv.Login = usr.Username


	return retv
}
