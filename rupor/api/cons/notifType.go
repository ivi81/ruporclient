package cons

type NotifCiType string

func (m NotifCiType) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("notifCiTypes", data); err == nil {
		m = NotifCiType(s)
		return nil
	}
	return err
}

type NotifCaType string

func (m NotifCaType) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("notifCaTypes", data); err == nil {
		m = NotifCaType(s)
		return nil
	}
	return err
}

type NotifCvType string

func (m NotifCvType) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("notifCvTypes", data); err == nil {
		m = NotifCvType(s)
		return nil
	}
	return err
}
