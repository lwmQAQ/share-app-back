package config

type ServerConfig struct {
	Server        Server      `yaml:"server"`
	Elasticsearch ESConfig    `yaml:"elasticsearch"`
	Mongo         MongoConfig `yaml:"mongo"`
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

type MongoConfig struct {
	Host         string `yaml:"host"`
	Port         int32  `yaml:"port"`
	DataBaseName string `yaml:"databaseName"`
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
}
