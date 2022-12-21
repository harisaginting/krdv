package auth

type ResponseList struct{
	Items 	[]User 		`json:"items"`
	Total 	int 		`json:"total"`
}

type User struct{
	ID 				string `json:"id"`
	Username 		string `json:"username"`
	Fullname 		string `json:"first_name"`
}

type PayloadUserRegister struct{
	ID 				string `json:"id",gorm:"primaryKey"`
	Username 		string `json:"username"`
	Fullname 		string `json:"Fullname"`
	Password 		string `json:"password"`
}

type ResponseUserRegister struct{
	Token 		string `json:"token"`
}

type ResponseMe struct{
	Username 		string `json:"username"`
	Fullname 		string `json:"Fullname"`
}

type PayloadUserLogin struct{
	ID 				string `json:"id",gorm:"primaryKey"`
	Username 		string `json:"username"`
	Fullname 		string `json:"Fullname"`
	Password 		string `json:"password"`
}

type ResponseUserLogin struct{
	Username 		string `json:"username"`
	Fullname 		string `json:"Fullname"`
	Token 			string `json:"token"`
}