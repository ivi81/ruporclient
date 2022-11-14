package cons

//MsgRes результат рассмотрения сообщения
type MsgType string

func (m MsgType) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("msgType", data); err == nil {
		m = MsgType(s)
		return nil
	}
	return err
}
