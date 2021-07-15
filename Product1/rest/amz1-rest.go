package rest

import (
	"Product1/controller"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() {
	r1 := gin.Default()
	grp2 := r1.Group("/v1/inventory")
	{
		grp2.GET("inv", controller.FetchProductsFromInventory)
		grp2.POST("inv", controller.AddProductToInventory)
		grp2.GET("inv/:prodname", controller.GetProductDetailsByName)
		grp2.GET("inv/id/:prodid", controller.GetProductDetailsById)
	}

	r1.Run(":7001")

}
