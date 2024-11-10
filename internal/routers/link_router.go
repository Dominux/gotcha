package routers

import (
	"net/http"

	"github.com/Dominux/gotcha/internal/services"
	"github.com/labstack/echo/v4"
)

type LinkRouter struct {
	e            *echo.Group
	service      *services.LinkService
	tgBotService *services.TelegramBotService
	notFoundUrl  *string
}

func newLinkRouter(g *echo.Group, service *services.LinkService, tgBotService *services.TelegramBotService, notFoundUrl *string) *LinkRouter {
	router := &LinkRouter{g, service, tgBotService, notFoundUrl}

	g.GET("/:id", router.logNRedirect)

	return router
}

func (r *LinkRouter) logNRedirect(c echo.Context) error {
	id := c.Param("id")

	// trying to get link by id
	link, err := r.service.Get(id)
	if err != nil {
		// return c.String(http.StatusNotFound, "Not Found")
		return c.Redirect(http.StatusMovedPermanently, *r.notFoundUrl)
	}

	clientIP := c.RealIP()
	r.tgBotService.SendGotcha(link.DestinationLink, clientIP)

	return c.Redirect(http.StatusMovedPermanently, link.DestinationLink)
}
