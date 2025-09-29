package api

import (
	"crypto/tls"

	"net/http"

	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/config"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/logger"
)

// NewApiClient создает и настраивает клиент API Rupor-a
func NewApiClient(cfg *config.ClientRuporCfg, logg logger.Logger) (RuporApiClient, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := RuporApiClient{
		Url:    cfg.Host,
		Client: http.Client{Transport: tr},
		Header: *createHeader(cfg),
		log:    logg,
	}
	return c, nil
}

// CreateHeader - создает заголовок HTTP-запроса
func createHeader(cfg *config.ClientRuporCfg) *http.Header {

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
