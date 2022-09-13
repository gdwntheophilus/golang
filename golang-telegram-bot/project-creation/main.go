package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	botToken := "5665585256:AAGlVYVVvVIjyD-u-XHN5i3vn_mCN_YRNzU"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	fmt.Println(botUrl)
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Fatalln(err)
		}
		for _, update := range updates {
			fmt.Println(update)
			err = sendResponse(botUrl, update)
			offset = update.UpdateId + 1
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resposebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseObject ResponseObject
	err = json.Unmarshal(resposebody, &responseObject)
	if err != nil {
		return nil, err
	}

	fmt.Println(responseObject.Result)

	return responseObject.Result, nil
}

func sendResponse(botUrl string, update Update) error {
	fmt.Println(update)
	// var jsonStr = []byte(`{"text":"Hello!"}`)
	// values := map[string]string{"chat_id": strconv.Itoa(update.UpdateId), "text": update.Message.Text}
	// buff, _ := json.Marshal(values)
	var botResponse BotMessage
	botResponse.ChatId = update.Message.Chat.ChatId //update.UpdateId
	botResponse.Text = update.Message.Text
	botResponse.ReplyMesageId = update.Message.MessageId

	fmt.Println(botResponse)
	buff, err := json.Marshal(botResponse)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buff)
	request, err := http.NewRequest("POST", botUrl+"/sendMessage", bytes.NewBuffer(buff))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	// request, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buff))
	// request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	return nil

}

//https://api.telegram.org/bot5665585256:AAGlVYVVvVIjyD-u-XHN5i3vn_mCN_YRNzU/getMe
//https://api.telegram.org/bot5665585256:AAGlVYVVvVIjyD-u-XHN5i3vn_mCN_YRNzU/getUpdates
