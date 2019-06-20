package model

type User struct {
	BaseModel
	Username 	string		`json:"username"`
	Phone 		string		`json:"phone"`
	Email 		string		`json:"email"`
}

//func (User) TableName() string {
//	return "user"
//}
