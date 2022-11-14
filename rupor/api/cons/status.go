package cons

var (
	msgStatus   = [...]string{"Новое", "В работе", "Принято решение", "Взаимодействие завершено", "Отправлено в архив"}
	notifStatus = [...]string{"Создано", "Зарегистрировано", "Требуется дополнение", "Проверка НКЦКИ", "Принято решение", "Отправлено в архив"}
)

const (
	NotifCREATED = NotifStatus(iota)
	NotifREGISTERED
	NotifREQADD
	NotifVERIFICATION
	NotifDESISION
	NotifARCHIVED
)

//MsgStatus статус Уведомления
type NotifStatus int

func (m NotifStatus) IsValid() bool {

	switch m {
	case
		NotifCREATED,
		NotifREGISTERED,
		NotifREQADD,
		NotifVERIFICATION,
		NotifDESISION,
		NotifARCHIVED:

		return true
	}
	return false
}

func (m NotifStatus) String() string {
	if m.IsValid() {
		return notifStatus[m]
	}
	return ""
}

const (
	MsgNEW = MsgStatus(iota)
	MsgINWORK
	MsgDSISION
	MsgCOMPLITED
	MsgARCHIVED
)

//MsgStatus статус Сообщения
type MsgStatus int

func (m MsgStatus) IsValid() bool {

	switch m {
	case
		MsgNEW,
		MsgINWORK,
		MsgDSISION,
		MsgCOMPLITED,
		MsgARCHIVED:

		return true
	}
	return false
}

func (m MsgStatus) String() string {
	if m.IsValid() {
		return msgStatus[m]
	}
	return ""
}

/*type MsgStatus string

func (m MsgStatus) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("msgStatus", data); err == nil {
		m = MsgStatus(s)
		return nil
	}
	return err
}*/
