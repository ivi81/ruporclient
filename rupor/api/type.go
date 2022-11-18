package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/api/param"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/api/resp"
	"gitlab.cloud.gcm/i.ippolitov/go-servicelogger/logger"
)

//
type RuporApiClient struct {
	Client  http.Client
	Header  http.Header
	GetReq  *http.Request
	PostReq http.Request
	Url     string
}

//NewGetRequest метод RuporApiClient создает новый объект  Get-запроса
func (c *RuporApiClient) NewGetRequest(ctx context.Context, uri string, header http.Header, params *param.GetParametrs) error {

	var (
		err error
		req *http.Request
		b   []byte
	)

	url := fmt.Sprintf("http://%s/%s/", c.Url, uri)

	if req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil); err != nil {
		return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	query := req.URL.Query()

	if header != nil {
		req.Header = header
	} else if c.Header != nil {
		req.Header = c.Header
	}

	if params != nil {
		if len(params.Filter) != 0 {
			if b, err = json.Marshal(params.Filter); err != nil {
				return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
			}
			query.Add("filter", fmt.Sprintf(string(b)))
		}
		if len(params.Projection) != 0 {
			if b, err = json.Marshal(params.Projection); err != nil {
				return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
			}
			query.Add("fields", fmt.Sprintf(string(b)))
		}
		if len(params.Sort) != 0 {
			if b, err = json.Marshal(params.Sort); err != nil {
				return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
			}
			query.Add("sort", fmt.Sprintf(string(b)))
		}
		if params.Limit != 0 {
			query.Add("limit", strconv.Itoa(params.Limit))
		}
		if params.Offset != 0 {
			query.Add("offset", strconv.Itoa(params.Offset))
		}
		req.URL.RawQuery = query.Encode()
	}
	c.GetReq = req
	return nil
}

//NewPostRequest метод RuporApiClient создает новый объект  Post-запроса
func (c *RuporApiClient) NewPostRequest() error {
	return nil
}

//DoGet выполняет Get-запрос к Api Rupor-а
func (c *RuporApiClient) DoGet(handler func(*http.Response) (*resp.Response, error)) (*resp.Response, error) {

	if c.GetReq.URL.Path == "" {
		err := fmt.Errorf("%s : Несформирован объект Request, Get-запрос не может быть выполнен", debugging.GetFuncName())
		return nil, err
	}

	response, err := c.Client.Do(c.GetReq)
	//log.Printf("%s : send next GET-request", time.Now().Format(time.UnixDate))
	if err != nil {
		err := fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		return nil, err
	}
	return handler(response)
}

func (c *RuporApiClient) DoPost() error {
	return nil
}

//GetRequest осуществляет GET-запрос к указанному URI и порционный возврат результата
//endPoint - URI api rupor
//getParams - параметры фильтрации докумнтов
//out - канал в который отправляются результаты запроса
func (c *RuporApiClient) GetRequest(ctx context.Context, endPoint string, getParams param.GetParametrs, out chan<- []json.RawMessage, endSignalCh chan<- struct{}, logg *logger.Logger) {

	var (
		response *resp.Response
		err      error
	)

	for {
		c.NewGetRequest(ctx, endPoint, nil, &getParams)
		logg.Debugf("new GET-request URI: %s, docs offset: %s", endPoint, getParams.Offset)
		if response, err = c.DoGet(resp.Handler); err != nil || response.Data.IsEmpty() {

			//	getParams.Offset = 0
			if err != nil {
				logg.Error(err.Error())
			}
			break
		}

		out <- response.Data.Result
		logg.Debugf("new GET-response URI: %s, docs offset: %s, docs count: %s was send to chanel", endPoint, getParams.Offset, len(response.Data.Result))
		if !response.Data.Next {
			logg.Debugf("complite GET-request UDI: %s, docs offset: %s", endPoint, getParams.Offset)
			//	docParams.Offset = 0
			break
		}
		getParams.Offset = getParams.Offset + len(response.Data.Result)
	}
	endSignalCh <- struct{}{}
}
