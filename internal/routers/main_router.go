package routers

import (
	"github.com/Dominux/gotcha/internal/services"
	"github.com/labstack/echo/v4"
)

type MainRouter struct {
	e          *echo.Echo
	linkRouter *LinkRouter
}

func NewMainRouter() *MainRouter {
	e := echo.New()
	return &MainRouter{e, nil}
}

func (r *MainRouter) AddLinkRouter(service *services.LinkService) {
	linkGroup := r.e.Group("/l")
	r.linkRouter = newLinkRouter(linkGroup, service)
}

func (r *MainRouter) RunRouter(port string) {
	r.e.Logger.Fatal(r.e.Start(":" + port))
}