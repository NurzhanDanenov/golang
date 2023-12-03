package entity

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
