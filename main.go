package main

import (
	"fmt"
	"os"

	"footy_res/api"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	url := os.Getenv("API_URL")
	key := os.Getenv("API_KEY")

	data,err := api.FetchData(url,key)

	fmt.Println(data.Matches[0].HomeTeam.Name)
}
