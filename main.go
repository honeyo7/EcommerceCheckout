package main

import (
	"log"

	r "github.com/honeyo7/EcommerceCheckout/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	myRouter := gin.Default()
	r.RoutesCheckout(myRouter)

	log.Println("Server started on: http://localhost:8080")
	myRouter.Run(":8080")
}
