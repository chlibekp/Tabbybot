package commands

import (
	"fmt"
	"runtime"
	"time"

	"tabbybot/internal/config"
	"tabbybot/internal/metrics"

	"github.com/bwmarrin/discordgo"
)

var InfoCommand = &discordgo.ApplicationCommand{
	Name:        "info",
	Description: "Get information about the bot",
}

func Info(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	uptime := time.Since(metrics.StartTime).Round(time.Second)
	goVersion := runtime.Version()

	guildCount := len(s.State.Guilds)
	shardInfo := fmt.Sprintf("Shard %d/%d", s.ShardID, s.ShardCount)

	// Build a beautiful embed
	embed := &discordgo.MessageEmbed{
		Title:       "Status",
		Description: "Here are the latest metrics for the bot.",
		Color:       0xfecb4d, // Yellow
		Timestamp:   time.Now().Format(time.RFC3339),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Uptime",
				Value:  uptime.String(),
				Inline: true,
			},
			{
				Name:   "Guilds",
				Value:  fmt.Sprintf("%d", guildCount),
				Inline: true,
			},
			{
				Name:   "Shard",
				Value:  shardInfo,
				Inline: true,
			},
			{
				Name:   "Go runtime",
				Value:  goVersion,
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Version: %s", config.VERSION),
		},
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
