package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lulzshadowwalker/saoirse/internal/bot"
)

func main() {
	bot.Init()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
