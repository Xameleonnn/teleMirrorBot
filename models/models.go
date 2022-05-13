package models

type Update struct {
	Update_id int     `json:"update_id"`
	Message   Message `json:"message"`
}

type Message struct {
	Message_id int    `json:"message_id"`
	Chat       Chat   `json:"chat"`
	Text       string `json:"text"`
}

type Chat struct {
	Chat_id int `json:"id"`
}

type RestResponse struct {
	Updates []Update `json:"result"`
}

type BotMessage struct {
	Chat_id int    `json:"chat_id"`
	Text    string `json:"text"`
}
