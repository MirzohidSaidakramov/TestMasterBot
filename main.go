package main

import (
	"flag"
	"log"
)

func main() {
	// token = flags.Get(token)
	t := mustToken()

	// tgClient = telegram.New(token)

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
