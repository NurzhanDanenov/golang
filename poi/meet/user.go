package meet

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
}
