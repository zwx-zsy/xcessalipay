package API

import (
	"XcessAlipay/Config"
	"XcessAlipay/proto"
	"context"
	"github.com/smartwalle/alipay"
)

func (s *Myserver) TradePagePay(c context.Context, in *service.AliPayRequest) (result *service.AliPageResponse,err error) {
	result = &service.AliPageResponse{}
	var p = alipay.AliPayTradePagePay{}
	p.NotifyURL = in.Data["NotifyURL"]
	p.ReturnURL = in.Data["ReturnURL"]
	p.Body = in.Data["Body"]
	p.Subject = in.Data["Subject"]
	p.OutTradeNo = in.Data["OutTradeNo"]
	p.TotalAmount = in.Data["TotalAmount"]
	p.ProductCode = in.Data["ProductCode"]
	urls, err := Config.Client.TradePagePay(p)
	if err != nil {
		return nil,err
	}
	result.Payurl = urls.String()
	return result,nil
}