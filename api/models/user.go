package models

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
}

type Login struct {
	Name     string `form:"Name" json:"Name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
