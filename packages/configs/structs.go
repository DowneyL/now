package configs

import (
	"time"
)

// Config configs
type Config struct {
	Server
	File
	Connect
}

// Server configs
type Server struct {
	Mode            string
	RuntimeRootPath string
	Lang            string
	AcceptLang      []string
	Http            `ini:"server.http"`
}

type Http struct {
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	DefaultPageSize int
	JwtSecret       string
}

// File configs
type File struct {
	UploadPath string
	Image      `ini:"file.image"`
	Log        `ini:"file.log"`
}

type Image struct {
	UploadPath   string
	MaxSize      int
	AllowExtends []string
}

type Log struct {
	Path       string
	NamePrefix string
	FileExtend string
	TimeFormat string
}

// Connect configs
type Connect struct {
	ReadDatabase  Database `ini:"connect.database.read"`
	WriteDatabase Database `ini:"connect.database.write"`
	Redis         `ini:"connect.redis"`
}

type Database struct {
	Type        string
	User        string
	Password    string
	Name        string
	Host        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}
