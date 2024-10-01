// docCategory.go -содержит константы описывающие набор допустимых значений поля "category" в Сообщениях и Уведомлениях
// и перечислимый тип значениями которого могут быть эти константы
package cons

//go:generate enummethods -type DocCategory
var (
	docCategory = [...]string{"Сообщение от НКЦКИ",
		"Уведомление о компьютерном инциденте",
		"Уведомление о компьютерной атаке",
		"Уведомление о наличии уязвимости"}
)

const (
	DocMSG = DocCategory(iota)
	DocNOTIFCI
	DocNOTIFCA
	DocNOTIFCV
)

// MsgCategory категории сообщений
type DocCategory int
