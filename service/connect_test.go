package main

import (
	"XcessAlipay/Config"
	"github.com/smartwalle/alipay"
	"testing"
)

func TestAliPay_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	var p = alipay.AliPayTradeRefund{}
	p.RefundAmount = "10"
	p.OutTradeNo = "trade_no_20170623011121"
	rsp, err := Config.Client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp.AliPayTradeRefund)
}