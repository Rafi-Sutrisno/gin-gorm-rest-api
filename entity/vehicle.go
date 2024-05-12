package entity

import "mime/multipart"

type Car struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Tipe string `json:"tipe" binding:"required"`
}

type Bike struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Tipe string `json:"tipe" binding:"required"`
	Foto string `json:"foto" binding:"required"`
}

type CarImage struct {
	File     *multipart.FileHeader `form:"file"`
	Path     string                `form:"path"`
	CarField Car                   `form:"carField"`
}
