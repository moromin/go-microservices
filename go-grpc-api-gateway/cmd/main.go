package main

import (
	"fmt"
	"log"

	"github.com/moromin/go-grpc-api-gateway/pkg/config"
)

func main() {
	// check config
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", c)
}
