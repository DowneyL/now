package configs

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	conf *Config
	cfg  *ini.File
	err  error

	server  = &Server{}
	file    = &File{}
	connect = &Connect{}
)

func init() {
	cfg, err = ini.Load(getPath("server"), getPath("file"), getPath("connect"))
	if err != nil {
		log.Fatalf("configs load err: %v\n", err)
	}
	cfg.NameMapper = ini.TitleUnderscore
}

// New configs
func New() *Config {
	if conf == nil {
		return Load()
	}

	return conf
}

// Load configs
func Load() *Config {
	loadServerConfig()
	loadFileConfig()
	loadConnectConfig()
	conf = &Config{
		Server:  *server,
		File:    *file,
		Connect: *connect,
	}

	return conf
}

func getPath(name string) string {
	return fmt.Sprintf("./conf/%s.ini", name)
}

func loadConfig(section string, v interface{}) {
	if err := cfg.Section(section).MapTo(v); err != nil {
		log.Fatalf("%s configs map to struct err: %v\n", section, err)
	}
}

func loadServerConfig() {
	loadConfig("server", server)
	loadConfig("server.http", &server.Http)
}

func loadFileConfig() {
	loadConfig("file", file)
	loadConfig("file.image", &file.Image)
	loadConfig("file.log", &file.Log)
}

func loadConnectConfig() {
	loadConfig("database.read", &connect.ReadDatabase)
	loadConfig("database.write", &connect.WriteDatabase)
	loadConfig("redis", &connect.Redis)
}

func (conf *Config) GetImageUploadPath() string {
	return conf.File.Image.UploadPath
}

func (conf *Config) GetFullImageUploadPath() string {
	return conf.RuntimeRootPath + conf.GetImageUploadPath()
}

func (conf *Config) GetHttpPort() string {
	return fmt.Sprintf(":%d", conf.Server.Http.Port)
}

func (conf *Config) GetHttpReadTimeout() time.Duration {
	return conf.Http.ReadTimeout * time.Second
}

func (conf *Config) GetHttpWriteTimeout() time.Duration {
	return conf.Http.WriteTimeout * time.Second
}
