package commands

import "github.com/bwmarrin/discordgo"

var InfoCommand = &discordgo.ApplicationCommand{
	Name:        "info",
	Description: "Get information about the bot",
}

func Info() {

}
