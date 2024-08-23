package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sagleft/uchatbot-engine"
)

type App interface {
	Run() error
}

type app struct {
	chatBot *uchatbot.ChatBot
}

func New() (App, error) {
	a := &app{}
	return a, a.Run()
}

func (a *app) Run() error {
	data := uchatbot.ChatBotData{}

	var err error
	a.chatBot, err = uchatbot.NewChatBot(data)
	if err != nil {
		return fmt.Errorf("create chatbot: %w", err)
	}

	a.waitForFinish()
	return nil
}

func (a *app) waitForFinish() {
	cancelChan := make(chan os.Signal, 1)
	// catch SIGETRM or SIGINTERRUPT
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	<-cancelChan
}
