package dto

type GetSettingsRequest struct {
	IdInstance       string `json:"id_instance"`
	ApiTokenInstance string `json:"api_token_instance"`
}

type GetStateInstanceRequest struct {
	IdInstance       string `json:"id_instance"`
	ApiTokenInstance string `json:"api_token_instance"`
}

type SendMessageRequest struct {
	IdInstance       string `json:"id_instance"`
	ApiTokenInstance string `json:"api_token_instance"`
	ChatID           string `json:"chat_id"`
	Message          string `json:"message"`
}

type SendFileByUrlRequest struct {
	IdInstance       string `json:"id_instance"`
	ApiTokenInstance string `json:"api_token_instance"`
	ChatID           string `json:"chat_id"`
	FileUrl          string `json:"file_url"`
}
