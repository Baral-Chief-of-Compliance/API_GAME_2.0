package model

type UserEnter struct {
	Login    string `json:"login"`
	PassHash string `json:"pass"`
}
