package dto

import "mime/multipart"

type CreateCarDTO struct {
	Name string `json:"name" binding:"required"`
	Tipe string `json:"tipe" binding:"required"`
}

type CarImageDTO struct {
	File *multipart.FileHeader `form:"file"`
}

type LoginDTO struct {
	Name string `json:"name" binding:"required"`
}
