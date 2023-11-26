package entity

type Image struct {
	Id          int    `json:"id"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description"`
}
