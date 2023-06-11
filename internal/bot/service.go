package bot

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lulzshadowwalker/saoirse/internal/helpers"
	"github.com/lulzshadowwalker/saoirse/internal/weather"
)

var config *BotConfig

func Init() {
	config = readConfig()

	sesh, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(fmt.Sprintf(`
error creating a discord session.
err: %q
`, err))
	}

	sesh.Identify.Intents = discordgo.IntentsGuildMessages

	sesh.AddHandler(handleMessage)

	err = sesh.Open()
	defer sesh.Close()

	if err != nil {
		panic(fmt.Sprintf(`
error connecting to Discord.
%q		
`, err))
	}

	fmt.Println("Saoirse is runnin 🏃🏻‍♀️")
}

func readConfig() *BotConfig {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config, err := helpers.ReadConfig[BotConfig](filepath.Join(cwd, "config/bot-config.json"))
	if err != nil {
		panic(fmt.Sprintf(`
error reading bot config
err: %q`, err))
	}
	return config
}

func handleMessage(sesh *discordgo.Session, msg *discordgo.MessageCreate) {
	fmt.Println("message event received", msg.Content)

	if msg.Author.ID == sesh.State.User.ID || string(msg.Content[0]) != config.Prefix {
		return
	}

	args := strings.Split(msg.Content, " ")

	var res string
	switch strings.ToLower(args[1]) {
	case "hai", "henlo", "hello", "privet", "Dia dhuit":
		res = fmt.Sprintf("henlo %s", msg.Author.Username)

	case "weather":
		city := args[2]

		weather, err := weather.Fetch(city)
		if err != nil || weather.Location.Name == "" {
			res = fmt.Sprintf("I don't know about %s but you are sure as hell fkin hot today <:horny:1117392856286511104>", city)
		} else {
			res = strconv.FormatFloat(weather.Current.TempC, 'f', 1, 32)
		}
	}

	if res != "" {
		sesh.ChannelMessageSend(msg.ChannelID, res)
	}
}
