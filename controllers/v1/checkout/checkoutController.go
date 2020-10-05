package checkout

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/honeyo7/EcommerceCheckout/models/v1/checkout"
)

func GetFinalAmt(c *gin.Context) {

	res := model.GetCheckoutAmt(c)

	//res = model.GetGamerCount(c)
	c.JSON(http.StatusOK, res)
}
