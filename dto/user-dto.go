package dto

type CreateCarDTO struct {
	Name string `json:"name" binding:"required"`
	Tipe string `json:"tipe" binding:"required"`
}
