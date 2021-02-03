package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

func main() {
	fmt.Println("vim-go")

	bot, err := discordgo.New("Bot " + os.Getenv("token"))

	if err != nil {
		panic(err)
	}

	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

	err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session:", err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "with Go")
	fmt.Println("logged in as user " + string(s.State.User.ID))
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Your code goes here

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
