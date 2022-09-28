package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RespBody struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	http.HandleFunc("/data", HandleGetRequest)
	http.ListenAndServe(":8080", nil)
}

func HandleGetRequest(rw http.ResponseWriter, r *http.Request) {
	inputData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(inputData))
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(resp.Request.URL)
	// fmt.Println(resp.Header)
	// fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	// fmt.Println(body)
	log.Println(string(body))
	var respBody RespBody
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.MarshalIndent(respBody, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
}

// fd656012187b434eb6c152050222709
