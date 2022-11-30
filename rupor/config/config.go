//config.go расширяет функционал go-microconfig в части настроек клиента к rupor
package config

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"

//RuporClientCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
type RuporClientCfg struct {
	microconfig.ClientHttpCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает значение перменных среды которые имеют префикс заданный в envPrefix
//в структуру RuporClientCfg
func (cfg *RuporClientCfg) SetValuesFromEnv(envPrefix string) {
	envPref := microconfig.JoinStr(envPrefix, RuporPrefix)
	cfg.ClientHttpCfg.SetValuesFromEnv(envPref)
}
