// config.go описание конфигурации клиента RUPOR
package config

import "time"

// ClientRuporCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
type ClientRuporCfg struct {
	Host          string        `yaml:"url" env:"HOST"`
	Port          int           `yaml:"port" env:"PORT"`
	ApiKey        string        `yaml:"apiKey" env:"API_KEY"`
	UserAgent     string        `yaml:"userAgent" env:"USER_AGENT"`
	Timeout       time.Duration `yaml:"timeout" env:"TIMEOUT"`
	SleepTime     time.Duration `yaml:"sleepTime" env:"SLEEP_TIME"`
	ResponseLimit int           `yaml:"responseLimit" env:"RESPONCE_LIMIT"`
}
