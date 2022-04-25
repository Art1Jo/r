package main

import (
	"log"
	handlers2 "reddit/pck/handlers"
	"reddit/pck/repositories"
	services2 "reddit/pck/services"
)

func main() {
	serverInstance := new(Server)
	postgresConfig := repositories.Config{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	}
	database, err := repositories.NewPostgresDB(postgresConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repositories.NewRepository(database)
	services := services2.NewService(repos)
	handler := handlers2.NewHandler(services)
	runServer(serverInstance, handler)
}

func runServer(serverInstance *Server, handlerLayer *handlers2.Handler) {
	port := "8080"
	router := handlerLayer.InitRouters()

	if err := serverInstance.Run(port, router); err != nil {
		log.Fatal(err.Error())
	}

	log.Print("server started successfully")
}
