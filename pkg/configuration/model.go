package configuration

type Configuration struct {
	App      string
	Version  string
	Http     HttpConfig
	Database struct {
		Mysql MysqlConfig
	}
	Logger
}

type HttpConfig struct {
	Port string
}

type MysqlConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

type Logger struct {
	SaveLogFile        bool
	SysLogFileLocation string
	TDRLogFileLocation string
}
