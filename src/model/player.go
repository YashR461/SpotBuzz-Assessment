package model

type Player struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	Score int `json:"score"`
}