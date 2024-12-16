package config

type ServerConfig struct {
	Server    Server      `yaml:"server"`
	RpcServer RpcServer   `yaml:"rpcserver"`
	Email     EmailConfig `yaml:"email"`
	JWT       JWTConfig   `yaml:"jwt"`
	Mysql     MysqlConfig `yaml:"mysql"`
	Redis     RedisConfig `yaml:"redis"`
	Etcd      EtcdConfig  `yaml:"etcd"`
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

type RpcServer struct {
	ServerName string `yaml:"servername"`
	Host       string `yaml:"host"`
	Port       int32  `yaml:"port"`
}
type JWTConfig struct {
	Key string `yaml:"key"`
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

type EtcdConfig struct {
	Addrs   []string `yaml:"addrs"`
	Timeout int      `yaml:"timeout"`
}
