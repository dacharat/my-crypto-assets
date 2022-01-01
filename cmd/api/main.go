package main

import (
	"github.com/dacharat/my-crypto-assets/cmd/api/route"
	"github.com/dacharat/my-crypto-assets/pkg/config"
)

func main() {
	config.NewConfig()

	router := route.NewRouter()
	router.Run()
}
