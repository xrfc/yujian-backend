package model

type DBConfig struct {
	UserName string
	PassWord string
	Host     string
	DBName   string
	Charset  string
	TimeZone string
}

// CreateDsn 生成数据库连接字符串。
func (config *DBConfig) CreateDsn() string {
	// 该方法根据DBConfig结构体中的配置信息，构造并返回一个数据库连接字符串。
	// 这个连接字符串可以用于建立与数据库的连接。
	// todo[ruibo]
	return ""
}

type LogConfig struct {
	FileName string // 日志文件名
	LogLevel string // 日志级别
}

type ServerConfig struct {
	Port string
}

type AppConfig struct {
	DB     DBConfig
	Log    LogConfig
	Server ServerConfig
}
