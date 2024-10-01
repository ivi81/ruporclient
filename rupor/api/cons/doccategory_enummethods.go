
//doccategory_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package cons

import (
	
	"github.com/ivi81/enummethods/enumerator"
)


//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m DocCategory) IsValid() bool {

	switch m {
	case
		DocMSG,
		DocNOTIFCI,
		DocNOTIFCA,
		DocNOTIFCV:
		return true
	}
	return false
}


//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m DocCategory) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *DocCategory) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}


//MarshalYAML - реализует метод интерфейса yaml.Marshaler
func (m DocCategory) MarshalYAML() ([]byte, error) {
	return enumerator.MarshalYAML(m)
}

//UnmarshalYAML - реализует метод интерфейса json.UnMarshaler
func (m *DocCategory) UnmarshalYAML(data []byte) error {
	return enumerator.UnmarshalYAML(m, data)
}


//String конвертация значения типа в строку. 
//Реализует интерфейс Stringer
func (m DocCategory) String() string {
	if m.IsValid() {
		return docCategory[m]
	}
	return ""
}


//SetValue - конвертация строки в значение перечислимого типа.
//Реализует интерфейс Unstringer
func (m *DocCategory) SetValue(s string) bool {
	for i, v := range docCategory {
		if v == s {
			*m = DocCategory(i)
			return true
		}
	}
	return false
}
