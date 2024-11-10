package main

import (
	"github.com/Dominux/gotcha/internal/models"
	"github.com/Dominux/gotcha/internal/repositories"
	"github.com/Dominux/gotcha/internal/routers"
	"github.com/Dominux/gotcha/internal/services"
)

func main() {
	// creating link repo
	linkRepo := repositories.NewLinkRepository()

	// creating link service
	linkService := services.NewLinkService(linkRepo)

	// creating links
	println(linkService.Create(&models.LinkDataModel{"https://goog.le", 2, 1}))
	println(linkService.Create(&models.LinkDataModel{"https://vk.com", 1, 2}))
	println(linkService.Create(&models.LinkDataModel{"https://youtu.be", 3, 2}))

	// adding removing cycle
	go linkService.RunLinksRemovingCycle()

	// creating main router and running it
	r := routers.NewMainRouter()
	r.AddLinkRouter(linkService)

	r.RunRouter("8000")
}
