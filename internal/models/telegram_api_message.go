package models

type TelegramApiMessageModel struct {
	MessageId uint64               `json:"message_id"`
	From      TelegramApiUserModel `json:"from"`
	Text      string               `json:"text"`
}
