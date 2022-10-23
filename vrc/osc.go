package vrc

import (
	"fmt"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/m-oons/headpat-osc/config"
	"github.com/m-oons/headpat-osc/util"
)

const inputAddress = "/chatbox/input"

var Client *osc.Client

var lastMessage time.Time
var messageLimit = time.Second * 2

func SetupOsc() {
	Client = osc.NewClient(config.Current.Osc.Host, int(config.Current.Osc.SendPort))

	dispatcher := osc.NewStandardDispatcher()
	dispatcher.AddMsgHandler("/avatar/parameters/Headpat", func(msg *osc.Message) {
		if len(msg.Arguments) == 0 {
			return
		}

		headpat, ok := msg.Arguments[0].(bool)
		if !ok {
			return
		}

		if headpat {
			AddHeadpat()
			SendMessage(config.Current.Headpat.Message)
			util.Log("Started receiving headpat (Total: %d)", GetHeadpats())
		} else {
			util.Log("Stopped receiving headpat")
		}
	})

	server := osc.Server{
		Addr:       fmt.Sprintf("%s:%d", config.Current.Osc.Host, config.Current.Osc.ReceivePort),
		Dispatcher: dispatcher,
	}
	util.Log("Waiting for headpats...")
	server.ListenAndServe()
}

func SendMessage(text string) error {
	if time.Since(lastMessage) < messageLimit { // rate limit
		return nil
	}

	text = buildMessage(text)
	msg := osc.NewMessage(inputAddress)
	msg.Append(text)
	msg.Append(true)
	err := Client.Send(msg)
	lastMessage = time.Now()
	return err
}
