package export

type Observ struct {
	Value   string `json:"value"`
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
	Lir     string `json:"lir,omitempty"`
	AsName  string `json:"as_name,omitempty"`
	Asn     string `json:"asn,omitempty"`
	Reg     string `json:"regdomain,omitempty"`
}

//ObservIP - oserv ip indicate ip v 4,6
/*type ObservIP struct {
	Value   string `json:"value"`
	Country string `json:"country,omitempty"`
	Lir     string `json:"lir,omitempty"`
	AsName  string `json:"as_name,omitempty"`
	Asn     string `json:"asn,omitempty"`
}

type ObservHash struct {
	Value string `json:"value"`
}

type ObservEmail struct {
	Value string `json:"value"`
}

type IndVuln struct {
	Value string `json:"value"`
}

type ObservService struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}
type ObservDomain struct {
	Value string `json:"value"`
	Reg   string `json:"regdomain,omitempty"`
}

type ObservMalwareDomain struct {
	ObservDomain
	Function string `json:"function"`
	Sinkhole bool   `json"sinkhole"`
}

type ObservMalwareIP struct {
	ObservIP
	Function string `json:"function"`
	Sinkhole bool   `json"sinkhole"`
}

type ObservUri struct {
	Value string `json:"value"`
}
*/
