package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Host string
	Port string
}

type Server struct {
	Server *echo.Echo
}

func NewServer(cfg Config) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "../static/index.html")

	return e.Start(cfg.Host)

}
