package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
)

type carService struct {
	carRepository repository.CarRepository
}

type CarService interface {
	// functional
	CreateCar(ctx context.Context, carDTO dto.CreateCarDTO) (entity.Car, error)
	GetAllCar(ctx context.Context) ([]entity.Car, error)
	GetCarById(ctx context.Context, id uint64) (entity.Car, error)
}

func NewCarService(cr repository.CarRepository) CarService {
	return &carService{
		carRepository: cr,
	}
}

func (cs *carService) CreateCar(ctx context.Context, carDTO dto.CreateCarDTO) (entity.Car, error) {
	newCar := entity.Car{
		Name: carDTO.Name,
		Tipe: carDTO.Tipe,
	}

	return cs.carRepository.InsertCar(ctx, newCar)
}

func (cs *carService) GetAllCar(ctx context.Context) ([]entity.Car, error) {
	return cs.carRepository.GetAllCar(ctx)
}

func (cs *carService) GetCarById(ctx context.Context, id uint64) (entity.Car, error) {
	return cs.carRepository.GetCarById(ctx, id)
}
