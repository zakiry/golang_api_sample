package entity

type DbConf struct {
	Dialect  string
	User     string `env:"MysqlUser"`
	Password string `env:"MysqlPassword"`
	Address  string `env:"MysqlAddress"`
	Port     int    `env:"MysqlPort"`
	Dbname   string `env:"MysqlDbname"`
}
