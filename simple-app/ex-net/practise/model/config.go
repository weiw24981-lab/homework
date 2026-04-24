package model

type Config struct {
	Server      ServerConfig   `mapstructure:"server"`
	JWT         JWTConfig      `mapstructure:"jwt"`
	DataTypeCon DataTypeConfig `mapstructure:"dataTypeCon"`
	SqliteCon   SqliteConfig   `mapstructure:"sqliteCon"`
	MysqlCon    MysqlConfig    `mapstructure:"mysqlCon"`
	ProgresCon  ProgresConfig  `mapstructure:"progresCon"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type DataTypeConfig struct {
	DataType string `mapstructure:"dataType"`
}

type SqliteConfig struct {
	Filename string `mapstructure:"filename"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type ProgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}
