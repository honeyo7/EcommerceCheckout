package routes

import (
	chkout "github.com/honeyo7/EcommerceCheckout/controllers/v1/checkout"

	"github.com/gin-gonic/gin"
)

func RoutesCheckout(route *gin.Engine) {
	getAmt := route.Group("/api/v1/checkout")
	{
		getAmt.POST("/getAmt", chkout.GetFinalAmt)
	}
}
