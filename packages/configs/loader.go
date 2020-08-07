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
	if conf != nil {
		return conf
	}
	conf = &Config{}
	load("server", &conf.Server)
	load("file", &conf.File)
	load("connect", &conf.Connect)

	return conf
}

func getPath(name string) string {
	return fmt.Sprintf("./conf/%s.ini", name)
}

func load(section string, v interface{}) {
	if err := cfg.Section(section).MapTo(v); err != nil {
		log.Fatalf("%s configs map to struct err: %v\n", section, err)
	}
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

func (conf *Config) GetDefaultLanguage() string {
	return conf.Server.Lang
}

func (conf *Config) GetAcceptLang() []string {
	if len(conf.AcceptLang) > 0 {
		return conf.AcceptLang
	}

	return []string{conf.GetDefaultLanguage()}
}

func (conf *Config) GetNeedLoadLangFile() (files []string) {
	for _, lang := range conf.GetAcceptLang() {
		if lang != "" && lang != "und" {
			files = append(files, fmt.Sprintf("./conf/lang/%s.json", lang))
		}
	}
	return
}
