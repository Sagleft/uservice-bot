package app

import (
	"bot/pkg/config"
	"bot/pkg/db"
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
	dbConn  db.DB
	chatBot *uchatbot.ChatBot
}

func New() (App, error) {
	a := &app{}
	return a, a.Run()
}

func (a *app) Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load: %w", err)
	}

	a.dbConn, err = db.New(cfg.DB, cfg.DB.TablePrefix)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}

	data := uchatbot.ChatBotData{
		Config: cfg.Utopia,
		// TODO: other data
	}

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
