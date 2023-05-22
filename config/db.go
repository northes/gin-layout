package config

import "fmt"

type DBDriverType string

var DBDriver = struct {
	MySQL   DBDriverType
	SQLite  DBDriverType
	SQLite3 DBDriverType
}{
	MySQL:   "mysql",
	SQLite:  "sqlite",
	SQLite3: "sqlite3",
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

func (m *MySQL) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.DBName,
	)
}

func (s *SQLite) GetDSN() string {
	return fmt.Sprintf("file:%s?cache=shared&_fk=1", s.Name)
}
