package config

type config struct {
	App   app
	DB    database
	Redis redis
	Log   log
}

type app struct {
	Port string
	Mode string
}

type database struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
}

type log struct {
	Path  string
	Level string
}

type redis struct {
	Addr string
}
