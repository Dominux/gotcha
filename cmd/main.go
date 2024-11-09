package main

import (
	"github.com/Dominux/gotcha/internal/models"
	"github.com/Dominux/gotcha/internal/repositories"
	"github.com/Dominux/gotcha/internal/routers"
)

func main() {
	// creating urls repo
	urlsRepo := repositories.NewLinkRepository()

	println(urlsRepo.Create(&models.LinkDataModel{"https://goog.le", 0, 0, 0}))
	println(urlsRepo.Create(&models.LinkDataModel{"https://vk.com", 0, 0, 0}))
	println(urlsRepo.Create(&models.LinkDataModel{"https://youtu.be", 0, 0, 0}))

	r := routers.NewRouter(urlsRepo)
	r.RunRouter("8000")
}
