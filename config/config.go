package config

import (
	"fmt"
	"github.com/eduardogpg/gonv"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Debug    bool
}

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("dbUSERNAME", "root")
	database.Password = gonv.GetStringEnv("dbPASSWORD", "11597")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("HOST", 3306)
	database.Database = gonv.GetStringEnv("DATABASE", "project_go_web")
	database.Debug = gonv.GetBoolEnv("DEBUG", true)
}

func (d *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", d.Username, d.Password, d.Host, d.Port, d.Database)
}

func GetDebug() bool {
	return database.Debug
}

//<username>:<password>@tcp(<host>:<port>)/<database>
func GetUrlDatabase() string {
	return database.url()
}
