package skills

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/aggyomfg/creampie-bot/internal/app/model"
	"github.com/aggyomfg/creampie-bot/internal/app/store"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
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
func GetHoliday(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, store store.Store, logger *logrus.Logger) {
	log = logger
	rand.Seed(time.Now().UnixNano())
	user := *msg.From
	todayHolidays, err := store.HolidayToday().GetAllHolidaysToday()
	if err != nil {
		log.Debugf("%s: No holidays in store, getting it now...", err)
	}
	if len(todayHolidays) == 0 {
		getHolidaysFromAPI()
		store.HolidayToday().SetAllHolidaysToday(holidayResult)
		todayHolidays, _ = store.HolidayToday().GetAllHolidaysToday()
	}
	mainHoliday := todayHolidays[0]
	userHoliday, err := store.HolidayToday().FindByUser(user)
	if err != nil {
		log.Debugf("%s: Holiday for user %s not found! Creating now...", err, user.UserName)
		store.HolidayToday().Create(&model.HolidayToday{
			User:    user,
			Holiday: todayHolidays[rand.Intn(len(todayHolidays))],
		})
		userHoliday, _ = store.HolidayToday().FindByUser(user)
	}
	message := fmt.Sprintf("🥳🥳🥳\nГлавный праздник на сегодня: 🎉 %s 🎉\n⚡️Но %s сегодня будет отмечать:\n🎉 %s 🎉\n🥳🥳🥳", mainHoliday, user.UserName, userHoliday)
	sendMsg := tgbotapi.NewMessage(msg.Chat.ID, message)
	bot.Send(sendMsg)
}

func getHolidaysFromAPI() {
	var holidayPageUrl []string
	_, month, day := time.Now().Date()
	currentDate := formatApiReadableDate(day, month.String())
	holidayPageUrl = append(holidayPageUrl, fmt.Sprintf("https://kakoysegodnyaprazdnik.ru/baza/%s", currentDate))
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: holidayPageUrl,
		ParseFunc: parseHolidayResult,
	}).Start()
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
