package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
)

type Config struct {
	ServerAddr string `env:"SERVER_ADDR,required"`
}

func main() {
	cfg := new(Config)
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("could not parse config: %v", err)
		return
	}

	engine := gin.New()

	engine.Handle(http.MethodGet, "/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	log.Println("starting server on", cfg.ServerAddr)
	log.Println("shutdown error", engine.Run(cfg.ServerAddr))
}
