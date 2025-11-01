package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	// Project
	VERSION = os.Getenv("VERSION")
	ENV     = os.Getenv("ENV")

	// Discord
	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	BOT_ID        = os.Getenv("BOT_ID")

	// HTTP
	HTTP_PORT = os.Getenv("HTTP_PORT")
)
