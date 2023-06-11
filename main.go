package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token  string
	Prefix string `json:"command-prefix"`
}

func (c Config) String() string {
	return fmt.Sprintf(`
token: %q
command-prefix: %q
`, c.Token, c.Prefix)
}

func main() {
	configFile, err := os.Open("./config.json")

	if err != nil {
		panic(fmt.Sprintf("error reading config.json\nerr: %q", err))
	}

	config := Config{}
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		panic(fmt.Sprintf(`
error decoding config.json.
err: %q
`, err))
	}

	fmt.Println(config)

	sesh, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(fmt.Sprintf(`
error decoding config.json.
err: %q
`, err))
	}

	sesh.Identify.Intents = discordgo.IntentsGuildMessages

	sesh.AddHandler(func(sesh *discordgo.Session, msg *discordgo.MessageCreate) {
		fmt.Println("message event received", msg.Content)

		if msg.Author.ID == sesh.State.User.ID || string(msg.Content[0]) != config.Prefix {
			return
		}

		var res string
		switch strings.ToLower(strings.Trim(msg.Content[1:], " ")) {
		case "hai", "henlo", "hello", "privet", "Dia dhuit":
			res = fmt.Sprintf("henlo %s", msg.Author.Username)
		}

		if res != "" {
			sesh.ChannelMessageSend(msg.ChannelID, res)
		}
	})

	err = sesh.Open()
	defer sesh.Close()

	if err != nil {
		panic(fmt.Sprintf(`
error decoding config.json.
err: %q
`, err))
	}

	fmt.Println("Saoirse is running 🏃🏻‍♀️")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
