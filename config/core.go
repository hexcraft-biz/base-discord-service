package config

import "github.com/bwmarrin/discordgo"

type ConfigInterface interface {
	GetDG() *discordgo.Session
	GetDiscordToken() string
	GetDiscordIntents() discordgo.Intent
	GetDiscordServerID() string
	GetTrustProxy() string
}
