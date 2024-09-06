package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/erfanwd/exchangeto/config"
	"github.com/erfanwd/exchangeto/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CurrencyPrices struct {
	USD float64 `json:"USD"`
}

type ExchangePrices map[string]CurrencyPrices

var Exchanges ExchangePrices

func Trigger(bot *tgbotapi.BotAPI) {
	exchanges := "BTC,ETH,XRP,USDT"
	apiURL := config.Config("EXCHANGE_API_URL") + "&fsyms=" + exchanges + "&api_key=" + config.Config("EXCHANGE_API_TOKEN")
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fetchData(apiURL, bot)
			}
		}
	}()
}

func fetchData(url string, bot *tgbotapi.BotAPI) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	var prices ExchangePrices
	if err := json.Unmarshal(body, &prices); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	Exchanges = prices

	reminders, _ := repositories.GetAllReminders()

	for _, reminder := range reminders {

		price := Exchanges[reminder.Exchange.Code]

		if reminder.Strategy == "higher" && reminder.Value <= int64(price.USD) {
			valueStr := strconv.FormatInt(reminder.Value, 10)
			priceStr := strconv.FormatFloat(price.USD, 'f', 2, 64)
			text := "قیمت " + reminder.Exchange.Name + " از " + valueStr + " دلار گذشت و به " + priceStr + " دلار رسید."

			msg := tgbotapi.NewMessage(reminder.User.ChatId, text)

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
				continue
			}
			repositories.DeleteReminder(reminder)
		}

		if reminder.Strategy == "lower" && reminder.Value >= int64(price.USD) {
			valueStr := strconv.FormatInt(reminder.Value, 10)
			priceStr := strconv.FormatFloat(price.USD, 'f', 2, 64)
			text := "قیمت " + reminder.Exchange.Name + " از " + valueStr + " دلار پایین تر اومد و به " + priceStr + " دلار رسید."

			msg := tgbotapi.NewMessage(reminder.User.ChatId, text)

			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
				continue
			}
			repositories.DeleteReminder(reminder)
		}

	}

}
