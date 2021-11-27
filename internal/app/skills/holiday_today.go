package skills

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	holidayResult []string
)

func formatApiReadableDate(day int, month_original string) string {
	month := map[string]string{
		"January":   "yanvar",
		"February":  "fevral",
		"March":     "mart",
		"April":     "aprel",
		"May":       "may",
		"June":      "iyun",
		"July":      "iyul",
		"August":    "avgust",
		"September": "sentyabr",
		"October":   "oktyabr",
		"November":  "noyabr",
		"December":  "dekabr",
	}
	return fmt.Sprintf("%s/%d", month[month_original], day)
}

// Get Holiday ...
func GetHoliday(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	var holidayPageUrl []string
	_, month, day := time.Now().Date()
	currentDate := formatApiReadableDate(day, month.String())
	holidayPageUrl = append(holidayPageUrl, fmt.Sprintf("https://kakoysegodnyaprazdnik.ru/baza/%s", currentDate))
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: holidayPageUrl,
		ParseFunc: parseHolidayResult,
	}).Start()
	rand.Seed(time.Now().UnixNano())
	user := msg.From.UserName
	message := fmt.Sprintf("🥳🥳🥳\nГлавный праздник на сегодня: 🎉 %s 🎉\n⚡️Но %s сегодня будет отмечать:\n🎉 %s 🎉\n🥳🥳🥳", holidayResult[0], user, holidayResult[rand.Intn(len(holidayResult))])
	sendMsg := tgbotapi.NewMessage(msg.Chat.ID, message)
	bot.Send(sendMsg)
}

func parseHolidayResult(g *geziyor.Geziyor, r *client.Response) {
	findJSPath := "body > div.wrap > div:nth-child(2) > div > div"
	r.HTMLDoc.Find(findJSPath).Each(func(i int, s *goquery.Selection) {
		holidaysList := s.Find("div > span")
		holidaysList.Each(func(i int, s *goquery.Selection) {
			holidayResult = append(holidayResult, s.Text())
		})
	})
}
