package config

type ServerConfig struct {
	Server      Server      `yaml:"server"`
	EmailConfig EmailConfig `yaml:"email"`
}

type EmailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"userName"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}
type JWTConfig struct {
	Key string `yaml:"key"`
}
