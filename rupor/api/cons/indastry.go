package cons

type Indastry string

func (m Indastry) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("indastry", data); err == nil {
		m = Indastry(s)
		return nil
	}
	return err
}
