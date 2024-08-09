package setting

type Config struct {
	Mysql  MysqlSQLSetting `mapstructure:"mysql"`
	Logger LoggerSetting   `mapstructure:"logger"`
}

type MysqlSQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	MaxSize     int    `mapstructure:"maxSize"`
	Compress    bool   `mapstructure:"compress"`
}
