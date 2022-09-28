package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SlackResponse struct {
	Challenge string `json:"challenge"`
}

func main() {
	http.HandleFunc("/slackbot-event-subs", HandleSlackbotEventSubs)
	http.ListenAndServe(":8080", nil)
}

func HandleSlackbotEventSubs(w http.ResponseWriter, r *http.Request) {

	getData, err := http.Get("https://api.weatherapi.com/v1/current.json?key=fd656012187b434eb6c152050222709&q=Bangalore&aqi=no")
	if err != nil {
		fmt.Println(err)
	}

	ioData, _ := ioutil.ReadAll(getData.Body)
	fmt.Println(string(ioData))

	fmt.Println(r.Body)
	var slackResponse SlackResponse
	slackResponse.Challenge = "This is the response"
	data, err := json.Marshal(slackResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func srting(ioData []byte) {
	panic("unimplemented")
}
