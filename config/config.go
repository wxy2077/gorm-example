package config

type Config struct {
	MainMySQL *MySQL
	Runtime   *Runtime
}

type MySQL struct {
	DNS         string
	Database    string
	MaxIDleConn int
	MaxOpenConn int
	MaxLifeTime int
}

type Runtime struct {
	HttpPort       int64
	JaegerHostPort string
	ServerName     string
}
