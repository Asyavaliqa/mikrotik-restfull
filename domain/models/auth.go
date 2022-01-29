package models

type Auth struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
