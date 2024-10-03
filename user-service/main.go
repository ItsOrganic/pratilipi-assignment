package main

import (
	"log"
	"user-service/db"
	"user-service/handler"
	"user-service/metrics"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	err = db.Connect("mongodb://mongodb:27017", "user-service", "users")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// utils.InitMQ()
	metrics.Init()
	// defer utils.CloseMQ()

	// err = utils.InitMQ("amqp://rabbitmq:5672")
	// if err != nil {
	// 	log.Fatalf("Error connecting to RabbitMQ: %v", err)
	// }
	// defer utils.CloseMQ()

	router := gin.Default()
	router.Use(middleware.PrometheusMiddleware())
	router.GET("/metrics", metrics.PrometheusHandler)
	router.POST("/register", handler.RegisterUser)
	router.POST("/login", handler.AuthenticateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/user/:id", handler.GetUser)
	router.GET("/profile/:id", handler.GetProfile)
	router.PUT("/profile/:id", handler.UpdateProfile)
	router.Run(":8081")
}
