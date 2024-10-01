//config.go расширяет функционал go-microconfig в части настроек клиента к rupor
package config

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"

//RuporClientCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
type RuporClientCfg struct {
	Url string `yaml:"url"`
	Key string `yaml:"apiKey"`
}

//SetValuesFromEnv загружает значение перменных среды которые имеют префикс заданный в envPrefix
//в структуру RuporClientCfg
func (cfg *RuporClientCfg) SetValuesFromEnv(envPrefix string) {
	envPref := microconfig.JoinStr(envPrefix, RuporPrefix)
	if url, ok := microconfig.LookupEnv(microconfig.JoinStr(envPref, "URL")); ok {
		cfg.Url = url
	}

	if apiKey, ok := microconfig.LookupEnv(microconfig.JoinStr(envPref, "API-KEY")); ok {
		cfg.Key = apiKey
	}
}
