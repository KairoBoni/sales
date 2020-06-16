package server

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml: port`
}

type Server struct {
	route *echo.Echo
	port  string
}

func getConfig(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}

func NewServer(configPath string) (*Server, error) {
	cfg, err := getConfig(configPath)
	if err != nil {
		return nil, err
	}
	s := &Server{
		route: echo.New(),
		port:  cfg.Port,
	}
	s.setupRoutes()

	return s, nil

}

func (s *Server) setupRoutes() {
	s.route.Use(middleware.Logger())
	s.route.Use(middleware.Recover())

	s.route.File("/", "static/index.html")
}

func (s *Server) Run() error {
	return s.route.Start(":" + s.port)
}
