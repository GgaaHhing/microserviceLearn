package internal

type AccountSrvConfig struct {
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type AccountWebConfig struct {
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type AppConfig struct {
	DBConfig         DBConfig         `mapstructure:"db" json:"db"`
	RedisConfig      RedisConfig      `mapstructure:"redis" json:"redis"`
	ConsulConfig     ConsulConfig     `mapstructure:"consul" json:"consul"`
	AccountWebConfig AccountWebConfig `mapstructure:"account_web" json:"account_web"`
	AccountSrvConfig AccountSrvConfig `mapstructure:"account_srv" json:"account_srv"`
	JWTConfig        JWTConfig        `mapstructure:"jwt" json:"jwt"`
}
