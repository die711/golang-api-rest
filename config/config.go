package config

import (
	"fmt"
	"github.com/eduardogpg/gonv"
)

type Config interface {
	url() string
}

type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
}

type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	database.username = gonv.GetStringEnv("dbUSERNAME", "root")
	database.password = gonv.GetStringEnv("dbPASSWORD", "11597")
	database.host = gonv.GetStringEnv("HOST", "localhost")
	database.port = gonv.GetIntEnv("HOST", 3306)
	database.database = gonv.GetStringEnv("DATABASE", "project_go_web")
	database.debug = gonv.GetBoolEnv("DEBUG", true)

	server = &ServerConfig{}
	server.host = gonv.GetStringEnv("HOST", "localhost")
	server.port = gonv.GetIntEnv("PORT", 3000)
	server.debug = gonv.GetBoolEnv("DEBUG", true)

}

func (d *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", d.username, d.password, d.host, d.port, d.database)
}

func Debug() bool {
	return server.debug
}

//<username>:<password>@tcp(<host>:<port>)/<database>
func UrlDatabase() string {
	return database.url()
}

func UrlServer() string {
	return server.url()
}

func (s *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func ServerPort() int {
	return server.port
}
