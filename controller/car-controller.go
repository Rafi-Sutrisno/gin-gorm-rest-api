package controller

import (
	"fmt"
	"mods/dto"
	"mods/service"
	"mods/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type carController struct {
	carService service.CarService
	jwtService service.JWTService
}

type CarController interface {
	// regist login
	InsertCar(ctx *gin.Context)
	GetAllCar(ctx *gin.Context)
	GetCarById(ctx *gin.Context)
	InsertImage(ctx *gin.Context)
	CarToken(ctx *gin.Context)
}

func NewCarController(cs service.CarService, jwt service.JWTService) CarController {
	return &carController{
		carService: cs,
		jwtService: jwt,
	}
}

func (cc *carController) RetrieveID(ctx *gin.Context) (uint64, error) {
	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)

	return cc.jwtService.GetUserIDByToken(token)
}

func (cc *carController) CarToken(ctx *gin.Context) {
	var carLogin dto.LoginDTO
	if tx := ctx.ShouldBind(&carLogin); tx != nil {
		res := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	car, err := cc.carService.GetCarByName(ctx, carLogin.Name)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to login, car no registered", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := cc.jwtService.GenerateToken(car.ID, car.Name)
	fmt.Print(token)
	res := utils.BuildResponse("Successful login", http.StatusOK, token)
	ctx.JSON(http.StatusOK, res)
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

func (cc *carController) GetCarById(ctx *gin.Context) {
	Carid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	CarByid, err := cc.carService.GetCarById(ctx, Carid)
	if err != nil {
		res := utils.BuildErrorResponse("failed to get car id info", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to get car info", http.StatusOK, CarByid)
	ctx.JSON(http.StatusOK, res)

}

func (cc *carController) InsertImage(ctx *gin.Context) {

	var newImage dto.CarImageDTO
	err := ctx.ShouldBind(&newImage)
	if err != nil {
		ctx.String(http.StatusBadRequest, "get form error %s", err.Error())
	}

	carId, err := cc.RetrieveID(ctx)
	if err != nil {
		response := utils.BuildErrorResponse("invalid token", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := cc.carService.CarImage(ctx, newImage, carId)
	if err != nil {
		res := utils.BuildErrorResponse("failed to insert", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to upload image", http.StatusOK, result)
	ctx.JSON(http.StatusOK, res)
}
