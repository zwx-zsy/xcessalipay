package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAliPay_TradeRefund(t *testing.T) {
	float := strconv.FormatFloat(9.9, 'f', 2, 64)
	fmt.Println(float)
	//t.Log("========== TradeRefund ==========")
	//var p = alipay.TradeRefund{}
	//p.RefundAmount = "10"
	//p.OutTradeNo = "trade_no_20170623011121"
	//rsp, err := Config.Client.TradeRefund(p)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("%v", rsp.Content)
}
