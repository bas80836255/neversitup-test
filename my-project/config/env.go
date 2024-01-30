package config

var (
	Env = environmentConfig{}
)

// Env data structure
type environmentConfig struct {
	AppEnv     string `env:"ENV" envDefault:"local"`
	AppVersion string `env:"APP_VERSION" json:"version"`
}
