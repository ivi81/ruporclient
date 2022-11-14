package cons

import (
	"fmt"

	"github.com/tidwall/gjson"
)

var (
	Consts = ConstMap{}
)

const (
	constFile = "/api/data/ruporApiConst.yaml"
)

func init() {
	/*	path := os.Args
		fmt.Println(path)
		//pwd, _ := os.Getwd()
		pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		yamlFile, err := ioutil.ReadFile(pwd + constFile)
		if err != nil {
			log.Panic("read error yaml file    #%v ", constFile, err)
		}
		err = yaml.Unmarshal(yamlFile, &Consts)
		if err != nil {
			log.Fatalf("error: %v", err)
		}*/
}

func umarshalNameField(key string, data []byte) (string, error) {
	var (
		slice ConstArr
		err   error
	)

	name := gjson.Get(string(data), "name")
	s := name.String()

	if slice, err = Consts.GetSlice(key); err != nil {
		return "", fmt.Errorf("error: for key '%s' %s", key, err.Error())
	}

	if err = slice.Exists(s); err != nil {
		return "", err
	}

	return s, err
}
