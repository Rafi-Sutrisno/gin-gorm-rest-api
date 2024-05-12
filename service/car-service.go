package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type carService struct {
	carRepository repository.CarRepository
}

type CarService interface {
	// functional
	CreateCar(ctx context.Context, carDTO dto.CreateCarDTO) (entity.Car, error)
	GetAllCar(ctx context.Context) ([]entity.Car, error)
	GetCarById(ctx context.Context, id uint64) (entity.Car, error)
	CarImage(ctx *gin.Context, imageDTO dto.CarImageDTO, carID uint64) (entity.CarImage, error)
	GetCarByName(ctx context.Context, name string) (entity.Car, error)
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

func (cs *carService) CarImage(ctx *gin.Context, imageDTO dto.CarImageDTO, carID uint64) (entity.CarImage, error) {
	filename := filepath.Base(imageDTO.File.Filename)

	err := ctx.SaveUploadedFile(imageDTO.File, "./image/"+filename)
	if err != nil {
		return entity.CarImage{}, err
	}

	path := "./image/" + filename

	newImage := entity.CarImage{
		Path:  path,
		CarID: carID,
	}

	return cs.carRepository.CarImage(ctx, newImage)

}

func (cs *carService) GetCarByName(ctx context.Context, name string) (entity.Car, error) {
	return cs.carRepository.GetCarByName(ctx, name)
}
