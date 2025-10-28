package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	BOT_ID        = os.Getenv("BOT_ID")
	ENV           = os.Getenv("ENV")
)
