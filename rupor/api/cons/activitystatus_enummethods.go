
//activitystatus_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package cons

import (
	
	"github.com/ivi81/enummethods/enumerator"
)


//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m ActivityStatus) IsValid() bool {

	switch m {
	case
		ActivityTAKENMEAS,
		ActivityCARRIEDOUTMEAS,
		ActivityRESUMEDMEAS,
		ActivityNOTCONFIRMEDINCIDENT:
		return true
	}
	return false
}


//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m ActivityStatus) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *ActivityStatus) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}


//MarshalYAML - реализует метод интерфейса yaml.Marshaler
func (m ActivityStatus) MarshalYAML() ([]byte, error) {
	return enumerator.MarshalYAML(m)
}

//UnmarshalYAML - реализует метод интерфейса json.UnMarshaler
func (m *ActivityStatus) UnmarshalYAML(data []byte) error {
	return enumerator.UnmarshalYAML(m, data)
}


//String конвертация значения типа в строку. 
//Реализует интерфейс Stringer
func (m ActivityStatus) String() string {
	if m.IsValid() {
		return activityStatus[m]
	}
	return ""
}


//SetValue - конвертация строки в значение перечислимого типа.
//Реализует интерфейс Unstringer
func (m *ActivityStatus) SetValue(s string) bool {
	for i, v := range activityStatus {
		if v == s {
			*m = ActivityStatus(i)
			return true
		}
	}
	return false
}
