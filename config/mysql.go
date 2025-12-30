package config

type Mysql struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
}
