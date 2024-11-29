package model

type DBConfig struct {
	UserName string
	PassWord string
	Host     string
	DBName   string
	Charset  string
	TimeZone string
}

func (config *DBConfig) CreateDsn() string {

}

type LogConfig struct {

}

type ServerConfig struct {

}

type AppConfig struct {
	DB     DBConfig
	Log    LogConfig
	Server ServerConfig
}
