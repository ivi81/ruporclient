package export

import (
	"time"
)

//CommonInf объект содержащий общие для всех типо документов информационные поля
type CommonObj struct {
	Uuid             string
	Category         CustomString `json:"category"`
	Company          CustomString `json:"company"`
	CreateTime       time.Time    `json:"createtime"`       //Дата и время регистрации. Дата и время регистрации сообщения от НКЦКИ. Заполняется по правилам стандарта ISO 8601. Используется timeZone UTC
	DetectTime       time.Time    `json:"detecttime"`       //Дата и время выявления. Дата и время выявления инцидента.
	EndTime          time.Time    `json:"endtime"`          //Дата и время завершения. Дата и время завершения инцидента.
	Updated          time.Time    `json:"updated"`          //Дата и время последнего обновления карточки документа
	Ownername        string       `json:"ownername"`        //Владелец информационного ресурса
	EventDescription string       `json:"eventdescription"` //Краткое описание события ИБ
	RegNumber        string       `json:"regnumber"`        //Регистрационный номер
	Status           CustomString `json:"status"`           //Статус документа (у каждого типа документов свой статус)
	Tlp              CustomString `json:"tlp"`              //Ограничительный маркер TLP
	Type             CustomString `json :"type"`            //Тип события ИБ
}

//Message Cообщение НКЦКИ
type Message struct {
	CommonObj
}

type Notification struct {
	CommonObj
	Assistance               bool         `json:"assistance"`         //"Необходимость привлечения сил ГосСОПКА"
	Activitystatus           CustomString `json:"activitystatus"`     //"Статус реагирования на инцидент.
	DetectionTool            string       `json:"detectiontool"`      // "Сведения о средстве или способе выявления инцидента"
	AffectedSystemName       string       `json:"affectedsystemname"` //"Наименование контролируемого ресурса, на котором был выявлен компьютерный инцидент"
	AffectedSystemCategory   CustomString `json:"affectedsystemcategory"`
	AffectedSystemFunction   CustomString `json:"affectedsystemfunction"`    //"Сфера функционирования субъекта. Значение передается строго из справочника"
	AffectedSystemConnection bool         `json:"affectedsystemconnection:"` //Наличие подключения к сети Интернет. Если в данном ключе передается значение true, то API ЛК ожидает в запросе также блок полей с описанием технических сведений об атакованном объекте и атакующей системе"
	IntegrityImpact          CustomString `json:"integrityimpact"`
	AvailabilityImpact       CustomString `json:"availabilityimpact"`    //"Влияние на доступность
	ConfidentialityImpact    CustomString `json:"confidentialityimpact"` //"Влияние на конфиденциальность
	CustomImpact             string       `json:"customimpact"`          //"Краткое описание иной формы последствий компьютерного инцидента"
	Location                 []Location   `json:"location"`              //"Страна/регион. Значение из справочника ISO-3166-2"
	City                     string       `json:"city"`                  //"Населенный пункт или геокоординаты"

}

//Comment коментарий к карточке Уведоления или Сообщения
type Comment struct {
	Id         int       `json: "id"`
	IncUuid    string    `json: "incident.uuid"`
	Text       string    `json: "text"`
	Login      string    `josn: "login"`
	IncidentId string    `json: "incident.id"`
	CreateTime time.Time `json: "create_time"`
}
