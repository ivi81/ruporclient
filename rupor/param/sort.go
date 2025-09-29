package param

//ElemtSort элемент списка сортировки
type ElemSort struct {
	Property  string `json:"property"`
	Direction string `json:"direction"`
}

//Sort список сортировки
type Sort []ElemSort
