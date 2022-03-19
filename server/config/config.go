package config

type Config struct {
	App    app    `mapstructure:"app" json:"app" yaml:"app"`
	Mysql  mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt    jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Onoff  onoff  `mapstructure:"on-off" json:"onoff" yaml:"on-off"`
	System system `mapstructure:"system" json:"system" yaml:"system"`
}
