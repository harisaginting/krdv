package user

type ResponseList struct {
	Items []User `json:"items"`
	Total int    `json:"total"`
}

type User struct {
	ID       int    `json:"id" `
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}
