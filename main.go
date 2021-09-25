package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := "TOKEN that you get on Discord API"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln(err)
	}
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	
	err = dg.Open()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Command on Discord" {
		s.ChannelMessageSend(m.ChannelID, "Bot reply on discord")
	}
}