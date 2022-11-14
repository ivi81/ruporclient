package cons

//Tlp значения протокола TLP
type Tlp string

func (m Tlp) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("tlp", data); err == nil {
		m = Tlp(s)
		return nil
	}
	return err
}