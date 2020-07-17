package forms

// LoginForm ...
type LoginForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	Pass string `form:"pass" json:"pass" binding:"required"`
}
