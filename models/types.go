package models

type Data struct {
	Filters   Filters   `json:"filters"`
	ResultSet ResultSet `json:"resultSet"`
	Matches   []Match   `json:"matches"`
}

type Filters struct {
	DateFrom   string `json:"dateFrom"`
	DateTo     string `json:"dateTo"`
	Permission string `json:"permission"`
}

type ResultSet struct {
	Count        int    `json:"count"`
	Competitions string `json:"competitions"`
	First        string `json:"first"`
	Last         string `json:"last"`
	Played       int    `json:"played"`
}

type Match struct {
	Area        Area        `json:"area"`
	Competition Competition `json:"competition"`
	Season      Season      `json:"season"`
	ID          int         `json:"id"`
	UtcDate     string      `json:"utcDate"`
	Status      string      `json:"status"`
	Matchday    int         `json:"matchday"`
	Stage       string      `json:"stage"`
	Group       string      `json:"group"`
	LastUpdated string      `json:"lastUpdated"`
	HomeTeam    Team        `json:"homeTeam"`
	AwayTeam    Team        `json:"awayTeam"`
	Score       Score       `json:"score"`
	Odds        Odds        `json:"odds"`
	Referees    []Referee   `json:"referees"`
}

type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type Competition struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}

type Season struct {
	ID             int         `json:"id"`
	StartDate      string      `json:"startDate"`
	EndDate        string      `json:"endDate"`
	CurrentMatchday int        `json:"currentMatchday"`
	Winner         interface{} `json:"winner"`
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Tla       string `json:"tla"`
	Crest     string `json:"crest"`
}

type Score struct {
	Winner    string     `json:"winner"`
	Duration  string     `json:"duration"`
	FullTime  TimeScore  `json:"fullTime"`
	HalfTime  TimeScore  `json:"halfTime"`
}

type TimeScore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Odds struct {
	Msg string `json:"msg"`
}

type Referee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Nationality string `json:"nationality"`
}

type Error struct {
	Message string `json:"message"`
}