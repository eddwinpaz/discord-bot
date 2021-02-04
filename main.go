package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/eddwinpaz/discord-bot/entities"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "hola" {
		s.ChannelMessageSend(m.ChannelID, "Hola! mi rey")
	}

	if m.Content == "sol de mexico" {
		s.ChannelMessageSend(m.ChannelID, "asi me dicen, que quieres mi rey?")
	}

	if m.Content == "/opciones" {
		s.ChannelMessageSend(m.ChannelID, "Que te puedo ayudar mi rey? \n**uf** - retorna el valor de la uf del dia \n**dolar** - retorna el valor del dolar del dia")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "uf" {
		uf := GetEconomyValue("uf")
		s.ChannelMessageSend(m.ChannelID, uf)
	}

	if m.Content == "dolar" {
		uf := GetEconomyValue("dolar")
		s.ChannelMessageSend(m.ChannelID, uf)
	}

	if strings.Contains(m.Content, "empleos como") {
		word := strings.Replace(m.Content, "empleos como", "", -1)
		query := strings.Replace(word, " ", "+", -1)
		jobs := SearchGetOnBoardJobsByTitle(query)
		s.ChannelMessageSend(m.ChannelID, jobs)
	}
}

// GetEconomyValue returns value from API for UF and US Dollars of todays value
func GetEconomyValue(indicator string) string {
	url := fmt.Sprintf("https://mindicador.cl/api/%s", indicator)
	responseChannel := make(chan string)
	msg := "Lo siento wey! el servicio no me dio el valor del dia. seguire buscando; **palabra de honor**"
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			responseChannel <- msg
			close(responseChannel)
			return
		}
		defer resp.Body.Close()
		var response entities.Indicator
		_ = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			responseChannel <- msg
			close(responseChannel)
			return
		}
		firstValue := response.Serie[0]
		if indicator == "uf" {
			msg = fmt.Sprintf("wey! **%s** esta hoy a **$ %.2f %s**",
				response.Nombre, firstValue.Valor, response.UnidadMedida)
		} else if indicator == "dolar" {
			msg = fmt.Sprintf("wey! **%s** esta hoy a **$ %.2f %s**",
				response.Nombre, firstValue.Valor, response.UnidadMedida)
		}
		responseChannel <- msg
		close(responseChannel)
	}()
	return <-responseChannel

}

// SearchGetOnBoardJobsByTitle search jobs on getonbrd.com
func SearchGetOnBoardJobsByTitle(title string) string {

	url := fmt.Sprintf("https://www.getonbrd.com/api/v0/search/jobs?query=%s", title)

	responseChannel := make(chan string)

	msg := "Lo siento wey! el servicio no me dio el valor del dia. seguire buscando; **palabra de honor**"

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			responseChannel <- msg
		}
		defer resp.Body.Close()
		var response entities.GetOnBoard

		_ = json.NewDecoder(resp.Body).Decode(&response)

		if err != nil {
			responseChannel <- msg
			close(responseChannel)
			return
		}

		if len(response.Data) == 0 {
			responseChannel <- msg
			close(responseChannel)
			return
		}

		firstValue := response.Data[0]
		msg = fmt.Sprintf("wey! en getonbrd.com hay una empresa que busca **%s** en **%s** pagando salario de $ %.2f USD más info aqui %s",
			firstValue.Attributes.Title, firstValue.Attributes.Country, firstValue.Attributes.MaxSalary, firstValue.Links.PublicURL)

		responseChannel <- msg
		close(responseChannel)
	}()
	return <-responseChannel

}
