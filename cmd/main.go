package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moaabb/product-api/pkg/config"
	"github.com/moaabb/product-api/pkg/db"
	"github.com/moaabb/product-api/pkg/handlers"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	h := db.Init(c.DBUrl)
	ph := handlers.NewProductHandler(h)

	r := gin.Default()

	// Routes
	r.GET("/", ph.GetProducts)
	r.POST("/", ph.CreateProduct)

	r.Run(c.Port)

}
