package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bunin/imss/app"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := app.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("failed to parse env", err)
	}
	if err := cfg.Validate(); err != nil {
		log.Fatalln("invalid env", err)
	}
	a := app.New(cfg)
	go func() {
		if err := a.Run(); err != nil {
			log.Fatalln("failed to start app", err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	for s := range c {
		log.Println("got signal " + s.String() + ". stopping.")
		if err := a.Stop(); err != nil {
			log.Fatalln("failed to stop the app", err)
		}
		os.Exit(0)
	}
}
