# tgalert [WIP]

Execute a command and receive a Telegram alert when it finish

## Configuration

```
export TGALERT_CHAT_ID="999999999"
export TGALERT_APIKEY="999999999:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
export TGALERT_HOSTNAME="username@hostname" (optional)
```

To get the CHAT_ID just https://api.telegram.org/bot$TGALERT_APIKEY/getUpdates and speak with him.

## Run

```
go run main.go ls
LICENSE
Makefile
README.md
main.go
```

This will also send a telegram message like this one:
```
Finished the execution of the command "ls":

LICENSE
Makefile
README.md
main.go
```
