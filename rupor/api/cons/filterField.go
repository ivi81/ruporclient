package cons

var (
	filterDocFields     = [...]string{"uuid", "regnumber", "category", "status", "activityStatus", "createtime", "updated"}
	filterCommentFields = [...]string{"incident.uuid", "login", "create_time"}
)

const (
	Uuid = DocFilterField(iota)
	Regnumber
	Category
	Status
	ActivityStatus
	DocCreateTime
	Updated
)

//DocFilterField представляет название поля по которому производится фильтрация message
type DocFilterField int

func (m DocFilterField) IsValid() bool {

	switch m {
	case
		Uuid,
		Regnumber,
		Category,
		Status,
		ActivityStatus,
		DocCreateTime,
		Updated:

		return true
	}
	return false
}

func (m DocFilterField) String() string {
	if m.IsValid() {
		return filterDocFields[m]
	}
	return ""
}

const (
	IncUuid = CommentFilterField(iota)
	Login
	CommentCreateTime
)

type CommentFilterField int

func (m CommentFilterField) IsValid() bool {

	switch m {
	case
		IncUuid,
		Login,
		CommentCreateTime:

		return true
	}
	return false
}

func (m CommentFilterField) String() string {
	if m.IsValid() {
		return filterCommentFields[m]
	}
	return ""
}

//type DocFilterField string

/*func (m DocFilterField) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("msgFilterField", data); err == nil {
		m = DocFilterField(s)
		return nil
	}
	return err
}*/

//NotifFilterField поле по которому производится фильтрация notification
/*type NotifFilterField string

func (m NotifFilterField) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("notifFilterField", data); err == nil {
		m = NotifFilterField(s)
		return nil
	}
	return err
}

//FilterField поле по которому производится фильтрация comments
type CommentFilterField string

func (m CommentFilterField) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("commentFilterField", data); err == nil {
		m = CommentFilterField(s)
		return nil
	}
	return err
}

//BullFilterField поле по которому производится фильтрация bulletin
type BullFilterField string

func (m BullFilterField) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)

	if s, err = umarshalNameField("bullFilterField", data); err == nil {
		m = BullFilterField(s)
		return nil
	}
	return err
}
*/
