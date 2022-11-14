package cons

var (
	activityStatus = [...]string{"Меры приняты", "Проводятся мероприятия по реагированию", "Возобновлены мероприятия по реагированию", "Инцидент не подтвержден"}
)

const (
	NotifTAKENMEAS = ActivitySt(iota)
	NotifCARRIEDOUTMEAS
	NotifRESUMEDMEAS
	NotifNOTCONFIRMEDINCIDENT
)

type ActivitySt int

func (m ActivitySt) IsValid() bool {

	switch m {
	case
		NotifTAKENMEAS,
		NotifCARRIEDOUTMEAS,
		NotifRESUMEDMEAS,
		NotifNOTCONFIRMEDINCIDENT:

		return true
	}
	return false
}

func (m ActivitySt) String() string {
	if m.IsValid() {
		return activityStatus[m]
	}
	return ""
}
