package entity

type Person struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}
