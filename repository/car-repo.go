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
	GetCarById(ctx context.Context, id uint64) (entity.Car, error)
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

func (db *carConnection) GetCarById(ctx context.Context, id uint64) (entity.Car, error) {
	var Car entity.Car

	tx := db.connection.Find(&Car, id)

	if tx.Error != nil {
		return entity.Car{}, tx.Error
	}

	return Car, nil
}