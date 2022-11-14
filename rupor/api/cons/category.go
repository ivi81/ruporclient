package cons

var (
	categorys = [...]string{"Сообщение от НКЦКИ",
		"Уведомление о компьютерном инциденте",
		"Уведомление о компьютерной атаке",
		"Уведомление о наличии уязвимости"}
)

const (
	DocMSG = DocCategory(iota)
	DocNOTIFCI
	DocNOTIFCA
	DocNOTIFCV
)

//MsgCategory категории сообщений
type DocCategory int

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

func (m DocCategory) String() string {
	if m.IsValid() {
		return categorys[m]
	}
	return ""
}

/*type MsgCategory string

func (m MsgCategory) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)
	if s, err = umarshalNameField("msgCategorys", data); err == nil {
		m = MsgCategory(s)
		return nil
	}
	return err
}

//NotifCategory категории уведомлений
type NotifCategory string

func (m NotifCategory) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)
	if s, err = umarshalNameField("notifCategorys", data); err == nil {
		m = NotifCategory(s)
		return nil
	}
	return err
}*/
