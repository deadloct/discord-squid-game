package main

import (
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/deadloct/bitheroes-hg-bot/cmd"
	"github.com/deadloct/bitheroes-hg-bot/game"
	"github.com/deadloct/bitheroes-hg-bot/settings"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("verbose logs enabled")
	log.SetLevel(log.DebugLevel)

	settings.ImportData()
}

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("BITHEROES_HG_BOT_AUTH_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	data, err := os.ReadFile(path.Join(settings.DataLocation, settings.PhrasesFile))
	if err != nil {
		log.Panic(err)
	}
	commandManager := cmd.NewManager(data)

	// Listen for server messages only
	session.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMessageReactions | discordgo.IntentMessageContent
	session.AddHandler(commandManager.CommandHandler)
	session.AddHandler(game.ManagerInstance(session).ReactionHandler)
	if err := session.Open(); err != nil {
		log.Panic(err)
	}

	err = commandManager.RegisterCommands(session)
	if err != nil {
		log.Panicf("error registering slash commands: %v", err)
	}
	defer commandManager.DeregisterCommmands(session)

	log.Info("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Info("Bot exiting...")
}
