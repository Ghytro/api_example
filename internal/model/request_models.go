package model

type PriceRequest struct {
	Id       int    `json:"id"`
	Strength string `json:"strength"`
	Volume   int    `json:"volume"`
	Dopping  string `json:"dopping"`
}

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Name     string   `json:"string"`
	Contacts Contacts `json:"contacts"`
}

type Contacts struct {
	Phone  string `json:"phone"`
	VkLink string `json:"vk_link,omitempty"`
	TgLink string `json:"tg_link,omitempty"`
}
