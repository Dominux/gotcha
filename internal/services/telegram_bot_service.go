package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Dominux/gotcha/internal/common"
	"github.com/Dominux/gotcha/internal/models"
)

const linksCreatedMsgTemplate = `
Links created:

<i>gotcha</i>: <code>%s</code>

Integrations:
<i>shorturl.at</i>: <code>%s</code>
`

const gotchaMsgTemplate = `
<b>Gotcha</b>

<b>original link:</b> <i>%s</i>
<b>IP:</b> <code>%s</code>
<b>User Agent:</b> <i>%s</i>
`

type TelegramBotService struct {
	linkService  *LinkService
	botToken     *string
	rateLimit    time.Duration
	userId       uint64
	nextUpdateId uint64
}

func NewTelegramBotService(linkService *LinkService, config *common.Config) *TelegramBotService {
	tgRateLimit := time.Millisecond * time.Duration(config.TGRateLimitInMs)
	return &TelegramBotService{linkService, &config.TGBotToken, tgRateLimit, config.TGUserId, 0}
}

func (s *TelegramBotService) RunCheckingUpdatesCycle() {
	println("Ran checking updates cycle")

	for {
		s.processCheckingUpdatesIteration()
	}
}

func (s *TelegramBotService) SendGotcha(destinationLink, ip string, userAgent string) {
	text := fmt.Sprintf(gotchaMsgTemplate, destinationLink, ip, userAgent)
	s.sendMessage(s.userId, text)
}

func (s *TelegramBotService) processCheckingUpdatesIteration() {
	time.Sleep(s.rateLimit)

	updatesURL := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&limit=10&allowed_updates=[\"message\"]", *s.botToken, s.nextUpdateId)

	res, err := http.Get(updatesURL)
	if err != nil {
		s.raiseTgApiError(err)
		return
	}

	defer res.Body.Close()

	var body models.TelegramApiResultModel
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		println(err.Error())
	}

	for _, update := range body.Result {
		s.processMessage(update.Message)

		s.nextUpdateId = update.UpdateId + 1
	}
}

func (s *TelegramBotService) processMessage(msg models.TelegramApiMessageModel) {
	// Restricting messages from others
	if msg.From.Id != s.userId {
		return
	}

	if msg.Text == "/start" {
		text := "Sup! Send me a link and I'll provide the gotcha link"
		if err := s.sendMessage(s.userId, text); err != nil {
			s.raiseTgApiError(err)
		}
		return
	}

	// validating url
	if _, err := url.ParseRequestURI(msg.Text); err != nil {
		text := "This message is not a valid URL"
		if err := s.sendMessage(s.userId, text); err != nil {
			s.raiseTgApiError(err)
		}
		return
	}

	// creating gotcha link
	linkData := models.LinkDataModel{msg.Text, 1, 7}
	gotchaLink, shortUrlLink := s.linkService.Create(&linkData)
	msgToSend := fmt.Sprintf(linksCreatedMsgTemplate, gotchaLink, shortUrlLink)

	if err := s.sendMessage(s.userId, msgToSend); err != nil {
		s.raiseTgApiError(err)
	}
}

func (s *TelegramBotService) sendMessage(chatId uint64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s&parse_mode=html", *s.botToken, chatId, url.QueryEscape(text))
	_, err := http.Get(url)
	return err
}

func (s TelegramBotService) raiseTgApiError(err error) {
	println("[Telegram API Error]", err.Error())
}
