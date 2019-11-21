package main

import (
	//"./cmd"
	. "./parser"
	. "./story"
	"encoding/json"
	"log"
	"time"
)

var storyModel StoryModel

func main() {
	//cmd.Execute()

	app := NewFeatureApp()
	results := app.Start("tests/test.feature")
	jsonStr, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal(jsonStr, &storyModel)
	storyModel.StartDate, _ = time.Parse(time.RFC3339, results["startDate"])
	storyModel.EndDate, _ = time.Parse(time.RFC3339, results["endDate"])
	log.Println(storyModel)
}