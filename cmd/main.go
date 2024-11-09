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

	println(linkService.Create(&models.LinkDataModel{"https://goog.le", 2}))
	println(linkService.Create(&models.LinkDataModel{"https://vk.com", 1}))
	println(linkService.Create(&models.LinkDataModel{"https://youtu.be", 3}))

	r := routers.NewLinkRouter(linkService)
	r.RunLinkRouter("8000")
}
