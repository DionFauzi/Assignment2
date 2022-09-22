package routers

import (
	"ASSIGNMENT2/controllers"

	"github.com/gin-gonic/gin"
)

func RouteBrand(route *gin.Engine) {
	brand := route.Group("/brand")

	brand.POST("", controllers.CreateBrand)
	brand.PUT("/:brandID", controllers.UpdateBrand)
	brand.GET("/:brandID", controllers.GetBrand)
}
