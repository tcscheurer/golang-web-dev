package main

import (
	"encoding/json"
	"fmt"
)

type code struct {
	Code        string `json:"code"`
	Description string `json:"desc"`
}

type codes []code

func main() {
	var c codes
	rcvd := `[{"code":"200","desc":"StatusOK"},{"code":"301","desc":"StatusMovedPermanently"},{"code":"302","desc":"StatusFound"},{"code":"303","desc":"StatusSeeOther"},{"code":"307","desc":"StatusTemporaryRedirect"},{"code":"400","desc":"StatusBadRequest"},{"code":"401","desc":"StatusUnauthorized"},{"code":"402","desc":"StatusPaymentRequired"},{"code":"403","desc":"StatusForbidden"},{"code":"404","desc":"StatusNotFound"},{"code":"405","desc":"StatusMethodNotAllowed"},{"code":"418","desc":"StatusTeapot"},{"code":"500","desc":"StatusInternalServerError"}]`
	if e := json.Unmarshal([]byte(rcvd), &c); e != nil {
		fmt.Println(e)
	}
	fmt.Println(c)
}
