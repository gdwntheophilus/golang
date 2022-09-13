package main

type ResponseObject struct {
	Status bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int    `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat`
	Text      string `json:"text"`
}

type From struct {
	FromId          int    `json:"id"`
	IsBot           bool   `json:"is_bot"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	LanguageEncoded string `json:"language_encoded"`
}

type Chat struct {
	ChatId    int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}

type BotMessage struct {
	ChatId        int    `json:"chat_id"`
	Text          string `json:"text"`
	ReplyMesageId int    `json:"reply_to_message_id"`
}
