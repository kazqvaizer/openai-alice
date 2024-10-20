package main

import (
	"alice-talking/dialog"
	"alice-talking/server"
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

func main() {

	err := env.Load("../.env")
	if err != nil {
		fmt.Printf("Problem with `.env` file: %v\n", err)
		return
	}

	var DialogConfig dialog.DialogConfig
	DialogConfig.ApiKey = os.Getenv("OPENAI_API_KEY")

	if DialogConfig.ApiKey == "" {
		fmt.Printf("API Key not set!\n")
		return
	}

	var ServerConfig server.ServerConfig
	ServerConfig.Host = os.Getenv("HOST")
	ServerConfig.Port = os.Getenv("PORT")
	ServerConfig.Token = os.Getenv("TOKEN")

	server.StartServer(ServerConfig, DialogConfig)

}
