package cons

type Impact string

func (m Impact) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("impact", data); err == nil {
		m = Impact(s)
		return nil
	}
	return err
}
