package controller

import (
	"mods/dto"
	"mods/service"
	"mods/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type carController struct {
	carService service.CarService
}

type CarController interface {
	// regist login
	InsertCar(ctx *gin.Context)
	GetAllCar(ctx *gin.Context)
}

func NewCarController(cs service.CarService) CarController {
	return &carController{
		carService: cs,
	}
}

func (cc *carController) InsertCar(ctx *gin.Context) {

	var car dto.CreateCarDTO
	if tx := ctx.ShouldBind(&car); tx != nil {
		res := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := cc.carService.CreateCar(ctx.Request.Context(), car)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to register user", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("Success to register car", http.StatusOK, result)
	ctx.JSON(http.StatusOK, res)
}

func (cc *carController) GetAllCar(ctx *gin.Context) {
	carList, err := cc.carService.GetAllCar(ctx)
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success ini mobil mu", http.StatusOK, carList)
	ctx.JSON(http.StatusOK, res)

}
