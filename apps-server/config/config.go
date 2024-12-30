package config

type ServerConfig struct {
	Server Server      `yaml:"server"`
	Mysql  MysqlConfig `yaml:"mysql"`
	Redis  RedisConfig `yaml:"redis"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbName"`
}

type RedisConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   int    `yaml:"dbName"`
}
