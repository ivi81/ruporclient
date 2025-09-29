package config_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2/env"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/config"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("./test_data/.test.env"); err != nil {
		log.Println(" No .env file found")
	}
	//Создаем перменную окружения хранящую путь к файлам конфигурации
	EnvKeyConfigPath := env.JoinStr(config.ClientRuporPrefix, "CONFIG_PATH")
	os.Setenv(EnvKeyConfigPath, "./test_data/config")

	//Задаем название среды развертывания
	os.Setenv("STAGE", "test")

	os.Exit(m.Run())
}

func TestClientNatsCfg(t *testing.T) {

	t.Run("TEST : load Cfg from env", func(t *testing.T) {
		testCfg := config.ClientRuporCfg{}

		err := microconfig.CfgLoad(&testCfg, config.ClientRuporPrefix, true)

		if !assert.NoError(t, err, "CfgLoad WITHOUT ERROR") {
			assert.ErrorContains(t, err, "no such file or directory")
		}
	})

}
