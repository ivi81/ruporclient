
//msgstatus_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package cons

import (
	
	"github.com/ivi81/enummethods/enumerator"
)


//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m MsgStatus) IsValid() bool {

	switch m {
	case
		MsgNEW,
		MsgINPROCESS,
		MsgDESISION,
		MsgCOMPLITED,
		MsgARCHIVED:
		return true
	}
	return false
}


//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m MsgStatus) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *MsgStatus) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}


//MarshalYAML - реализует метод интерфейса yaml.Marshaler
func (m MsgStatus) MarshalYAML() ([]byte, error) {
	return enumerator.MarshalYAML(m)
}

//UnmarshalYAML - реализует метод интерфейса json.UnMarshaler
func (m *MsgStatus) UnmarshalYAML(data []byte) error {
	return enumerator.UnmarshalYAML(m, data)
}


//String конвертация значения типа в строку. 
//Реализует интерфейс Stringer
func (m MsgStatus) String() string {
	if m.IsValid() {
		return msgStatus[m]
	}
	return ""
}


//SetValue - конвертация строки в значение перечислимого типа.
//Реализует интерфейс Unstringer
func (m *MsgStatus) SetValue(s string) bool {
	for i, v := range msgStatus {
		if v == s {
			*m = MsgStatus(i)
			return true
		}
	}
	return false
}