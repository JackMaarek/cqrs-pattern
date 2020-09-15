package main

import (
	"github.com/JackMaarek/cqrsPattern/application/models"
	routes "github.com/JackMaarek/cqrsPattern/application/router"
	"github.com/JackMaarek/cqrsPattern/application/validators"
	"github.com/JackMaarek/cqrsPattern/domain"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"log"
)

type config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	models.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
	models.MakeMigrations()
	router := gin.Default()
	routes.SetupRouter(router)
	//go bus.HandleRequest(&channels.C)
	validators.InitValidator()
	domain.InitBuses()

	log.Fatal(router.Run(":8080"))
}
