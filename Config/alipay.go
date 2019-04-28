package Config

import (
	"github.com/smartwalle/alipay"
)

var Client *alipay.AliPay

func init() {
	Client = alipay.New(APPID, ALIPUBLICKEY, PRIVATEKEY, false)
}