package config

type ServerConfig struct {
	Server        Server      `yaml:"server"`
	Elasticsearch ESConfig    `yaml:"elasticsearch"`
	Mongo         MongoConfig `yaml:"mongo"`
	Redis         RedisConfig `yaml:"redis"`
	Mysql         MysqlConfig `yaml:"mysql"`
}
type ESConfig struct {
	Url      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
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
type MongoConfig struct {
	Host         string `yaml:"host"`
	Port         int32  `yaml:"port"`
	DataBaseName string `yaml:"databaseName"`
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type RedisConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   int    `yaml:"dbName"`
}
