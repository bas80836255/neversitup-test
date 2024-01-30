package config

var (
	App = configHolder{}
)

// ConfigHolder data structure
type configHolder struct {
	Application configApplication `yaml:"application" json:"application"`
	Server      configServer      `yaml:"server" json:"server"`
}

type configApplication struct {
	Name string `yaml:"name" json:"name"`
}

type configServer struct {
	Port int `yaml:"port" json:"port"`
}
