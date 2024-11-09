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

	println(linkService.Create(&models.LinkDataModel{"https://goog.le", 0, 0, 0}))
	println(linkService.Create(&models.LinkDataModel{"https://vk.com", 0, 0, 0}))
	println(linkService.Create(&models.LinkDataModel{"https://youtu.be", 0, 0, 0}))

	r := routers.NewLinkRouter(linkService)
	r.RunLinkRouter("8000")
}
