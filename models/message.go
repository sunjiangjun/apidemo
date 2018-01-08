package models

import "github.com/astaxie/beego"

type Message struct {
	Id   string
	Body string
}

func GetMessage() *Message {
	return &Message{"1", beego.AppConfig.String("tag")}
}
