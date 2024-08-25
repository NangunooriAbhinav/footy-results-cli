# Footy Res CLI

**Footy Res CLI** is a command-line application that allows you to fetch and filter football match results based on different leagues and teams. The application retrieves data from a football API and provides the results in a user-friendly format.

## Features

- Fetch football match results from various leagues.
- Filter results by league and team.
- Display match details including teams, scores, and formatted match dates.
- Supports multiple leagues, including Premier League, Ligue 1, UEFA Champions League, Bundesliga, World Cup, and Serie A.

## Requirements

- Go 1.16 or higher
- A valid API key for the football data source (stored in a `.env` file).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/footy_res.git
   cd footy_res
   ```

2. Create a `.env` file in the root directory and add your API details:

   ```bash
   API_URL="https://api.football-data.org/v4/matches"
   API_KEY="your_api_key_here"
   ```

3. Build the application:

   ```bash
   go build -o footy_res
   ```

## Usage

Run the application with different flags to filter the results by league and optionally by team.

### Fetch all matches for a specific league:

```bash
./footy_res -p  # Premier League
./footy_res -f  # Ligue 1
./footy_res -u  # UEFA Champions League
./footy_res -b  # Bundesliga
./footy_res -w  # World Cup
./footy_res -i  # Serie A
```

### Fetch all matches for a specific team within a league:

```bash
./footy_res -p Arsenal  # Premier League matches for Arsenal
./footy_res -b Bayern   # Bundesliga matches for Bayern Munich
```

### Display help information:

```bash
./footy_res -h
```

### Example Output

```bash
Results for Premier League - Team: Arsenal

Arsenal FC vs Chelsea FC
Score: 3 - 1
Date: Saturday, 24-Aug-2024 16:30
----------------------------------------
```

## Project Structure

```
footy_res/
├── api/
│   ├── fetch.go   # Contains API fetching logic
├── models/
│   ├── types.go   # Contains all data models and types
├── main.go        # Main entry point for the application
├── .env           # Environment variables (not included in the repo)
└── README.md      # Project documentation
```

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Football Data API](https://www.football-data.org/documentation) for providing the match data.
- [GoDotEnv](https://github.com/joho/godotenv) for handling environment variables.
