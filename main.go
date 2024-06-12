package main

import (
	"blog/configs"
	"blog/routers"
	"fmt"
)

func main() {
	var router = routers.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
	configs.InitConfig()
	fmt.Println("Server is running on: http://localhost:8080")
}
