package config

type redis struct {
	Ip       string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
