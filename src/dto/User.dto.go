package dto

type UserDto struct {
	Username string `json:"name" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}
