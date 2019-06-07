package model

//  use tags on struct field declarations to customize the encoded JSON key names
type Account struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
