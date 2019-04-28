package Config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	APPID string = "2016093000634363"
	ALIPUBLICKEY string = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsx3wqFgTIGQ+uiNbsnygLG0NFcZUgokYMXFyyFM3fmxjp8WhgeCuI6NMXWf934KkEV2R6+FUEYR0hIDePQKTsOlCi6GanlNJI/9xgUBuJuZWQ7Elcr6e+PzxsoyKEH2ophrmN3hMdUZMSmLX+LWDuNXJ+135EBuLQLEI06K4f/VETAB3xYswm8GC8KrRPn56umKaJYYlyPne1Sz/RjqGeM4W++oMMhUE7zmClHohNpzsnSlZpeQj+JTs+Wx9lxYhySZ4yEAv+WiSVKbVhUVDCU+yH0z71m+thmMnCjs/V3P2/oT/AtU4Htc89kjx83Etx7/PwFvujSdcgQzbXdrV0wIDAQAB"
	PRIVATEKEY string = "MIIEogIBAAKCAQEAy9hTe5u6w9MTCwSq7l+djb+qN4DbkV+b6OZgx5TPQHNLFzjiAs1yYfhK7kIsmhMA0FMlKAp+KvthQAETYW9JpgfIyxpB+GZTWmL+rHUGztzz7xLQuuw2078+W3gB24UoXTG013BlTmWuBL2Su1/ejIxF+8KKWUoy4YJYaR3rUF5KvTsuJxZso8rKHJHeVCXpLbKQVzSKDf8wHFmWXAsPJtElJxxyrx+Uw3O0LKsVewWyIRegWcv9SQe18VvRf3b07EIsXUgn+1d0bZ4fJdMCeuff88IfPNaYpmoVahG/CmiQzEFUhWrEjuQUTTSk9JrnYUuTBUFc9QjcsNvfQ1fpbwIDAQABAoIBAFWRCwoaBr3ovlejr/S74jD8QYFm78QINDQBcKKWjqDpRJ4M3KJ8hwUk6p802/AZBgBsxm4WaXu1Zc8uJGqEApBVOC45aXPR0C9i2rCdtyvih+rVJRUxzpn37z8KEw2IybbFBLpijcPo3rFSumM5DPxGbaIOuXBduxn6Uw5c5sgm7Jggya19AZ5+1mevFo1+gnvP+Av5AKTDCIJX7UzB34ZJB650Jgcrvx77qBRLVuh/dfs7NoIsk5mzTa4ChqAL8pZF3lx7swojVCp3HUDQ3Uma4UATW61S9BpECGb4HdSzX6feu1vsOImpBX2ZvEq9cloVKj0GUikUV/zWk/lS2bECgYEA7lB6vcMfTD+/M/73ANXqFL42iJm2IhfEx0f+jy0Ci8RKXymg63Hf09zyOlWyK3dI2j7hY5LN5uYPmPVHUYQ3soT7Xx9d7jvszpycZkVspTC+XTRiOQU4QZSuL6FzYNlWII9enclDb4zWoRUWwC/Lx8kT1YiG6LF/rr0YXbfEtkUCgYEA2vj+lZeWR/HXZEA1UbKFmw2yJ05DH329qfmq2xF8bCgGj0cEj0RhGjp5JZFVdR4pldgJ9plwt51dB7MaUcOTNacNu7Tlq5Oe2ZOiCyPA462oPxrTT5t7Q0+r53fNk+DFa8v3b3wBBoZYuPgQQKFbVObuDfrVhdHoklmLVEh95iMCgYBxU4oaxTgT7ViCZGOXXSI0fQO3z8jBy6XC0mSmN1Q3nL91I3mnnP5AZJg3z1qruszFJBPQYxZRD+Qo67Lfe810sjJLWMI1MP/6fJUPOUJiB5zu+6Md1HKdFRQECHWyf4eqciMfS38rA8Fo1nt8Jv5z4hXxfDwTWZ0LPlbg8iu0NQKBgDddITE9ZHfIQ8dLHLANTQ6JBfQ+K2Z46k1hnLtKzoSiEhaDxDlc9pglFKBVAKBeAfGq78nzvbYPCYh61jq6EixMkoJVGISrDEzCsQPUBF8Mdy+NXsQP8bhxGmuptMag3AmHRoZfefrmWORbg/BzW545zRKnyvjMu9rU3Q9wTLS1AoGAN3m8sW47Vokt9XR0qhqnG2oExlCAwOc8KSg2dUGLAsaLmLP7kj40eQeeguo7omBBs+jIXJt6rWLYNCEGoG0Y1TmMor1c19LD5d+HRSjsGD+3bF6nLgNyvv0oO8dLe72Tcs1xYrcKfRbARFQpjRTG2BkZVLVBiKjV6NnnDPzFBw4="
)

const (
	//CONFPATH string = "/etc/tl/conf.yaml" // 配置文件地址
	CONFPATH string = "/septnet/config/conf.yaml" // 配置文件地址
	CONFKEY string = "Config" //配置文件的key
)


var ServerConf *Yaml

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
	//router.Use(ServerConf.ReadConfig)
}


func (this *MongoDB) String() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", this.User, this.Password, this.Host, this.Port, this.AuthDBName)
}