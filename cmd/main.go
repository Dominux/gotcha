package main

import (
	"os"
	"strconv"

	"github.com/Dominux/gotcha/internal/common"
	"github.com/Dominux/gotcha/internal/repositories"
	"github.com/Dominux/gotcha/internal/routers"
	"github.com/Dominux/gotcha/internal/services"
)

func main() {
	// creating config from env vars
	var config common.Config
	{
		tgRateLimitInMs, err := strconv.ParseUint(os.Getenv("TG_RATE_LIMIT_IN_MS"), 10, 32)
		if err != nil {
			panic(err)
		}

		tgUserId, err := strconv.ParseUint(os.Getenv("TG_USER_ID"), 10, 32)
		if err != nil {
			panic(err)
		}

		config = common.Config{
			TGBotToken:      os.Getenv("TG_BOT_TOKEN"),
			TGRateLimitInMs: tgRateLimitInMs,
			TGUserId:        tgUserId,
			URLBase:         os.Getenv("URL_BASE"),
			NotFoundUrl:     os.Getenv("NOT_FOUND_URL"),
		}
	}

	// creating link repo
	linkRepo := repositories.NewLinkRepository()

	// creating link service
	linkService := services.NewLinkService(linkRepo, &config.URLBase)

	// creating telegram bot service
	tgBotService := services.NewTelegramBotService(linkService, &config)

	// adding removing cycle
	go linkService.RunLinksRemovingCycle()

	// adding tg bot
	go tgBotService.RunCheckingUpdatesCycle()

	// creating main router and running it
	r := routers.NewMainRouter()
	r.AddLinkRouter(linkService, tgBotService, &config.NotFoundUrl)

	r.RunRouter("8000")
}
