package API

import (
	"XcessAlipay/Config"
	"XcessAlipay/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/smartwalle/alipay"
)

func (s *Myserver) TradeRefund(ctx context.Context, in *service.AliRefundRequest) (resp *service.AliRefoundResponse, err error) {
	resp = &service.AliRefoundResponse{}
	var p = alipay.AliPayTradeRefund{}
	p.RefundAmount = in.RefundAmount
	p.OutTradeNo = in.OutTradeNo
	rsp, err := Config.Client.TradeRefund(p)
	if err != nil {
		return resp, err
	} else {
		resp.Sign = rsp.Sign
		resp.OutTradeNo = rsp.AliPayTradeRefund.OutTradeNo
		resp.BuyerUserId = rsp.AliPayTradeRefund.BuyerUserId
		resp.TradeNo = rsp.AliPayTradeRefund.TradeNo
		resp.Code = rsp.AliPayTradeRefund.Code
		return resp, err
	}

}

func (s *Myserver) FundTransToAccountTransfer(ctx context.Context, in *service.AliPayRequest) (resp *service.AliPayFundTransToAccountTransferResponse, err error) {
	resp = &service.AliPayFundTransToAccountTransferResponse{}
	var p = alipay.AliPayFundTransToAccountTransfer{}
	p.OutBizNo = in.Data["OutBizNo"]//
	p.PayeeType = "ALIPAY_LOGONID"
	p.PayeeAccount = in.Data["PayeeAccount"] //收账支付宝账号
	p.Amount = in.Data["Amount"]             //转账金额
	rsp, err := Config.Client.FundTransToAccountTransfer(p)
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}
	jsonBytes, err := json.Marshal(rsp.Body)
	fmt.Println("jsonBytes", string(jsonBytes))
	json.Unmarshal([]byte(jsonBytes), resp)
	return resp, err
}
