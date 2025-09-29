package resp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
)

//Handler обработчик ответа на запрос
func Handler(resp *http.Response) (*Response, error) {

	var (
		err      error
		response = &Response{}
		respBody []byte
	)

	defer resp.Body.Close()

	if respBody, err = io.ReadAll(resp.Body); err != nil {
		err := fmt.Errorf("error Read Body : %s : %s", debugging.GetFuncName(), err.Error())
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusInternalServerError, http.StatusNotFound:

		errStr := string(respBody)
		response, err = initRespError(resp, errStr)
	case http.StatusBadGateway:

		response, err = initRespError(resp, resp.Status)
	case http.StatusOK, http.StatusBadRequest:

		if err = json.Unmarshal(respBody, response); err != nil {
			err := fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
			return nil, err
		}

		if !response.Success {
			_, err = initRespError(resp, response.Err)
		}
	default:

		response, err = initRespError(resp, "Unprocessed Error")
	}
	return response, err
}

func initRespError(httpResp *http.Response, errMsg string) (*Response, error) {
	resp := Response{
		StatusCode: httpResp.StatusCode,
		Err:        errMsg,
		Success:    false,
	}
	err := fmt.Errorf("%s : %s : %s", httpResp.Proto, httpResp.Status, errMsg)

	return &resp, err
}
