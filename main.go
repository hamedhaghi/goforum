package main

import (
	"github.com/hamedhaghi/goforum/api"
	"github.com/hamedhaghi/goforum/config"
)

func main() {
	api.Run(config.GetConfig())
}
