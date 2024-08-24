package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	url := os.Getenv("API_URL")
	key := os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Status code: ", res.StatusCode)
	}

	body,err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	// var data 
	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(data.Matches[0].HomeTeam.Name)

}
