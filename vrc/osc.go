package vrc

import (
	"github.com/hypebeast/go-osc/osc"
	"github.com/m-oons/headpat-osc/config"
)

const inputAddress = "/chatbox/input"

var Client *osc.Client

func init() {
	Client = osc.NewClient(config.Current.Osc.Host, int(config.Current.Osc.Port))
}

func SendMessage(text string) error {
	text = buildMessage(text)
	msg := osc.NewMessage(inputAddress)
	msg.Append(text)
	msg.Append(true)
	return Client.Send(msg)
}
