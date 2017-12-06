package main

import (
    "flag"
    "fmt"
    "strings"
    "strconv"
    "os"
    "os/exec"
    "gopkg.in/telegram-bot-api.v4"
)

const (
    ENV_APIKEY = "TGALERT_APIKEY"
    ENV_CHAT_ID = "TGALERT_CHAT_ID"
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
        fmt.Println("No command specified.");
        os.Exit(1);
    }

    apikey, ok := os.LookupEnv(ENV_APIKEY);
    if !ok {
        fmt.Printf("The following enviorment variable needs to be defined: %s\n", ENV_APIKEY);
	os.Exit(1);
    }

    chatid_s, ok2 := os.LookupEnv(ENV_CHAT_ID);
    if !ok2 {
        fmt.Printf("The following enviorment variable needs to be defined: %s\n", ENV_CHAT_ID);
	os.Exit(1);
    }

    chatid, err := strconv.Atoi(chatid_s)
    if err != nil {
        fmt.Println(err)
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
    
    msg := tgbotapi.NewMessage(int64(chatid), fmt.Sprintf("Executed command \"%s\": ", strings.Join(args, " ")))
    _, err2 := bot.Send(msg)
    fmt.Println(err2)

    msg2 := tgbotapi.NewMessage(int64(chatid), string(out))
    bot.Send(msg2)

}
