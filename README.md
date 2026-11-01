# go-projects

Small practice repo, built up over three PRs. It's one HTTP server that greets a user by name and sends the same greeting to a Telegram chat.

Stack: Go 1.24, `net/http`, go-telegram-bot-api v5.

## How to run

The bot token and the chat id come from environment variables:

```bash
export TELEGRAM_TOKEN=<token from BotFather>
export TELEGRAM_CHAT_ID=<chat id to send greetings to>
go run ./main
```

The server listens on `:8080`.

## Endpoints

- `GET /` prints "Привет" to stdout and returns an empty 200.
- `POST /sayHello` with `{"name": "Irena"}` sends "Привет, Irena" to the chat and returns the same text. Other methods get 405.

```bash
curl -X POST localhost:8080/sayHello -d '{"name":"Irena"}'
```
