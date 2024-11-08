package internal

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	e *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()
	e.GET("/:id", logNRedirect)

	return &Router{e}
}

func (r *Router) RunRouter(port string) {
	r.e.Logger.Fatal(r.e.Start(":" + port))
}

func logNRedirect(c echo.Context) error {
	clientIP := c.RealIP()
	fmt.Println(clientIP)
	return c.Redirect(http.StatusMovedPermanently, "http://goog.le")
}
