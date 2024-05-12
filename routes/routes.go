package routes

import (
	"mods/controller"
	"mods/service"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, carController controller.CarController, jwtService service.JWTService) {
	carPublic := router.Group("/car")
	{
		// public can access
		carPublic.POST("/insert", carController.InsertCar)
		carPublic.GET("", carController.GetAllCar)
		carPublic.GET("/:id", carController.GetCarById)
		carPublic.POST("/image", carController.InsertImage)
		carPublic.POST("/login", carController.CarToken)
		carPublic.GET("/img/:path", carController.GetImage)
		carPublic.POST("/predict/:path", carController.Predict)
	}

}
