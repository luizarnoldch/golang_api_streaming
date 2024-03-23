package main

import (
	"main/app"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.Info("Server Starting")
	app.Start()
	log.Info("Server Finished")
}