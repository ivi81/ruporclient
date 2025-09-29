package export

import (
	"fmt"

	"github.com/tidwall/gjson"
	cons "gitlab.cloud.gcm/i.ippolitov/go-ruporclient/rupor/cons"
)

const timeLayout = "2022-06-20T16:36:28.000Z"

// Company Краткое наименование организации являющейся субъектом ГосСОПКА "АО \"Ромашка\""
type Company CustomString

type Location struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Linked     bool   `json:"linked"`
	Coordinate string `json:"coordinate"`
}

//Uuid уникальный идентификатор Документа "4f00b46c-4543-47b2-93f3-3bf029563623"
//type Uuid string

//Domain Информация о доменном имени
/*type Domain struct {
	Value     string `json:"value"`
	RegDomain string `json:"regdomain"`
}*/

//type IpV4 string

//type IpV6 string

//Lir Наименование LIR (Локальная Регистратура Интернета) "BBQ"
//type Lir string

//Asn Номер ASN (Номер автономной система. AS - cодержит в себе адресное пространство (IP-адреса),
// имеет уникальный ASN — номер, позволяющий однозначно идентифицировать AS в Интернете) "AS-5513"
//type Asn string

//AsName Наименование AS (Автономной системы)
//type AsName string

//Country Название страны
//type Country string

// Malware Информаци об вредоносной системе
type Malware struct {
	Function cons.MalwareFunc `json:"function"`
	Sinkhole bool             `json:"sinkhole"`
}

//Hash Хеш-сумма
//type Hash string

//Email Email-адрес вредоносного объекта
//type Email string

//Uri URI-адрес вредоносной системы "https://badboy.com//efweW42ERFSA"
//type Uri string

//VulnInfo Описание используемых уязвимостей "Пароль был admin"
//type VulnInfo string

//Service Сетевая служба и порт/протокол "443/https"
/*type Service struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}*/

//Path "AS-Path до атакованной Автономной системы (ASN)"
//type Path string

// ProdInfo  "Наименование и версия уязвимого продукта"
type ProdInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

//VulnId Идентификатор уязвимости "CVE-2020-1231"
//type VulnId string

// CustomString нужен для выбора поля name из вложенного JSON, остальные поля игнорируются
type CustomString string

func (s *CustomString) UnmarshalJSON(data []byte) error {

	str := string(data)
	switch str {
	case "null", "nil":
		return nil
	}

	name := gjson.Get(str, "name")
	if name.Exists() {

		*s = CustomString(name.String())
		return nil
	}

	return fmt.Errorf("error: field 'name' not in json field")
}

type NestedObj struct {
	Name string `json:"name"`
	Uuid string `json:"uuid,omitempty"`
	Id   int    `json:"id,omitempty"`
}
