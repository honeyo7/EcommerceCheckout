package checkout

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	chkoutvm "github.com/honeyo7/EcommerceCheckout/ViewModels/checkout"
	cmmn "github.com/honeyo7/EcommerceCheckout/class/clsCommon"
	db "github.com/honeyo7/EcommerceCheckout/clsodbc"
)

func GetCheckoutAmt(c *gin.Context) chkoutvm.ResCheckout {

	reqData := chkoutvm.ReqCheckout{}
	res := chkoutvm.ResCheckout{}

	res.StatusData.State = "getAmt"
	res.StatusData.Status = "False"

	res.ResData.NetAmt = 0.00
	res.ResData.TotalAmt = 0.00
	res.ResData.DiscAmt = 0.00

	if err := c.ShouldBindJSON(&reqData); err != nil {
		res.StatusData.ErrorCode = "-1"
		res.StatusData.ErrorMsg = err.Error()
		return res
	}

	i := cmmn.VerifyKey(reqData.AppData.StrKey)

	if i == 0 {
		res.StatusData.ErrorCode = "1"
		res.StatusData.ErrorMsg = "Invalid Key"
		return res
	}

	j := cmmn.VerifyAppVer(reqData.AppData.AppVer)

	if j == 0 {
		res.StatusData.ErrorCode = "2"
		res.StatusData.ErrorMsg = "App Version outdated"
		return res
	}

	var pdtPurchases []chkoutvm.PdtPurchase

	pdtPurchases = reqData.ReqData

	if len(pdtPurchases) == 0 {
		res.StatusData.ErrorCode = "3"
		res.StatusData.ErrorMsg = "No product found!"
		return res
	}

	var dblTotalAmt float64 = 0.00
	var dblDiscAmt float64 = 0.00

	for i := 0; i < len(pdtPurchases); i++ {
		var dblPdtAmt float64 = 0.00
		var dblOfferAmt float64 = 0.00
		var intPdtCount int64 = 0
		var intPdtQty int64 = pdtPurchases[i].IntQty
		intPdtCount, err := db.ExecuteQueryInt("SELECT COUNT(1) FROM ecom_pdt WHERE pdt_sku='" + pdtPurchases[i].StrSKU + "' and quantity>=" + strconv.FormatInt(intPdtQty, 10))
		if err != nil {
			res.StatusData.ErrorCode = "4"
			res.StatusData.ErrorMsg = err.Error()
			return res
		}

		if intPdtCount == 0 {
			res.StatusData.ErrorCode = "5"
			res.StatusData.ErrorMsg = "Product " + pdtPurchases[i].StrSKU + " is out of stock!"
			return res
		}

		selDB, err := db.ExecuteQueryRows("SELECT a.price,a.offer_flag,b.min_qty,b.offer_type,b.disc_per,b.offer_pdt_sku,b.offer_pdt_qty FROM ecom_pdt a LEFT JOIN ecom_offer b on a.pdt_sku=b.pdt_sku WHERE a.pdt_sku='" + pdtPurchases[i].StrSKU + "'")

		if err != nil {
			res.StatusData.ErrorCode = "3"
			res.StatusData.ErrorMsg = err.Error()
			return res
		}

		for selDB.Next() {
			var price, offer_flag, min_qty, offer_type, disc_per, offer_pdt_sku, offer_pdt_qty string

			err = selDB.Scan(&price, &offer_flag, &min_qty, &offer_type, &disc_per, &offer_pdt_sku, &offer_pdt_qty)

			dblPdtAmt, _ = strconv.ParseFloat(price, 64)
			intMinQty, _ := strconv.ParseInt(min_qty, 10, 64)
			dblDiscPer, _ := strconv.ParseFloat(disc_per, 64)

			if offer_flag == "1" && offer_type == "1" {

				if intMinQty <= intPdtQty {

					dblOfferAmt = dblPdtAmt * dblDiscPer / 100 * float64(intPdtQty)
				}
			} else if offer_flag == "1" && offer_type == "2" {
				intoffer_pdt_qty, _ := strconv.ParseInt(offer_pdt_qty, 10, 64)
				if pdtPurchases[i].StrSKU == offer_pdt_sku {

					if intPdtQty >= intMinQty+intoffer_pdt_qty {
						intOfferQty := intPdtQty / (intMinQty + intoffer_pdt_qty)

						dblOfferAmt = dblPdtAmt * float64(intOfferQty) * dblDiscPer / 100
					}
				} else {
					dblOfferAmt = getOfferAmt(intPdtQty, pdtPurchases[i].StrSKU, intMinQty, dblDiscPer, offer_pdt_sku, intoffer_pdt_qty, pdtPurchases)
				}
			}

		}

		dblTotalAmt = dblTotalAmt + (dblPdtAmt * float64(intPdtQty))
		dblDiscAmt = dblDiscAmt + dblOfferAmt

	}

	res.StatusData.Status = "True"

	res.ResData.NetAmt = math.Round((dblTotalAmt-(math.Round(dblDiscAmt*100)/100))*100) / 100
	res.ResData.TotalAmt = dblTotalAmt
	res.ResData.DiscAmt = math.Round(dblDiscAmt*100) / 100

	return res

}

func getOfferAmt(intPdtQty int64, pdt_sku string, intMinQty int64, disc_per float64, offer_pdt_sku string, offer_pdt_qty int64, arrPdt []chkoutvm.PdtPurchase) float64 {

	var dblOfferPrice float64 = 0.00
	if intPdtQty < intMinQty {
		return 0.00
	}

	intMaxOffer := intPdtQty / intMinQty

pdtLoop:
	for i := 0; i < len(arrPdt); i++ {
		if arrPdt[i].StrSKU == offer_pdt_sku {
			if offer_pdt_qty > arrPdt[i].IntQty {
				dblOfferPrice = 0.00
			} else {

				var intOfferQty int64 = 0

				if intMaxOffer > arrPdt[i].IntQty {
					intOfferQty = arrPdt[i].IntQty
				} else {
					intOfferQty = intMaxOffer
				}

				dblOfferPdtPrice, _ := db.ExecuteQueryFloat("SELECT price FROM ecom_pdt WHERE pdt_sku='" + arrPdt[i].StrSKU + "'")

				dblOfferPrice = dblOfferPdtPrice * float64(intOfferQty) * disc_per / 100
			}

			break pdtLoop
		}
	}

	return dblOfferPrice
}
