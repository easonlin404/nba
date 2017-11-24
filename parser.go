package nba

import (
	"log"
	"github.com/easonlin404/esrest"

)

func TodaysScores(){
	_, err := esrest.New().Debug(true).Get("https://data.nba.com/data/10s/v2015/json/mobile_teams/nba/2017/scores/00_todays_scores.json").Do()
	if err != nil {
		log.Fatal(err)
	}
	//TODO: parse response json
}