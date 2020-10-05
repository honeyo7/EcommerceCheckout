// +build unit
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	r "github.com/honeyo7/EcommerceCheckout/routes"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutAmt(t *testing.T) {
	myRouter := gin.Default()
	r.RoutesCheckout(myRouter)
	// Test data
	// 1. Scanned Items: MacBook Pro -1 , Raspberry Pi B-1: Output NetAmount:5399.99,TotalAmount:5429.99,DiscAmt:30
	// 2. Scanned Items: Google Home, Quantity:3 : Output NetAmount:99.98,TotalAmount:149.97,DiscAmt:49.99
	// 3. Scanned Items: Alexa Speaker, Quantity:3 : Output NetAmount:295.65,TotalAmount:328.5,DiscAmt:32.85
	// 4. Scanned Items: MacBook Pro -1 , Raspberry Pi B-2: Output NetAmount:5429.99,TotalAmount:5459.99,DiscAmt:30
	var tests = []struct {
		input    string
		expected string
	}{
		{"{\"AppData\":{\"StrKey\":\"hsulwody638js\",\"AppVer\":\"1.2\",\"Imei\":\"1.2\",\"Device_Model\":\"1.2\",\"OSVer\":\"1.2\",\"IPAddress\":\"12.12.23.34\"},\"Data\":[{\"StrSKU\":\"43N23P\",\"Quantity\":1},{\"StrSKU\":\"234234\",\"Quantity\":1}]}", "{\"StatusData\":{\"State\":\"getAmt\",\"Status\":\"True\",\"ErrorCode\":\"\",\"ErrorMsg\":\"\"},\"Data\":{\"NetAmount\":5399.99,\"TotalAmount\":5429.99,\"DiscAmt\":30}}"},
		{"{\"AppData\":{\"StrKey\":\"hsulwody638js\",\"AppVer\":\"1.2\",\"Imei\":\"1.2\",\"Device_Model\":\"1.2\",\"OSVer\":\"1.2\",\"IPAddress\":\"12.12.23.34\"},\"Data\":[{\"StrSKU\":\"120P90\",\"Quantity\":3}]}", "{\"StatusData\":{\"State\":\"getAmt\",\"Status\":\"True\",\"ErrorCode\":\"\",\"ErrorMsg\":\"\"},\"Data\":{\"NetAmount\":99.98,\"TotalAmount\":149.97,\"DiscAmt\":49.99}}"},
		{"{\"AppData\":{\"StrKey\":\"hsulwody638js\",\"AppVer\":\"1.2\",\"Imei\":\"1.2\",\"Device_Model\":\"1.2\",\"OSVer\":\"1.2\",\"IPAddress\":\"12.12.23.34\"},\"Data\":[{\"StrSKU\":\"A304SD\",\"Quantity\":3}]}", "{\"StatusData\":{\"State\":\"getAmt\",\"Status\":\"True\",\"ErrorCode\":\"\",\"ErrorMsg\":\"\"},\"Data\":{\"NetAmount\":295.65,\"TotalAmount\":328.5,\"DiscAmt\":32.85}}"},
		{"{\"AppData\":{\"StrKey\":\"hsulwody638js\",\"AppVer\":\"1.2\",\"Imei\":\"1.2\",\"Device_Model\":\"1.2\",\"OSVer\":\"1.2\",\"IPAddress\":\"12.12.23.34\"},\"Data\":[{\"StrSKU\":\"43N23P\",\"Quantity\":1},{\"StrSKU\":\"234234\",\"Quantity\":2}]}", "{\"StatusData\":{\"State\":\"getAmt\",\"Status\":\"True\",\"ErrorCode\":\"\",\"ErrorMsg\":\"\"},\"Data\":{\"NetAmount\":5429.99,\"TotalAmount\":5459.99,\"DiscAmt\":30}}"},
	}

	for _, test := range tests {
		payload := strings.NewReader(test.input)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/checkout/getAmt", payload)
		myRouter.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		if assert.Equal(t, test.expected, w.Body.String()) == false {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, w.Body.String())
		}
	}

}
