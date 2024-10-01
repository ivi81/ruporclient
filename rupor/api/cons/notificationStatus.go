// notificationStatus.go - содержит константы описывающие набор допустимых значений поля "status" в Уведомлении
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate enummethods -type NotifStatus
var (
	notifStatus = [...]string{"Создано", "Зарегистрировано", "Требуется дополнение", "Проверка НКЦКИ", "Принято решение", "Отправлено в архив"}
)

// MsgStatus статус Уведомления
type NotifStatus int

const (
	NotifCREATED = NotifStatus(iota)
	NotifREGISTERED
	NotifREQADD
	NotifVERIFICATION
	NotifDESISION
	NotifARCHIVED
)
