//config.go расширяет функционал go-microconfig в части настроек клиента к rupor
package config

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"

//RuporClientCfg - структура данных описывающая информацию о конфигурации модуля логирования микросервисов
type RuporClientCfg struct {
	microconfig.ClientHttpCfg
}

//SetValuesFromEnv загружает значение перменных среды которые имеют префикс заданный в envPrefix
//в структуру ServiceLoggerCfg
func (cfg *RuporClientCfg) SetValuesFromEnv(envPrefix string) {
	cfg.ClientHttpCfg.SetValuesFromEnv(envPrefix)
}
