package model

type Team struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Stars int    `json:"stars"`
}
