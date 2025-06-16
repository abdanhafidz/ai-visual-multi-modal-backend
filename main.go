package main

import (
	"fmt"

	config "github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	router "github.com/abdanhafidz/ai-visual-multi-modal-backend/router"
)

func main() {
	fmt.Println("Server started on ", config.TCP_ADDRESS, ", port :", config.HOST_PORT)
	router.StartService()

}
