package param

//GetParametrs параметры GET-запроса к API рупор-а
type GetParametrs struct {
	Filter     Filter
	Sort       Sort
	Projection Fields
	Limit      int
	Offset     int
}
