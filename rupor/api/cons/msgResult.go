package cons

//MsgResult результат рассмотрения сообщения
type MsgResult string

func (m MsgResult) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("msgResults", data); err == nil {
		m = MsgResult(s)
		return nil
	}
	return err
}
