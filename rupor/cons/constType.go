package cons

import "fmt"

//ConstArr используется для хранения списка строковых констант
type ConstArr []string

func (cArr *ConstArr) Exists(val string) error {
	for _, n := range *cArr {
		if val == n {
			return nil
		}
	}
	return fmt.Errorf("value %s not in range %v", val, cArr)
}

//ConstMap используется для хранения набора списков строковых констант
type ConstMap map[string]ConstArr

func (cMap *ConstMap) GetSlice(key string) (ConstArr, error) {

	if slice, ok := (*cMap)[key]; ok {
		return slice, nil
	}

	var keys []string
	for k := range *cMap {
		keys = append(keys, k)
	}
	return nil, fmt.Errorf("error: category '%s' not in range оf Constants list %s", key, keys)
}

func (cMap *ConstMap) Exists(key string, val string) bool {

	var (
		slice ConstArr
		ok    bool
	)

	if slice, ok = (*cMap)[key]; !ok {
		return ok
	}

	if err := slice.Exists(val); err != nil {
		return false
	}
	return true
}
