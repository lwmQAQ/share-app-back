package config

type ServerConfig struct {
	Server Server      `yaml:"server"`
	Minio  MinioConfig `yaml:"minio"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	BucketName      string `yaml:"bucketName"`
}
