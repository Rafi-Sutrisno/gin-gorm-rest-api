package entity

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