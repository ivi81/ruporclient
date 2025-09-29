package param

//ElementFilter элемент фильтра Инцидентов
type ElemFilter struct {
	Property string      `json:"property"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

//Filter фильт инцидентов (параметр GET-запроса)
type Filter []ElemFilter
