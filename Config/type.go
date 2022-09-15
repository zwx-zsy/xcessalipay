package Config

//start
type Yaml struct {
	PGDBConf      PostgresqlDB `yaml:"PostgresqlDB"`
	Server        Server       `yaml:"Server"`
	RedisConf     Redis        `yaml:"Redis"`
	JwtConf       JwtConf      `yaml:"JwtConf"`
	WeChat        WeChats      `yaml:"Wechat"`
	AliAccessKey  AliAccessKey `yaml:"AliAccessKey"`
	AliPayConf    AliPayConf   `yaml:"AliPayConf"`
	WXPayConf     WXPayConf    `yaml:"WXPayConf"`
	AliOssConf    AliOssConf   `yaml:"AliOss"`
	TimeLayoutStr string       `yaml:"TimeLayoutStr"`
	JdUnion       JdUnion      `yaml:"JdUnion"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type PostgresqlDB struct {
	User         string `yaml:"user"`
	Password     string `yaml:"passWord"`
	Address      string `yaml:"address"`
	DatabaseName string `yaml:"databaseName"`
	Port         int64  `yaml:"port"`
}

type Redis struct {
	Host                       string `yaml:"host"`
	PassWord                   string `yaml:"passWord"`
	Port                       string `yaml:"port"`
	DatabaseName               int    `yaml:"db"`
	VerificationRedisKeyPrefix string `json:"verificationRedisKeyPrefix"`
}

type JwtConf struct {
	IsSuer    string `yaml:"issuer"`
	ExpTime   int64  `yaml:"expTime"`
	NotBefore int64  `yaml:"notBefore"`
}

type JdUnion struct {
	AppSecret string `yaml:"appSecret"`
	AppKey    string `yaml:"appKey"`
}

//微信资源配置
type WeChats struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

type WXPayConf struct {
	AppID     string `yaml:"appID"`
	MchID     string `yaml:"mchID"`
	WXPApiKey string `yaml:"wxPApiKey"`
}

//阿里资源配置
type AliAccessKey struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	AliYunDomain    string `yaml:"aliYunDomain"`
	SmsVersion      string `yaml:"smsVersion"`
	SmsTemplateCode string `yaml:"smsTemplateCode"`
	SmsSignName     string `yaml:"smsSignName"`
	SmsRegionId     string `yaml:"smsRegionId"`
}

type AliPayConf struct {
	AppID        string `yaml:"appID"`
	AliPublicKey string `yaml:"aliPublicKey"`
	PrivateKey   string `yaml:"privateKey"`
	IsProduction bool   `yaml:"isProduction"`
}

type AliOssConf struct {
	AccessKeyID     string `yaml:"ossAccessKeyID"`
	AccessKeySecret string `yaml:"ossAccessKeySecret"`
	Bucket          string `yaml:"ossBucket"`
	EndPoint        string `yaml:"ossEndPoint"`
	CallBackHost    string `yaml:"ossCallBackHost"`
}
