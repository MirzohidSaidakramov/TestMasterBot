package main

import (
	"flag"
	"log"
	"testMaster/clients/telegram"
)

func main() {
	// token = flags.Get(token)

	tgClient = telegram.New("api.telegram.org", mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New

	// consumer.Start(fetcher,processor)

}

func mustToken() string {
	//
	token := flag.String("token-bot-token", "", "token for access to telehram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")

	}

	return *token
}
