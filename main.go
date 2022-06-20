package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/hexcraft-biz/base-discord-service/service"
)

// https://discord.com/api/oauth2/authorize?client_id=988268853006045194&permissions=8&scope=bot

func main() {
	appCfg := &AppConfig{}

	// Open a websocket connection to Discord and begin listening.
	appCfg.DG = service.NewDiscordWebSocket(appCfg)
	if err := appCfg.DG.Open(); err != nil {
		fmt.Println("error Discord WS opening connection,", err)
		return
	} else {
		fmt.Println("Discord Bot is now running.")
	}

	// TODO add discord event hanlder
	// Do something...

	ginEngine := service.New(appCfg)

	// TODO add other gin route.
	// Do something...

	ginEngine.Run(":" + os.Getenv("APP_PORT"))
}

func MustNot(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//================================================================
// AppConfig implement ConfigInterface
//================================================================
type AppConfig struct {
	DG *discordgo.Session
}

func (ac *AppConfig) GetDG() *discordgo.Session {
	return ac.DG
}

func (ac *AppConfig) GetDiscordToken() string {
	return os.Getenv("DISCORD_TOKEN")
}

func (ac *AppConfig) GetDiscordIntents() discordgo.Intent {
	return discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions
}

func (ac *AppConfig) GetTrustProxy() string {
	return os.Getenv("TRUST_PROXY")
}
