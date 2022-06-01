package config

type Config struct {
	DomainConfigs map[string]DomainConfig `mapstructure:"DomainConfig"`
	JWTConfig     JWTConfig
}

type DomainConfig struct {
	ServerConfig ServerConfig `mapstructure:"ServerConfig"`
	DBConfig     DBConfig     `mapstructure:"DBConfig"`
}

type ServerConfig struct {
	AppVersion       string
	Mode             string
	RoutePrefix      string
	Debug            bool
	Port             string
	TimeoutSecs      int64
	ReadTimeoutSecs  int64
	WriteTimeoutSecs int64
}

type JWTConfig struct {
	SessionTime int64
	SecretKey   string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Name     string
	Password string
}
