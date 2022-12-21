package watchlist

type ResponseList struct{
	Items 	[]User 		`json:"items"`
	Total 	int 		`json:"total"`
}

type User struct{
	ID 				string `json:"id",gorm:"primaryKey"`
	Username 		string `json:"first_name"`
	Fullname 		string `json:"first_name"`
}