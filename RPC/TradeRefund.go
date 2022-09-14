package RPC

import (
	"XcessAlipay/Config"
	"XcessAlipay/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

func (s *Myserver) TradeRefund(ctx context.Context, in *service.AliRefundRequest) (resp *service.AliRefoundResponse, err error) {
	resp = &service.AliRefoundResponse{}
	var p = alipay.TradeRefund{}
	p.RefundAmount = in.RefundAmount
	p.OutTradeNo = in.OutTradeNo
	rsp, err := Config.Client.TradeRefund(p)
	if err != nil {
		return resp, err
	} else {
		resp.Sign = rsp.Sign
		resp.OutTradeNo = rsp.Content.OutTradeNo
		resp.BuyerUserId = rsp.Content.BuyerUserId
		resp.TradeNo = rsp.Content.TradeNo
		resp.Code = rsp.Content.SubCode
		return resp, err
	}

}

func (s *Myserver) FundTransToAccountTransfer(ctx context.Context, in *service.AliPayRequest) (resp *service.AliPayFundTransToAccountTransferResponse, err error) {
	resp = &service.AliPayFundTransToAccountTransferResponse{}
	var p = alipay.FundTransUniTransfer{}
	PayeeInfoOne := &alipay.PayeeInfo{
		Identity:     in.Data["PayeeAccount"],
		IdentityType: "ALIPAY_LOGON_ID",
		Name:         in.Data["Name"],
	}
	p.BizScene = "DIRECT_TRANSFER"
	p.OutBizNo = in.Data["OutBizNo"]
	p.ProductCode = "TRANS_ACCOUNT_NO_PWD"
	p.PayeeInfo = PayeeInfoOne
	p.TransAmount = in.Data["Amount"]             //转账金额
	rsp, err := Config.Client.FundTransUniTransfer(p)
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}

	jsonBytes, err := json.Marshal(rsp.Content)
	fmt.Println("jsonBytes", string(jsonBytes))
	json.Unmarshal([]byte(jsonBytes), resp)
	return resp, err
}
