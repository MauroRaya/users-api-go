package domain

type User struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
