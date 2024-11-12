package conf

type JWTConfig struct {
	SigningKey string `mapstructure:"signing_key"`
}
