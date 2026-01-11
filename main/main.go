package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irenakk/go-projects.git/config"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

type notifyHandler struct {
	Bot    *tgbotapi.BotAPI
	ChatId int64
}

func (u User) SayHello() string {
	name := u.Name
	return "Привет, " + name
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Привет")
}

func (h *notifyHandler) handlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	msg := tgbotapi.NewMessage(h.ChatId, user.SayHello())
	if _, err := h.Bot.Send(msg); err != nil {
		log.Println("send error:", err)
		http.Error(w, "telegram send error", 500)
		return
	}

	w.Write([]byte(user.SayHello()))
	w.WriteHeader(http.StatusOK)
}

func main() {
	cfg := config.LoadConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}
	bot.Debug = false

	notifyHandler := &notifyHandler{
		Bot:    bot,
		ChatId: cfg.ChatId,
	}

	http.HandleFunc("/", handlerGet)
	http.HandleFunc("/sayHello", notifyHandler.handlerPost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
