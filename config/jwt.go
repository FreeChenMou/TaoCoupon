package config

type JWT struct {
	SigningKey  string `json:"signing-key" yaml:"signing-key"`
	ExpiresTime string `json:"expires-time" yaml:"expires-time"`
	BufferTime  string `json:"buffer-time" yaml:"buffer-time"`
	Issuer      string `json:"issuer" yaml:"issuer"`
}
