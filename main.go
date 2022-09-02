package main

import (
	"cool/config"
	"cool/router"
)

func main() {
	defer config.CloseDatabaseConnection(config.DB)

	router.InitApp()
}
