package main

import (
	"github.com/CatchZeng/feishu/pkg/feishu"
)

func SendMessage() string {
	token := "318e87b6-c6ee-4d66-8169-b60a881cac25"
	key := "wNeRboPX4HSHrrdG5l6oU"
	client := feishu.NewClient(token, key)
	msg := feishu.NewTextMessage()
	msg.Content.Text = "hello world"
	_, respone, err := client.Send(msg)
	if err != nil {
		panic(err)
	}
	return respone.Msg
}
func main() {
	SendMessage()
}
