package config

type Config struct {
	MainMySQL *MySQL
}

type MySQL struct {
	DNS         string
	Database    string
	MaxIDleConn int
	MaxOpenConn int
	MaxLifeTime int
}
