package Config

import (
	"fmt"
	"github.com/smartwalle/alipay"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	CONFPATH string = "/etc/xcesspay/conf.yaml" // 配置文件地址
	//CONFPATH string = "/septnet/config/conf.yaml" // 配置文件地址
	CONFKEY string = "Config" //配置文件的key
)

var ServerConf *Yaml
var Client *alipay.AliPay

func init() {

	yamlFile, err := ioutil.ReadFile(CONFPATH)
	if err != nil {
		log.Fatal(err)
	} else {
		err = yaml.Unmarshal(yamlFile, &ServerConf)
		if err != nil {
			log.Fatal(err)
		}
	}
	Client = alipay.New(ServerConf.AliPayConf.APPID, ServerConf.AliPayConf.ALIPUBLICKEY, ServerConf.AliPayConf.PRIVATEKEY, ServerConf.AliPayConf.IsPruduction)
	//router.Use(ServerConf.ReadConfig)
}

func (this *MongoDB) String() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", this.User, this.Password, this.Host, this.Port, this.AuthDBName)
}
