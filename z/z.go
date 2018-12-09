package z

import (
	"bytes"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"os"
)

var I int

type x struct {
	i int
}

var Xx x


func init () {
	I = rand.Int()
	Xx.i = rand.Int()
	fmt.Println(I)
	fmt.Println(Xx.i)
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

	bot, err := tgbotapi.NewBotAPI("634747898:AAHhwfMTu0tDMNtlkICGgpTKsWrDuIx7ZrQ")
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

func Send1() {
	const width, height = 1080, 1080

	bot, err := tgbotapi.NewBotAPI("634747898:AAHhwfMTu0tDMNtlkICGgpTKsWrDuIx7ZrQ")
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

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		//bot.Send(msg)

		//MakeHell(width,height)
		//file := tgbotapi.FileBytes{
		//	Name: "photo.jpg",
		//	Bytes: fb.Bytes(),
		//}
		f, err := os.Open("hell.png")
		if err != nil {panic("pizdec")}
		fmt.Println(f)
		fb := &bytes.Buffer{}
		fb.ReadFrom(f)
		fmt.Println(fb.Len())

		file := tgbotapi.FileBytes{
			Name:  "hell.png",
			Bytes: fb.Bytes(),
		}

		photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, file)

		bot.Send(photo)

	}
}
