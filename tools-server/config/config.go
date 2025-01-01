package config

type ServerConfig struct {
	Server Server      `yaml:"server"`
	Minio  MinioConfig `yaml:"minio"`
	Mysql  MysqlConfig `yaml:"mysql"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbName"`
}
