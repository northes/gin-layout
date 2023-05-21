package config

type DBDriverType string

var DBDriver = struct {
	MySQL  DBDriverType
	SQLite DBDriverType
}{
	MySQL:  "mysql",
	SQLite: "sqlite",
}

type DB struct {
	Driver  DBDriverType `mapstructure:"driver"`
	*MySQL  `mapstructure:"mysql"`
	*SQLite `mapstructure:"sqlite"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
}

type SQLite struct {
	Name string `mapstructure:"name"`
}
