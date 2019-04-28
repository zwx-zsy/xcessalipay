package Config

type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

//start
type Yaml struct {
	ConfigKey string  `yaml:"ConfigKey"`
	DBConf    MongoDB `yaml:"mongodb"`
	RedisConn string  `yaml:"redisConn"`
	Server    Server  `yaml:"Server"`
}


type MongoDB struct {
	User         string `yaml:"db_user"`
	Host         string `yaml:"db_host"`
	Password     string `yaml:"db_pass"`
	Port         string `yaml:"db_port"`
	DatabaseName string `yaml:"db_database_name"`
	AuthDBName   string `yaml:"db_auth_name"`
	Uri          string `yaml:"url"`
}