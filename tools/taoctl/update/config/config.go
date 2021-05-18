package config

import "manlu.org/tao/core/logx"

// Config defines a service configure for taoctl update
type Config struct {
	logx.LogConf
	ListenOn string
	FileDir  string
	FilePath string
}
