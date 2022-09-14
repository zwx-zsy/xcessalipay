package RPC

import (
	"XcessAlipay/Config"
	service "XcessAlipay/proto"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/smartwalle/alipay/v3"
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
	var p = alipay.TradeAppPay{}
	p.NotifyURL = in.Data["NotifyURL"]
	p.Body = in.Data["Body"]
	p.Subject = in.Data["Subject"]
	p.OutTradeNo = in.Data["OutTradeNo"]
	p.TotalAmount = in.Data["TotalAmount"]
	p.ProductCode = in.Data["ProductCode"]
	Param, err := Config.Client.TradeAppPay(p)
	res, err := url.ParseQuery(Param)
	//res := Param.(url.Values)
	resp.Codes = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]string{}
	resp.Data["app_id"] = res.Get("app_id")
	resp.Data["method"] = res.Get("method")
	resp.Data["timestamp"] = res.Get("timestamp")
	resp.Data["biz_content"] = res.Get("biz_content")
	resp.Data["sign"] = res.Get("sign")
	fmt.Println(resp.Data)
	return resp, nil
}
