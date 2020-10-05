package checkout

import (
	prop "github.com/honeyo7/EcommerceCheckout/class/clsCommon"
)

type PdtPurchase struct {
	StrSKU string `json:"StrSKU"`
	IntQty int64  `json:"Quantity"`
}

type checkOutRes struct {
	NetAmt   float64 `json:"NetAmount"`
	TotalAmt float64 `json:"TotalAmount"`
	DiscAmt  float64 `json:"DiscAmt"`
}

type ReqCheckout struct {
	AppData prop.Req_data `json:"AppData"`
	ReqData []PdtPurchase `json:"Data"`
}

type ResCheckout struct {
	StatusData prop.Res_data `json:"StatusData"`
	ResData    checkOutRes   `json:"Data"`
}
