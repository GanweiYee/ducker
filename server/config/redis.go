package config

type redis struct {
	Ip       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	DB       int    `json:"db" yaml:"db"`
	Password string `json:"password" yaml:"password"`
}
