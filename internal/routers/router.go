package routers

import (
	"fmt"
	"net/http"

	"github.com/Dominux/gotcha/internal/repositories"
	"github.com/labstack/echo/v4"
)

type Router struct {
	e    *echo.Echo
	repo *repositories.LinkRepository
}

func NewRouter(repo *repositories.LinkRepository) *Router {
	e := echo.New()
	return &Router{e, repo}
}

func (r *Router) RunRouter(port string) {
	r.e.GET("/:id", r.logNRedirect)
	r.e.Logger.Fatal(r.e.Start(":" + port))
}

func (r *Router) logNRedirect(c echo.Context) error {
	id := c.Param("id")

	// trying to get link by id
	link, err := r.repo.Get(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	clientIP := c.RealIP()
	fmt.Println(clientIP)
	return c.Redirect(http.StatusMovedPermanently, link.DestinationLink)
}
