package config

type Redis struct {
	Name     string `json:"name" yaml:"name"`
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}
