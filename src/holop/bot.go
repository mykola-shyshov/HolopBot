package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/robfig/cron"
	"holop/agents"
	"log"
	"math/rand"
	"strings"
)

var greetings = [5]string{"Да, господин", "К вашим услугам", "Покорнейше", "Растилаюсь у Ваших ног", "Выполняю"}
var chatID int64 = 388041982

func main() {

	bot, err := tgbotapi.NewBotAPI("439068576:AAF_H63gW2m7Pm2AFUTJTIW2OkUxePG_arQ")
	if err != nil {
		log.Panic(err)
	}

	cronObj := cron.New()
	cronObj.AddFunc("0 15 7 * * ?", func() {
		if (chatID == 0) {
			return
		}
		msg := tgbotapi.NewMessage(chatID, getRainPrediction())
		bot.Send( msg )
	})
	cronObj.Start()

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Chat id - %d", update.Message.Chat.ID)
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		message := handleMessage(update.Message.From.ID, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

		bot.Send(msg)
	}
}

func handleMessage(userId int, message string) string {
	command := strings.ToLower(message)

	if strings.Index(command, "hello") != -1 {
		return "hello!"

	} else if strings.Index(command, "wind") != -1 {
		return getWindPrediction()
		
	} else if strings.Index(command, "?") != -1 {
		return "Приклоняюсь пред Вами\nwind - скорость ветра\nrain - дождь сегодня "

	} else if strings.Index(command, "rain") != -1 {
		return getRainPrediction()

	}

	return "Повторите пожалуйста"
}

func getRainPrediction() string {
	p, _ := agents.GetRainProbability()

	var probability string
	if p.Probability > 0.8 {
		probability = "очень высокая вероятность осадков"
	} else if p.Probability > 0.5 {
		probability = "высокая вероятность осадков"
	} else if p.Probability > 0.1 {
		probability = "осадки вероятны"
	} else {
		probability = "осадки не ожидаются"
	}

	return fmt.Sprintf(
		"%s. В Киеве %s",
		randGreeting(),
		probability)
}

func getWindPrediction() string {
	f, _ := agents.GetCurrentWind()

	return fmt.Sprintf(
		"%s. В точке \"%s\" скорость ветра %4.1f м/c",
		randGreeting(),
		f.Items[0].Place.Title,
		f.Items[0].WindSpeed)
}

func randGreeting() string {
	return greetings[rand.Int31n(int32(len(greetings)))]
}
