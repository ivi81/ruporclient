// filterCommentField.go - содержит константы описывающие набор названий ключей документа "Уведомление",
// по которым может быть произведена фильтрация
// и перечислимый тип значениями которого могут быть эти константы

package cons

//go:generate enummethods -type FilterCommentFields
var (
	filterCommentFields = [...]string{"incident.uuid", "login", "create_time"}
)

type FilterCommentFields int

const (
	IncUuid = FilterCommentFields(iota)
	Login
	CommentCreateTime
)
