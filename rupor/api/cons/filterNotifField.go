// filterNotifField.go - содержит константы описывающие набор названий ключей документа "Уведомление",
// по которым может быть произведена фильтрация
// и перечислимый тип значениями которого могут быть эти константы

package cons

//go:generate enummethods -type FilterNotifField
var (
	filterNotifField = [...]string{"uuid", "regnumber", "category", "status", "activity_status", "createtime", "updated", "event_description"}
)

// DocFilterField представляет название поля по которому производится фильтрация message
type FilterNotifField int

const (
	Uuid = FilterNotifField(iota)
	Regnumber
	Category
	Status
	ActStatus
	NotifCreateTime
	Updated
	EventDescr
)
