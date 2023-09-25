package model

type UserEnter struct {
	Name     string `json:"name"`
	PassHash string `json:"passHash"`
}
