package cons

type ProductCategory string

func (m ProductCategory) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("productCategory", data); err == nil {
		m = ProductCategory(s)
		return nil
	}
	return err
}
