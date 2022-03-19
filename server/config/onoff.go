package config

type onoff struct {
	SyncDb bool `mapstructure:"sync-db" json:"syncDb" yaml:"sync-db"`
}
