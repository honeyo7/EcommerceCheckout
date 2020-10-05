package clsCommon

type Req_data struct {
	StrKey      string `json:"StrKey"`
	AppVer      string `json:"AppVer"`
	Imei        string `json:"Imei"`
	DeviceModel string `json:"DeviceModel"`
	OSVer       string `json:"OSVer"`
	IPAddress   string `json:"IPAddress"`
}

type Res_data struct {
	State     string `json:"State"`
	Status    string `json:"Status"`
	ErrorCode string `json:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg"`
}
