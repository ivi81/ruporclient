// messageSatatus.go -содержит константы описывающие набор допустимых значений поля "status" в Сообщении
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate enummethods -type MsgStatus
var (
	msgStatus = [...]string{"Новое", "В работе", "Принято решение", "Взаимодействие завершено", "Отправлено в архив"}
)

// MsgStatus статус Уведомления
type MsgStatus int

const (
	MsgNEW = MsgStatus(iota)
	MsgINPROCESS
	MsgDESISION
	MsgCOMPLITED
	MsgARCHIVED
)
