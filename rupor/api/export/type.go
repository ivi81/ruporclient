package export

import (
	"time"
)

// CommonInf объект содержащий общие для всех типо документов информационные поля
type CommonObj struct {
	Uuid             string
	Category         CustomString `json:"category"`
	Company          CustomString `json:"company"`
	CreateTime       time.Time    `json:"create_time"`       //Дата и время регистрации. Дата и время регистрации сообщения от НКЦКИ. Заполняется по правилам стандарта ISO 8601. Используется timeZone UTC
	DetectTime       time.Time    `json:"detect_time"`       //Дата и время выявления. Дата и время выявления инцидента.
	EndTime          time.Time    `json:"end_time"`          //Дата и время завершения. Дата и время завершения инцидента.
	Updated          time.Time    `json:"updated"`           //Дата и время последнего обновления карточки документа
	Ownername        string       `json:"ownername"`         //Владелец информационного ресурса
	EventDescription string       `json:"event_description"` //Краткое описание события ИБ
	RegNumber        string       `json:"reg_number"`        //Регистрационный номер
	Identifier       string       `json:"identifier"`        //идентификатор в web-интерфейсе
	Status           CustomString `json:"status"`            //Статус документа (у каждого типа документов свой статус)
	Tlp              CustomString `json:"tlp"`               //Ограничительный маркер TLP
	Type             CustomString `json :"type"`             //Тип события ИБ
}

// Message Cообщение НКЦКИ
type Message struct {
	CommonObj
}

type Notification struct {
	CommonObj
	Assistance               bool         `json:"assistance"`           //"Необходимость привлечения сил ГосСОПКА"
	Activitystatus           CustomString `json:"activity_status"`      //"Статус реагирования на инцидент.
	DetectionTool            string       `json:"detectiontool"`        // "Сведения о средстве или способе выявления инцидента"
	AffectedSystemName       string       `json:"affected_system_name"` //"Наименование контролируемого ресурса, на котором был выявлен компьютерный инцидент"
	AffectedSystemCategory   CustomString `json:"affected_system_category"`
	AffectedSystemFunction   CustomString `json:"affected_system_function"`    //"Сфера функционирования субъекта. Значение передается строго из справочника"
	AffectedSystemConnection bool         `json:"affected_system_connection:"` //Наличие подключения к сети Интернет. Если в данном ключе передается значение true, то API ЛК ожидает в запросе также блок полей с описанием технических сведений об атакованном объекте и атакующей системе"
	IntegrityImpact          CustomString `json:"integrity_impact"`
	AvailabilityImpact       CustomString `json:"availability_impact"`    //"Влияние на доступность
	ConfidentialityImpact    CustomString `json:"confidentiality_impact"` //"Влияние на конфиденциальность
	CustomImpact             string       `json:"custom_impact"`          //"Краткое описание иной формы последствий компьютерного инцидента"
	Location                 []Location   `json:"location"`               //"Страна/регион. Значение из справочника ISO-3166-2"
	City                     string       `json:"city"`                   //"Населенный пункт или геокоординаты"

}

// RespComment коментарий к карточке Уведоления или Сообщения
type RespComment struct {
	Id         int       `json:"id"`
	IncUuid    string    `json:"incident.uuid"`
	Text       string    `json:"text"`
	Login      string    `josn:"login"`
	IncidentId string    `json:"incident.id"`
	CreateTime time.Time `json:"create_time"`
}

// Comment добавляемый новый коментарий к карточке Уведоления или Сообщения
type Comment struct {
	IncUuid string      `json:"incident.uuid"`
	Data    CommentText `json:"data"`
}

type CommentText struct {
	Comment string `json:"comment"`
}
