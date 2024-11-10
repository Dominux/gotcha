package models

type TelegramApiUpdateModel struct {
	UpdateId uint64                  `json:"update_id"`
	Message  TelegramApiMessageModel `json:"message"`
}
