package config

import (
	"api/pkg/handlers"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Server struct {
	config Config
	engine *gin.Engine
	httpServer *http.Server
}

func NewServer(config Config) *Server {
	handler := &handlers.Handler{}
	engine := handler.InitRoutes()
	
	return &Server{
		config: config,
		engine: engine,
	}
}

func (s *Server) Run() error {
	s.httpServer = &http.Server{
		Addr: fmt.Sprintf(":%s", s.config.Port),
		Handler: s.engine,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: s.config.Timeout,
		WriteTimeout: s.config.Timeout,
	}

	log.Printf("Server is running on port %s\n", s.config.Port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx *gin.Context) error {
	log.Println("Shutdown server...")
	return s.httpServer.Shutdown(ctx)
}


