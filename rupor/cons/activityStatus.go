// activityStatus.go - содержит константы описывающие набор допустимых значений поля "activity_status" в Уведомлении,
// описывает возможную активностть в отношении уведомления.
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate enummethods -type ActivityStatus
var (
	activityStatus = [...]string{"Меры приняты", "Проводятся мероприятия по реагированию", "Возобновлены мероприятия по реагированию", "Инцидент не подтвержден"}
)

const (
	ActivityTAKENMEAS = ActivityStatus(iota)
	ActivityCARRIEDOUTMEAS
	ActivityRESUMEDMEAS
	ActivityNOTCONFIRMEDINCIDENT
)

type ActivityStatus int
