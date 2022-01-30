package models

type Secret struct {
	ID         string `json:".id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Service    string `json:"service"`
	Profile    string `json:"profile"`
	LastLogout string `json:"last-logged-out"`
	Comment    string `json:"comment"`
}
