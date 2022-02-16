package server

import (
	"errors"
	"fmt"
	"strider-challenge/app/router"
	"strider-challenge/app/router/routehomepage"
	"strider-challenge/domain/service"
	"strider-challenge/infra/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type IRouter interface {
	Register(e *echo.Echo)
}

type Server struct {
	srv         *echo.Echo
	cfg         *config.Config
	routes      []IRouter
	middlewares []echo.MiddlewareFunc
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		srv: echo.New(),
		cfg: cfg,
	}
}

func (s *Server) addRoute(route IRouter) {
	s.routes = append(s.routes, route)
}

func (s *Server) addMiddleware(middleware echo.MiddlewareFunc) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) registerRoutes() {
	for _, route := range s.routes {
		route.Register(s.srv)
	}
}

func (s *Server) registerMiddlewares() {
	for _, middleware := range s.middlewares {
		s.srv.Use(middleware)
	}
}

func (s *Server) run() (err error) {

	s.srv.HideBanner = true
	s.srv.Debug = s.cfg.ServerDebug

	s.registerRoutes()
	s.registerMiddlewares()

	err = s.srv.Start(fmt.Sprintf(":%d", s.cfg.ServerPort))
	if err != nil {
		return errors.New("server: fail to start echo")
	}

	return nil
}

func (s *Server) InitServer(svc *service.Service) (err error) {

	logrus.Info("runninng server at localhost:", s.cfg.ServerPort)

	baseRouter := router.NewBaseRouter()

	homePageSvc := service.NewHomePageService(*svc)
	homePageCtrl := routehomepage.NewController(homePageSvc)
	homePageRoute := routehomepage.NewRoute("homepage", homePageCtrl)

	s.addRoute(baseRouter)
	s.addRoute(homePageRoute)

	s.addMiddleware(middleware.Logger())

	err = s.run()
	return err
}
