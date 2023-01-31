package main

import (
	"cool/config"
	"cool/router"
)

func main() {
	defer config.CloseDatabaseConnection(config.DB)
	defer config.RedisPool.Close()

	router.InitApp()
}
