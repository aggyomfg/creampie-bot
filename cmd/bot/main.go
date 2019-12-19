package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aggyomfg/creampie-bot/internal/app/bot"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	helpFlag := flag.Bool("help", false, "show help")
	flag.Parse()
	if *helpFlag == true {
		fmt.Printf("help")
	}

	config := bot.NewConfig()
	if err := env.Parse(config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", config)

	bot.Start(config)

}
