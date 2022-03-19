package config

type mysql struct {
	Ip       string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}
