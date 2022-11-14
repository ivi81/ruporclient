package cons

type SystemCategory string

func (m SystemCategory) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("systemCategory", data); err == nil {
		m = SystemCategory(s)
		return nil
	}
	return err
}
