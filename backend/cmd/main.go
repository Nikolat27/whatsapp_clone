package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"whatsapp_clone/internal/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("loading env variables. error: %s", err))
	}
	
	app, err := api.NewApp()
	if err != nil {
		panic(fmt.Errorf("creating a new app. error: %s", err))
	}
	
	err = app.RunServer()
	if err != nil {
		panic(fmt.Errorf("running the http server. error: %s", err))
	}
}
