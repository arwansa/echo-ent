package main

import (
	"log"

	"github.com/arwansa/echo-ent/config"
	"github.com/arwansa/echo-ent/handler"
	"github.com/arwansa/echo-ent/middleware"
	"github.com/arwansa/echo-ent/repository"
	"github.com/arwansa/echo-ent/usecase"
	"github.com/labstack/echo"
)

func main() {
	cfg := config.Get()

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()
	config.SetClient(client)

	e := echo.New()
	s := middleware.NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle)

	userRepo := repository.NewUserRepository(client)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(e, userUsecase)

	e.Logger.Fatal(e.Start(cfg.Server.RESTPort))
}
