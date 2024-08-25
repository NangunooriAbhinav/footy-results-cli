package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"footy_res/api"
	"footy_res/models"

	"github.com/joho/godotenv"
)

func filter(slice models.Data, league string, team string) models.Data {
	var filteredData models.Data
	for _, match := range slice.Matches {
		if match.Competition.Code == league {
			if team != "" {
				if strings.EqualFold(match.HomeTeam.ShortName, team) || strings.EqualFold(match.AwayTeam.ShortName, team) {
					filteredData.Matches = append(filteredData.Matches, match)
				}
			} else {
				filteredData.Matches = append(filteredData.Matches, match)
			}
		}
	}
	return filteredData
}

func formatDate(utcDate string) (string, error) {
	t, err := time.Parse(time.RFC3339, utcDate)
	if err != nil {
		return "", err
	}
	return t.Format("Monday, 02-Jan-2006 15:04"), nil
}

func printHelp() {
	fmt.Println("Usage: footy_res [flag] [team]")
	fmt.Println("Flags:")
	fmt.Println("  -p         Premier League")
	fmt.Println("  -f         Ligue 1")
	fmt.Println("  -u         UEFA Champions League")
	fmt.Println("  -b         Bundesliga")
	fmt.Println("  -w         World Cup")
	fmt.Println("  -i         Serie A")
	fmt.Println("  -h         Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  footy_res -p             # Show all Premier League matches")
	fmt.Println("  footy_res -p Arsenal     # Show all Arsenal matches in Premier League")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	url := os.Getenv("API_URL")
	key := os.Getenv("API_KEY")

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	if os.Args[1] == "-h" {
		printHelp()
		return
	}

	var league string
	switch os.Args[1] {
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
		fmt.Println("Invalid flag. Use -h for help.")
		return
	}

	var team string
	if len(os.Args) == 3 {
		team = os.Args[2]
	} else {
		team = ""
	}

	data, err := api.FetchData(url, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	data = filter(data, league, team)

	if len(data.Matches) == 0 {
		fmt.Println("No matches found.")
		return
	}

	fmt.Printf("%s", data.Matches[0].Competition.Name)
	if team != "" {
		fmt.Printf(" - Team: %s\n", team)
	}else{
		fmt.Println()
	}
	fmt.Println()

	for _, match := range data.Matches {
		formattedDate, err := formatDate(match.UtcDate)
		if err != nil {
			fmt.Println("Error formatting date:", err)
			formattedDate = match.UtcDate 
		}


		fmt.Printf("%s vs %s\n", match.HomeTeam.Name, match.AwayTeam.Name)
		fmt.Printf("Score: %d - %d\n", match.Score.FullTime.Home, match.Score.FullTime.Away)
		fmt.Printf("Date: %s\n", formattedDate)
		fmt.Println()
	}
}
