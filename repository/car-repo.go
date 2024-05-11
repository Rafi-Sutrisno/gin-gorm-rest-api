package repository

import (
	"context"
	"mods/entity"

	"gorm.io/gorm"
)

type carConnection struct {
	connection *gorm.DB
}

type CarRepository interface {
	// functional
	InsertCar(ctx context.Context, car entity.Car) (entity.Car, error)
	GetAllCar(ctx context.Context) ([]entity.Car, error)
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carConnection{
		connection: db,
	}
}

func (db *carConnection) InsertCar(ctx context.Context, car entity.Car) (entity.Car, error) {
	if err := db.connection.Create(&car).Error; err != nil {
		return entity.Car{}, err
	}

	return car, nil
}

func (db *carConnection) GetAllCar(ctx context.Context) ([]entity.Car, error) {
	var listCar []entity.Car

	tx := db.connection.Find(&listCar)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return listCar, nil
}
