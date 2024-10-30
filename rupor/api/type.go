package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"strconv"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/api/param"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/api/resp"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/logger"
)

type RuporApiClient struct {
	Client  http.Client
	Header  http.Header
	GetReq  *http.Request
	PostReq *http.Request
	Url     string
	log     logger.Logger
}

// NewGetRequest метод RuporApiClient создает новый объект  Get-запроса
func (c *RuporApiClient) NewGetRequest(ctx context.Context, uri string, header http.Header, params *param.GetParametrs) error {

	var (
		err error
		req *http.Request
	//	b   []byte
	)

	url := fmt.Sprintf("http://%s/%s/", c.Url, uri)

	if req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil); err != nil {
		return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	//заголовок запроса
	if header != nil {
		req.Header = header
	} else if c.Header != nil {
		req.Header = c.Header
	}

	//формируем параметры запроса
	if params != nil {
		if req.URL.RawQuery, err = getReqParamsToStr(params); err != nil {
			return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		}
	}

	c.GetReq = req
	return nil
}

// NewPostRequest - метод RuporApiClient создает новый объект  Post-запроса
func (c *RuporApiClient) NewPostRequest(ctx context.Context, uri string, header http.Header, body io.Reader) error {
	var (
		err error
		req *http.Request
	//	b   []byte
	)

	url := fmt.Sprintf("http://%s/%s/", c.Url, uri)

	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, body); err != nil {
		return fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	//заголовок запроса
	if header != nil {
		req.Header = header
	} else if c.Header != nil {
		req.Header = c.Header
	}

	c.PostReq = req
	return nil
}

// DoGet выполняет Get-запрос к Api Rupor-а
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

// DoGet выполняет Post-запрос к Api Rupor-а
func (c *RuporApiClient) DoPost(handler func(*http.Response) (*resp.Response, error)) (*resp.Response, error) {
	if c.PostReq.URL.Path == "" {
		err := fmt.Errorf("%s : Несформирован объект Request, Post-запрос не может быть выполнен", debugging.GetFuncName())
		return nil, err
	}

	response, err := c.Client.Do(c.PostReq)
	//log.Printf("%s : send next GET-request", time.Now().Format(time.UnixDate))
	if err != nil {
		err := fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		return nil, err
	}
	return handler(response)
}

//	GetRequest - осуществляет GET-запрос к указанному URI и порционный возврат результата (batch-ирование ответа)
//
//	Параметры:
//
// endPoint - URI api rupor
// getParams - параметры фильтрации документов
// out - канал в который отправляются результаты запроса
func (c *RuporApiClient) GetRequest(ctx context.Context, endPoint string, getParams param.GetParametrs, out chan<- []json.RawMessage, endSignalCh chan<- struct{}) {

	var (
		response *resp.Response
		err      error
	)

	for {
		c.NewGetRequest(ctx, endPoint, nil, &getParams)
		c.log.Debugf("new GET-request URI: %s, docs offset: %d", endPoint, getParams.Offset)
		if response, err = c.DoGet(resp.Handler); err != nil || response.Data.IsEmpty() {

			//	getParams.Offset = 0
			if err != nil {
				c.log.Error(err.Error())
			}
			break
		}

		out <- response.Data.Result
		c.log.Debugf("new GET-response URI: %s, docs offset: %d, docs count: %d was send to chanel", endPoint, getParams.Offset, len(response.Data.Result))
		if !response.Data.Next {
			c.log.Debugf("complite GET-request URI: %s, docs offset: %d", endPoint, getParams.Offset)
			//	docParams.Offset = 0
			break
		}
		getParams.Offset = getParams.Offset + len(response.Data.Result)
	}
	endSignalCh <- struct{}{}
}

//	PostRequest - осуществляет POST-запрос к указанному URI
//
//	Параметры:
//
// endPoint - URI api rupor
// body - тело запроса
func (c *RuporApiClient) PostRquest(ctx context.Context, endPoint string, body io.Reader) []json.RawMessage {
	var (
		response *resp.Response
		err      error
	)
	c.NewPostRequest(ctx, endPoint, nil, body)
	c.log.Debugf("new POST-request URI: %s", endPoint)
	if response, err = c.DoPost(resp.Handler); err != nil || response.Data.IsEmpty() {
		if err != nil {
			c.log.Error(err.Error())
		}
	}
	return response.Data.Result
}

// getReqParamsToStr - конвертирует параметры Get-запроса в строку
func getReqParamsToStr(params *param.GetParametrs) (string, error) {
	var (
		b   []byte
		err error
	)
	query := url.Values{}
	if len(params.Filter) != 0 {
		if b, err = json.Marshal(params.Filter); err != nil {
			return "", err
		}
		query.Add("filter", fmt.Sprintf(string(b)))
	}
	if len(params.Projection) != 0 {
		if b, err = json.Marshal(params.Projection); err != nil {
			return "", err
		}
		query.Add("fields", fmt.Sprintf(string(b)))
	}
	if len(params.Sort) != 0 {
		if b, err = json.Marshal(params.Sort); err != nil {
			return "", err
		}
		query.Add("sort", fmt.Sprintf(string(b)))
	}
	if params.Limit != 0 {
		query.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset != 0 {
		query.Add("offset", strconv.Itoa(params.Offset))
	}

	return query.Encode(), err
}

//Вместо этого участка кода теперь getReqParamsToStr
/*if params != nil {
	query := req.URL.Query()
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
}*/
