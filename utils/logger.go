package utils

import (
	"fmt"
	"log"
	"os"

	config "github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
)

func LogError(errorLogged error) {
	fmt.Println("There is an error!")
	file, err := os.OpenFile(config.LOG_PATH+"/error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Error Log :", errorLogged)
	log.SetOutput(file)

	log.Println("Error Log :", errorLogged)
}
