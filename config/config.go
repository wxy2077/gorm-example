package config

type Config struct {
	Runtime   *Runtime
	System    *System
	MainMySQL *MySQL
}

type System struct {
	PrefixUrl string
}

type Runtime struct {
	HttpPort       int64
	Mode           string
	JaegerHostPort string
	ServerName     string
}

type MySQL struct {
	DNS         string
	Database    string
	MaxIDleConn int
	MaxOpenConn int
	MaxLifeTime int
}
