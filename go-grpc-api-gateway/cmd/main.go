package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moromin/go-grpc-api-gateway/pkg/auth"
	"github.com/moromin/go-grpc-api-gateway/pkg/config"
	"github.com/moromin/go-grpc-api-gateway/pkg/order"
	"github.com/moromin/go-grpc-api-gateway/pkg/product"
)

func main() {
	// check config
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
