package server

import (
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
)

type config struct {
	Port      string `yaml:"port"`
	IndexPath string `yaml:"indexPath"`
}

type Server struct {
	route *echo.Echo
	cfg   *config
}

func getConfig(path string) (*config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &config{}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, err
	}
	fmt.Println(cfg)

	return cfg, nil

}

func NewServer(configPath string) (*Server, error) {
	cfg, err := getConfig(configPath)
	if err != nil {
		return nil, err
	}
	s := &Server{
		route: echo.New(),
		cfg:   cfg,
	}
	s.setupRoutes()

	return s, nil

}

func (s *Server) setupRoutes() {
	s.route.Use(middleware.Logger())
	s.route.Use(middleware.Recover())

	s.route.File("/", s.cfg.IndexPath)
}

func (s *Server) Run() error {
	return s.route.Start(":" + s.cfg.Port)
}
