package discord

import (
	"log/slog"
	"os"
	"tabbybot/internal/config"
	"tabbybot/internal/discord/commands"

	"github.com/bwmarrin/discordgo"
)

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Debug("Received interaction", "type", i.Type, "name", i.ApplicationCommandData().Name)
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
	err := b.Session.Open()
	if err != nil {
		panic(err)
	}
	slog.Info("Discord sessions opened.")
	slog.Info("Bot is now running. Press CTRL-C to exit.")
}

func (b *Bot) Close() {
	slog.Info("Closing Discord sessions...")
	err := b.Session.Close()
	if err != nil {
		panic(err)
	}
	slog.Info("Discord sessions closed.")
	os.Exit(0)
}
