package main

import (
	"flag"
	"fmt"
	"github.com/Erickype/HexagonalArchitecture/internal/adapters/handler"
	"github.com/Erickype/HexagonalArchitecture/internal/adapters/repository"
	"github.com/Erickype/HexagonalArchitecture/internal/core/services"
	"github.com/gin-gonic/gin"
)

var (
	repo      = flag.String("db", "postgres", "Database for storing messages")
	redisHost = "redis-17424.c11.us-east-1-2.ec2.cloud.redislabs.com:17424"
	svc       *services.MessengerService
)

func main() {
	flag.Parse()

	fmt.Printf("Application running using %s\n", *repo)
	switch *repo {
	case "redis":
		store := repository.NewMessengerRedisRepository(redisHost)
		svc = services.NewMessengerService(store)
	default:
		store := repository.NewMessengerPostgresRepository()
		svc = services.NewMessengerService(store)
	}

	InitRoutes()
}

func InitRoutes() {
	router := gin.Default()
	routesHandler := handler.NewHttpHandler(*svc)
	router.GET("/messages/:id", routesHandler.ReadMessage)
	router.GET("/messages", routesHandler.ReadMessages)
	router.POST("/messages", routesHandler.SaveMessage)
	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
