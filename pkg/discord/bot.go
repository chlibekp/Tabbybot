package discord

import (
	"log/slog"
	"os"
	"tabbybot/pkg/config"
	"tabbybot/pkg/discord/commands"
	"tabbybot/pkg/metrics"

	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var err error
	switch i.ApplicationCommandData().Name {
	case commands.InfoCommand.Name:
		err = commands.Info(s, i)
	}

	if err != nil {
		slog.Error("Error handling command", "error", err)

		errorUuid := uuid.New().String()

		// Add error message to embed
		errorEmbed := &discordgo.MessageEmbed{
			Title:       "Error | An error has occured",
			Description: "Please contact the bot developer with the following error ID \n `" + errorUuid + "`",
			Color:       0xFF0000,
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{errorEmbed},
			},
		})
	}
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Select handler based on the interaction type
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		slog.Debug("Received interaction", "type", i.Type, "name", i.ApplicationCommandData().Name)
		// Handle command
		go handleCommands(s, i)
	}
}

func registerCommands(s *discordgo.Session) error {
	commands := []*discordgo.ApplicationCommand{
		commands.InfoCommand,
	}

	slog.Info("Registering commands...")

	// Update all commands
	_, err := s.ApplicationCommandBulkOverwrite(config.BOT_ID, "", commands)
	if err != nil {
		return err
	}

	slog.Info("Commands registered.")
	return nil
}

type Bot struct {
	Session *discordgo.Session
}

func NewBot() *Bot {
	session, err := discordgo.New("Bot " + config.DISCORD_TOKEN)
	if err != nil {
		panic(err)
	}

	// Register commands
	if err := registerCommands(session); err != nil {
		panic(err)
	}

	// Add all handlers
	session.AddHandler(commandHandler)

	return &Bot{
		Session: session,
	}
}

func (b *Bot) Start() {
	slog.Info("Opening Discord sessions...")
	metrics.StartTime = time.Now()
	if err := b.Session.Open(); err != nil {
		panic(err)
	}
	slog.Info("Discord sessions opened.")
	slog.Info("Bot is now running. Press CTRL-C to exit.")
}

func (b *Bot) Close() {
	slog.Info("Closing Discord sessions...")
	if err := b.Session.Close(); err != nil {
		panic(err)
	}
	slog.Info("Discord sessions closed.")
	os.Exit(0)
}
