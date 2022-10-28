package conf

import "github.com/AH-dark/bing-wallpaper/pkg/util"

const defaultConfig = `[System]
Listen = :8080
Debug = false
SessionSecret = {{SessionSecret}}

[Database]
Type = sqlite
DBFile = bing-wallpaper.db

[Storage]
Type = local
BasePath = data
`

const (
	BackendVersion = "1.1.0"
)

var SystemConfig = &system{
	Listen:        util.EnvStr("LISTEN", ":8080"),
	Debug:         util.EnvStr("DEBUG", "false") == "true",
	SessionSecret: util.EnvStr("SESSION_SECRET", ""),
}

var DatabaseConfig = &database{
	Type:        util.EnvStr("DB_TYPE", "sqlite3"),
	Host:        util.EnvStr("DB_HOST", "localhost"),
	Port:        util.EnvInt("DB_PORT", 3306),
	Username:    util.EnvStr("DB_USERNAME", "root"),
	Password:    util.EnvStr("DB_PASSWORD", ""),
	Database:    util.EnvStr("DB_DATABASE", "bing_wallpaper"),
	Charset:     util.EnvStr("DB_CHARSET", "utf8"),
	DBFile:      util.EnvStr("DB_FILE", "bing-wallpaper.db"),
	TablePrefix: util.EnvStr("DB_TABLE_PREFIX", ""),
	SSLMode:     util.EnvStr("DB_SSL_MODE", "disable"),
}

var RedisConfig = &redis{
	Network:  util.EnvStr("REDIS_NETWORK", "tcp"),
	Server:   util.EnvStr("REDIS_SERVER", ""),
	Password: util.EnvStr("REDIS_PASSWORD", ""),
	DB:       util.EnvInt("REDIS_DB", 0),
}

var CORSConfig = &cors{
	AllowOrigins:     util.EnvArr("CORS_ALLOW_ORIGINS", []string{"*"}),
	AllowMethods:     util.EnvArr("CORS_ALLOW_METHODS", []string{"GET", "HEAD", "OPTIONS"}),
	AllowCredentials: util.EnvStr("CORS_ALLOW_CREDENTIALS", "true") == "true",
	MaxAge:           util.EnvInt("CORS_MAX_AGE", 600),
	AllowHeaders:     util.EnvArr("CORS_ALLOW_HEADERS", []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "accept", "origin", "Cache-Control", "X-Requested-With"}),
	ExposeHeaders:    util.EnvArr("CORS_EXPOSE_HEADERS", nil),
}

var StorageConfig = &storage{
	Type:      util.EnvStr("STORAGE_TYPE", "local"),
	Endpoint:  util.EnvStr("STORAGE_ENDPOINT", ""),
	Region:    util.EnvStr("STORAGE_REGION", "us-east-1"),
	AccessID:  util.EnvStr("STORAGE_ACCESS_ID", ""),
	AccessKey: util.EnvStr("STORAGE_ACCESS_KEY", ""),
	Bucket:    util.EnvStr("STORAGE_BUCKET", ""),
	BasePath:  util.EnvStr("STORAGE_BASE_PATH", "data"),
	BaseUrl:   util.EnvStr("STORAGE_BASE_URL", ""),
	SSL:       util.EnvStr("STORAGE_SSL", "false") == "true",
	ACL:       util.EnvStr("STORAGE_ACL", "public-read"),
}
