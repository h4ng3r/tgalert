package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	envAPIKey   = "TGALERT_APIKEY"
	envChatID   = "TGALERT_CHAT_ID"
	envHostname = "TGALERT_HOSTNAME"
	msgText     = "Finished the execution of the command"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: tgalert [comand]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("No command specified.")

	}

	apikey, ok := os.LookupEnv(envAPIKey)
	chatidS, ok2 := os.LookupEnv(envChatID)
	if !ok || !ok2 {
		fmt.Printf("Both environment variables %s and %s are mandatory", envAPIKey, envChatID)
		os.Exit(1)
	}

	chatid, err := strconv.Atoi(chatidS)
	if err != nil {
		fmt.Println("[ERROR] The chat id is not a number")
		os.Exit(1)
	}

	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprint(os.Stderr, string(out))
	} else {
		fmt.Print(string(out))
	}

	bot, err := tgbotapi.NewBotAPI(apikey)
	if err != nil {
	}

	var host string
	var ok3 bool
	host, ok3 = os.LookupEnv(envHostname)
	if !ok3 {
		user, _ := user.Current()
		hostname, _ := os.Hostname()
		host = fmt.Sprintf("%s@%s", user.Name, hostname)
	}

	msg := tgbotapi.NewMessage(int64(chatid), fmt.Sprintf("[%s] %s \"%s\": ", host, msgText, strings.Join(args, " ")))
	_, err2 := bot.Send(msg)
	if err2 != nil {
		fmt.Printf("[ERROR] %s\n", err2)
	}
	msg2 := tgbotapi.NewMessage(int64(chatid), string(out))
	bot.Send(msg2)
}
