package routers

import (
	"fmt"
	"net/http"

	"github.com/Dominux/gotcha/internal/services"
	"github.com/labstack/echo/v4"
)

type LinkRouter struct {
	e       *echo.Group
	service *services.LinkService
}

func newLinkRouter(g *echo.Group, service *services.LinkService) *LinkRouter {
	router := &LinkRouter{g, service}

	g.GET("/:id", router.logNRedirect)

	return router
}

func (r *LinkRouter) logNRedirect(c echo.Context) error {
	id := c.Param("id")

	// trying to get link by id
	link, err := r.service.Get(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	clientIP := c.RealIP()
	fmt.Println(clientIP)
	return c.Redirect(http.StatusMovedPermanently, link.DestinationLink)
}
