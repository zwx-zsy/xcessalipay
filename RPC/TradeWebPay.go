package RPC

import (
	"context"
	"github.com/smartwalle/alipay/v3"
	"github.com/yancyzhou/xcessalipay/Config"
	"github.com/yancyzhou/xcessalipay/proto"
)

func (s *Myserver) TradePagePay(c context.Context, in *service.AliPayRequest) (result *service.AliPageResponse, err error) {
	result = &service.AliPageResponse{}
	var p = alipay.TradePagePay{}
	p.NotifyURL = in.Data["NotifyURL"]
	p.ReturnURL = in.Data["ReturnURL"]
	p.Body = in.Data["Body"]
	p.Subject = in.Data["Subject"]
	p.OutTradeNo = in.Data["OutTradeNo"]
	p.TotalAmount = in.Data["TotalAmount"]
	p.ProductCode = in.Data["ProductCode"]
	urls, err := Config.Client.TradePagePay(p)
	if err != nil {
		return nil, err
	}
	result.Payurl = urls.String()
	return result, nil
}
