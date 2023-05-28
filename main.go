package main

import (
	"flag"
	"log"

	"github.com/VladPorvin/content-saver-bot/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	// token = flags.Get(token)
	tgClient := telegram.New(tgBotHost, mustToken())
	// fetcher := fetcher.New()
	// processor := processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}

	return *token
}
