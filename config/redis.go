package config

type Redis struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       string `mapstructure:"db" json:"db" yaml:"db"`
}
