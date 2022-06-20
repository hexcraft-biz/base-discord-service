package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/base-discord-service/config"
	"github.com/hexcraft-biz/base-discord-service/features"
)

func NewDiscordWebSocket(cfg config.ConfigInterface) *discordgo.Session {
	dg, err := discordgo.New("Bot " + cfg.GetDiscordToken())
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return nil
	}

	dg.Identify.Intents = cfg.GetDiscordIntents()

	return dg
}

func New(cfg config.ConfigInterface) *gin.Engine {
	engine := gin.Default()
	engine.SetTrustedProxies([]string{cfg.GetTrustProxy()})

	// common features
	features.LoadCommon(engine, cfg)

	return engine
}
