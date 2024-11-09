package routers

import (
	"fmt"
	"net/http"

	"github.com/Dominux/gotcha/internal/services"
	"github.com/labstack/echo/v4"
)

type LinkRouter struct {
	e       *echo.Echo
	service *services.LinkService
}

func NewLinkRouter(service *services.LinkService) *LinkRouter {
	e := echo.New()
	return &LinkRouter{e, service}
}

func (r *LinkRouter) RunLinkRouter(port string) {
	r.e.GET("/:id", r.logNRedirect)
	r.e.Logger.Fatal(r.e.Start(":" + port))
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
