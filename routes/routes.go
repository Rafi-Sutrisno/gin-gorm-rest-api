package routes

import (
	"mods/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, carController controller.CarController) {
	carPublic := router.Group("/car")
	{
		// public can access
		carPublic.POST("/insert", carController.InsertCar)
		carPublic.GET("", carController.GetAllCar)
		carPublic.GET("/:id", carController.GetCarById)
	}

}
