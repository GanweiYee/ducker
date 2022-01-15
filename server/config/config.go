package config

type Config struct {
	App     app     `json:"app" yaml:"app"`
	Mysql   mysql   `json:"mysql" yaml:"mysql"`
	Redis   redis   `json:"redis" yaml:"redis"`
	Jwt     jwt     `json:"jwt" yaml:"jwt"`
	Switchs switchs `json:"switchs" yaml:"switchs"`
}
