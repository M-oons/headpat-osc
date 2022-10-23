package vrc

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
	"github.com/m-oons/headpat-osc/config"
)

const inputAddress = "/chatbox/input"

var Client *osc.Client

func SetupOsc() {
	Client = osc.NewClient(config.Current.Osc.Host, int(config.Current.Osc.Port))

	dispatcher := osc.NewStandardDispatcher()
	dispatcher.AddMsgHandler("/avatar/parameters/Headpat", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	server := osc.Server{
		Addr:       fmt.Sprintf("%s:%d", config.Current.Osc.Host, config.Current.Osc.Port),
		Dispatcher: dispatcher,
	}
	server.ListenAndServe()
}

func SendMessage(text string) error {
	text = buildMessage(text)
	msg := osc.NewMessage(inputAddress)
	msg.Append(text)
	msg.Append(true)
	return Client.Send(msg)
}
