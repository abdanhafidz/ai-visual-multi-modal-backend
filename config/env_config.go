package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
var TCP_ADDRESS string
var LOG_PATH string
var HOST_ADDRESS string
var HOST_PORT string
var EMAIL_VERIFICATION_DURATION int
var OPEN_AI_API_KEY string
var REPLICATE_API_KEY string

func InitializeEnv() {
	godotenv.Load()
	HOST_ADDRESS = os.Getenv("HOST_ADDRESS")
	HOST_PORT = os.Getenv("HOST_PORT")
	TCP_ADDRESS = HOST_ADDRESS + ":" + HOST_PORT
	LOG_PATH = os.Getenv("LOG_PATH")
	EMAIL_VERIFICATION_DURATION, _ = strconv.Atoi(os.Getenv("EMAIL_VERIFICATION_DURATION"))
	OPEN_AI_API_KEY = os.Getenv("OPEN_AI_API_KEY")
	REPLICATE_API_KEY = os.Getenv("REPLICATE_API_KEY")
}