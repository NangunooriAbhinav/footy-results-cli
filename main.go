package main

import (
	"fmt"
	"os"

	"footy_res/api"
	"footy_res/models"

	"github.com/joho/godotenv"
)

func filter (slice models.Data,league string,team string) models.Data{
	var filteredData models.Data
	for _,match := range slice.Matches{
		if match.Competition.Code == league{
			if team != ""{
				if match.HomeTeam.ShortName == team || match.AwayTeam.ShortName == team{
					filteredData.Matches = append(filteredData.Matches,match)
				}
			}else{
				filteredData.Matches = append(filteredData.Matches,match)
			}
		}
	}
	return filteredData
}

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	url := os.Getenv("API_URL")
	key := os.Getenv("API_KEY")

	if len(os.Args) < 2 {
		fmt.Println("Usage : footy_res [flag]  Help: footy_res -h")
		return
	}
	
	var league string
	switch os.Args[1]{
		case "-p":
			league = "PL"
		case "-f":
			league = "FL1"
		case "-u":
			league = "ucl"
		case "-b":
			league = "BL1"
		case "-w":
			league = "world-cup"
		case "-i":
			league = "SA"
		default:
			league = ""
	}
	
	var team string
	if len(os.Args) == 3{
		team = os.Args[2]
	}else{
		team = ""
	}


	data,err := api.FetchData(url,key)
	if err != nil {
		fmt.Println(err)
		return
	}

	data = filter(data,league,team)

	fmt.Println(data.Matches[0].Competition.Name)
	fmt.Println()
	for _,match := range data.Matches{
		fmt.Println(match.HomeTeam.Name," vs ",match.AwayTeam.Name)
		fmt.Println(match.Score.FullTime.Home," - ",match.Score.FullTime.Away)
		fmt.Println()
	}
}
