package api

import (
	"crypto/tls"

	"net/http"

	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/config"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/logger"
)

//NewApiClient создает и настраивает клиент API Rupor-a
func NewApiClient(cfg *config.RuporClientCfg, logg logger.Logger) (RuporApiClient, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := RuporApiClient{
		Url:    cfg.Host,
		Client: http.Client{Transport: tr},
		Header: *createHeader(cfg),
	}
	return c, nil
}

//CreateHeader - создает заголовок HTTP-запроса
func createHeader(cfg *config.RuporClientCfg) *http.Header {

	header := http.Header{}
	if cfg.ApiKey != "" {
		header.Add("x-token", cfg.ApiKey)
	}
	if cfg.UserAgent != "" {
		header.Add("User-Agent", cfg.UserAgent)
	}

	header.Add("Accept", "application/json")
	header.Add("Connection", "keep-alive")
	return &header
}
