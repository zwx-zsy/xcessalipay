package Config

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	ServerConf                 *Yaml
	BaseYamlPath               string
	IrisPath                   string
	CertPath                   string
	VerificationRedisKeyPrefix string
	LocationZone               *time.Location
	Client                     *alipay.Client
)

func init() {
	LocationZone, _ = time.LoadLocation("Asia/Shanghai")

	if os.Getenv("Env") == "Develop" {
		BaseYamlPath = "./DockerBuild/configuration/confdev.yaml"
		CertPath = "./DockerBuild/alipaycert/"
	} else {
		BaseYamlPath = "/etc/xcessalipay/conf.yaml"
		IrisPath = "/etc/xcessalipay/iris.yaml"
		CertPath = "/etc/alipaycert/"
	}
	yamlFile, err := ioutil.ReadFile(BaseYamlPath)
	if err != nil {
		log.Fatal(err)
	} else {
		err = yaml.Unmarshal(yamlFile, &ServerConf)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("ServerConf.AliPayConf.AppID", ServerConf.AliPayConf.AppID)
	fmt.Println("ServerConf.AliPayConf.PrivateKey", ServerConf.AliPayConf.PrivateKey)
	Client, err = alipay.New(ServerConf.AliPayConf.AppID, ServerConf.AliPayConf.PrivateKey, ServerConf.AliPayConf.IsProduction)
	Client.LoadAppPublicCertFromFile(CertPath + "appCertPublicKey_2019050664411025.crt") // 加载应用公钥证书
	Client.LoadAliPayRootCertFromFile(CertPath + "alipayRootCert.crt")                   // 加载支付宝根证书
	Client.LoadAliPayPublicCertFromFile(CertPath + "alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书
}
