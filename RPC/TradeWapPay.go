package RPC

import (
	"context"
	"github.com/smartwalle/alipay/v3"
	"github.com/yancyzhou/xcessalipay/Config"
	"github.com/yancyzhou/xcessalipay/proto"
)

//手机页面支付接口
func (s *Myserver) TradeWapPay(ctx context.Context, in *service.AliPayRequest) (resp *service.AliPageResponse, err error) {
	resp = &service.AliPageResponse{}
	var p = alipay.TradeWapPay{}
	p.NotifyURL = in.Data["NotifyURL"]
	p.ReturnURL = in.Data["ReturnURL"]
	p.Subject = in.Data["Subject"]
	p.OutTradeNo = in.Data["OutTradeNo"]
	p.TotalAmount = in.Data["TotalAmount"]
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := Config.Client.TradeWapPay(p)
	resp.Payurl = url.String()
	return resp, nil
}
