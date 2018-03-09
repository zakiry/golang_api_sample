package entity

type DbConf struct {
	dialect  string
	user     string
	password string
	address  string
	port     int
	dbname   string
}

type DbConfBuilder struct {
	Dialect  string
	User     string `env:"MysqlUser"`
	Password string `env:"MysqlPassword"`
	Address  string `env:"MysqlAddress"`
	Port     int    `env:"MysqlPort"`
	Dbname   string `env:"MysqlDbname"`
}

func (builder DbConfBuilder) Build() DbConf {
	return DbConf{
		dialect:  builder.Dialect,
		user:     builder.User,
		password: builder.Password,
		address:  builder.Address,
		port:     builder.Port,
		dbname:   builder.Dbname,
	}
}

func (conf DbConf) Dialect() string {
	return conf.dialect
}

func (conf DbConf) User() string {
	return conf.user
}

func (conf DbConf) Password() string {
	return conf.password
}

func (conf DbConf) Address() string {
	return conf.address
}

func (conf DbConf) Port() int {
	return conf.port
}

func (conf DbConf) Dbname() string {
	return conf.dbname
}
