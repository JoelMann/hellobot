package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Token struct {
	Secret string
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "8ball" {
		answers := []string{"Yes", "No", "Maybe", "OHMYGOSH Gurl Abso-fucking-lutely!", "Oh Baby, your crazy, noooooo!", "Dunno dawg, seems sketch."}
		reply := answers[rand.Intn(len(answers))]
		s.ChannelMessageSend(m.ChannelID, reply)
	}
}

func main() {
	fmt.Println("Hello World")
	var TOKEN Token
	tokenFile, _ := os.Open("secrets.json")
	defer tokenFile.Close()
	tokenByte, _ := io.ReadAll(tokenFile)

	json.Unmarshal(tokenByte, &TOKEN)
	client, err := discordgo.New("Bot " + TOKEN.Secret)
	if err != nil {
		fmt.Println(err)
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	client.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	client.Identify.Intents = discordgo.IntentsGuildMessages

	client.Open()
	fmt.Println("BOT IS LIVE")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	client.Close()
}
