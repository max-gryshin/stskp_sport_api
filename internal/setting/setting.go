package setting

import (
	"os"
)

// App is a structure for storage app configuration
type App struct {
	JwtSecret string

	RuntimeRootPath string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// DbSetting is a structure for storage db configuration
type DbSetting struct {
	Database         string
	Username         string
	Password         string
	PostgresPassword string
	Url              string
}

// ServerSetting is a structure for storage server configuration
type ServerSetting struct {
	RunMode string
	Host    string
	Port    string
	//ReadTimeout  time.Duration
	//WriteTimeout time.Duration
	//Path string
}

type Setting struct {
	ServerConfig ServerSetting
	DbConfig     DbSetting
	App          App
}

var AppSetting = &Setting{}

// LoadSetting loads configuration from env variables
func LoadSetting() *Setting {
	// TODO: Try use go-env for easy unmarshalling https://github.com/Netflix/go-env
	AppSetting = &Setting{
		ServerConfig: ServerSetting{
			Host: getEnv("HOST", ""),
			Port: getEnv("PORT", ""),
		},
		DbConfig: DbSetting{
			Database:         getEnv("POSTGRESQL_DATABASE", ""),
			Username:         getEnv("POSTGRESQL_USERNAME", ""),
			Password:         getEnv("POSTGRESQL_PASSWORD", ""),
			PostgresPassword: getEnv("POSTGRESQL_POSTGRES_PASSWORD", ""),
			Url:              getEnv("DATABASE_URL", ""),
		},
		App: App{
			getEnv("JWT_SECRET", ""),
			getEnv("ROOT_PATH", ""),
			getEnv("SAVE_PATH", ""),
			getEnv("LOG_PATH", ""),
			getEnv("LOG_NAME", ""),
			getEnv("LOG_EXT", ""),
			getEnv("", ""),
		},
	}

	return AppSetting
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
