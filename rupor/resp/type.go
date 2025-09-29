package resp

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/export"
)

// NkckiGetResp ответ на GET-запрос к Rupor
type Response struct {
	Success    bool     `json:"success"`
	Data       RespData `json:"data,omitempty"`
	Msg        string   `json "message,omitempty"`
	Err        string   `json:"error,omitempty"`
	StatusCode int      `json:"statusCode,omitempty"`
}

// RespData поле данных ответа (содержит результаты запроса)
type RespData struct {
	Result []json.RawMessage //RespResult ////`json:"result"`
	Next   bool              ////`json:"next"`
}

func (r *RespData) UnmarshalJSON(data []byte) error {

	strData := string(data)

	if next := gjson.Get(strData, "next"); next.Exists() {
		r.Next = next.Bool()
	} else {
		r.Next = false
	}

	if result := gjson.Get(strData, "result"); result.Exists() {
		data = []byte(result.Raw)
	}

	if isJsonObject(data) {
		r.Result = append(r.Result, data)
	} else if err := json.Unmarshal(data, &r.Result); err != nil {
		return err
	}

	return nil
}

func (r *RespData) IsEmpty() bool {
	if len(r.Result) == 0 {
		return true
	}
	return false
}

// RespResult список документов возвращенных в результате запроса
type RespResult []interface{}

func (r *RespResult) UnmarshalJSON(data []byte) error {

	var (
		rawSlice []json.RawMessage
		//	respRsult RespResult
		err error
	)

	if err := json.Unmarshal(data, &rawSlice); err != nil {
		return err
	}

	for _, raw := range rawSlice {

		var f interface{}
		if f, err = decodeResult(raw); err != nil {
			return err
		}

		(*r) = append((*r), f)

	}

	return nil
}

// decodeResult распознает тип объекта и делает его Unmarshal
func decodeResult(b []byte) (interface{}, error) {

	category := gjson.Get(string(b), "category.name").String()
	var doc interface{}
	switch category {
	case cons.DocMSG.String():

		doc = export.Message{}
		if err := json.Unmarshal(b, &doc); err != nil {
			fmt.Println(string(b))
			return nil, err
		}
		return doc, nil
	case cons.DocNOTIFCA.String(),
		cons.DocNOTIFCI.String(),
		cons.DocNOTIFCV.String():

		doc = export.Notification{}
		if err := json.Unmarshal(b, &doc); err != nil {
			fmt.Println(string(b))
			return nil, err
		}
		return doc, nil
	}

	return nil, nil
}

// isNotJsonObject проверяет является ли срез байт json-объектом
func isJsonObject(b []byte) bool {

	dec := json.NewDecoder(bytes.NewBuffer(b))

	t, err := dec.Token()
	if err != nil {
		return false
	}
	if d, ok := t.(json.Delim); ok && d == '{' {
		return true
	}
	return false
}
