package entity

type AppConf struct {
	port int
}

type AppConfBuilder struct {
	Port int
}

func (builder AppConfBuilder) Build() AppConf {
	return AppConf{
		port: builder.Port,
	}
}

func (conf AppConf) Port() int {
	return conf.port
}
