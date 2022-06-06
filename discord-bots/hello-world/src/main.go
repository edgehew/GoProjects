package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "strings"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
    Token string
)

func init() {
    flag.StringVar(&Token, "t", "", "Bot Token")
    flag.Parse()
}

func createResponse(m *discordgo.MessageCreate) (s string) {
    var b strings.Builder
    if m.Content == "!hello" {
        fmt.Fprintf(&b, "Hello %s", m.Author.Username)
    } else if m.Content == "!complement" {
        complement := getRandomComplement()
        fmt.Fprintf(&b, "Hello %s! %s", m.Author.Username, complement)
    } else if m.Content == "!insult" {
        insult := getRandomInsult()
        fmt.Fprintf(&b, "Hello %s! %s", m.Author.Username, insult)
    }

    return b.String()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
        return
    }

    message := createResponse(m)
    _, err := s.ChannelMessageSend(m.ChannelID, message)

    if err != nil {
        fmt.Println(err)
    }
}

func main() {

    // Create a new Discord session using the provided bot token.
    dg, err := discordgo.New("Bot " + Token)
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
    fmt.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    dg.Close()
}