package RPC

import (
	"XcessAlipay/Config"
	"XcessAlipay/proto"
	"context"
	"github.com/smartwalle/alipay"
	"net/http"
	"net/url"
)

type Myserver struct {
}

type Result struct {
	AppId       string                 `json:"app_id"`
	Timestamp   string                 `json:"timestamp"`
	Biz_content map[string]interface{} `json:"biz_content"`
	Sign        string                 `json:"sign"`
}

func (s *Myserver) TradeAppPay(ctx context.Context, in *service.AliPayRequest) (resp *service.AliPayResponse, err error) {
	resp = &service.AliPayResponse{}
	var p = alipay.AliPayTradeAppPay{}
	p.NotifyURL = in.Data["NotifyURL"]
	p.Body = in.Data["Body"]
	p.Subject = in.Data["Subject"]
	p.OutTradeNo = in.Data["OutTradeNo"]
	p.TotalAmount = in.Data["TotalAmount"]
	p.ProductCode = in.Data["ProductCode"]
	param, err := Config.Client.TradeAppPay(p)
	res := param.(url.Values)
	resp.Codes = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]string{}
	resp.Data["app_id"] = res.Get("app_id")
	resp.Data["method"] = res.Get("method")
	resp.Data["timestamp"] = res.Get("timestamp")
	resp.Data["biz_content"] = res.Get("biz_content")
	resp.Data["sign"] = res.Get("sign")

	rs := &Result{}
	rs.AppId = res.Get("app_id")
	rs.Timestamp = res.Get("timestamp")

	return resp, nil
}
