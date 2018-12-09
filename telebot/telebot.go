package telebot

import (
	"bytes"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/instahell/hell"
	"log"
	"os"
)

type TeleBot struct {
	hell.TeleHell

}

func Send (){


	f, err := os.Open("hell.png")
	if err != nil {panic("pizdec")}
	defer f.Close()
	fb := &bytes.Buffer{}
	fb.ReadFrom(f)
	fmt.Println(fb.Len())

	file := tgbotapi.FileBytes{
		Name:  "hell.png",
		Bytes: fb.Bytes(),
	}

	bot, err := tgbotapi.NewBotAPI("603813331:AAGe1YrOra-SPHLd5O7KOljQOPsSJl6x3Os")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)


	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, file)
		bot.Send(photo)
	}
}