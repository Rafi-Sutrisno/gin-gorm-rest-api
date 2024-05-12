package entity

type Car struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" binding:"required"`
	Tipe      string     `json:"tipe" binding:"required"`
	ListImage []CarImage `json:"list_image,omitempty"`
}

type Bike struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Tipe string `json:"tipe" binding:"required"`
	Foto string `json:"foto" binding:"required"`
}

type CarImage struct {
	Path  string `json:"path"`
	CarID uint64 `gorm:"foreignKey" json:"car_id"`
	Car   *Car   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"car,omitempty"`
}
