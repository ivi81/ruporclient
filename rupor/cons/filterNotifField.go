// filterNotifField.go - содержит константы описывающие набор названий ключей документа "Уведомление",
// по которым может быть произведена фильтрация
// и перечислимый тип значениями которого могут быть эти константы

package cons

//go:generate enummethods -type FilterNotifField
var (
	filterNotifField = [...]string{
		"uuid",
		"reg_number",
		"category",
		"status",
		"activity_status",
		"create_time",
		"updated",
		"event_description",
		"detect_time",
		"end_time",
		"inrelated_observables_domain",
		"inrelated_observables_uri",
		"inrelated_observables_ipv4",
		"inrelated_observables_ipv6",
		"inrelated_indicators_domain_malware",
		"inrelated_indicators_domain",
		"inrelated_indicators_ipv4",
		"inrelated_indicators_ipv6",
		"inrelated_indicators_ipv4_malware",
		"inrelated_indicators_ipv6_malware",
		"inrelated_indicators_hash",
		"inrelated_indicators_email",
		"inrelated_indicators_uri",
		"related_observables_email",
		"related_observables_uri",
		"related_indicators_vuln",
		"related_indicators_as_lir",
		"related_indicators_asn",
		"related_indicators_as",
		"related_observables_service",
		"related_observables_as_path",
		"product_info",
		"vulnerability_id",
		"product_category",
		"malware_hash",
		"affected_system_name",
		"affected_system_category",
		"affected_system_function",
		"affected_system_connection",
		"availability_impact",
		"integrity_impact",
		"confidentiality_impact",
		"custom_impact",
	}
)

// DocFilterField представляет название поля по которому производится фильтрация message
type FilterNotifField int

const (
	Uuid = FilterNotifField(iota)
	Regnumber
	Category
	Status
	ActStatus
	NotifCreateTime
	Updated
	EventDescr
	DetectTime
	EndTime
	InrelatedObservablesDomain
	InrelatedObservablesUri
	InrelatedObservablesIpv4
	InrelatedObservablesIpv6
	InrelatedIndicatorsDomainMalware
	InrelatedIndicatorsDomain
	InrelatedIndicatorsIpv4
	InrelatedIndicatorsIpv6
	InrelatedIndicatorsIpv4Malware
	InrelatedIndicatorsIpv6Malware
	InrelatedIndicatorsHash
	InrelatedIndicatorsEmail
	InrelatedIndicatorsUri
	RelatedObservablesEmail
	RelatedObservablesUri
	RelatedIndicatorsVuln
	RelatedIndicatorsAsLir
	RelatedIndicatorsAsn
	RelatedIndicatorsAs
	RelatedObservablesService
	RelatedObservablesAsPath
	ProductInfo
	VulnerabilityId
	ProductCategory
	MalwareHash
	AffectedSystemName
	AffectedSystemCategory
	AffectedSystemfunction
	AffectedSystemConnection
	AvailabilityImpact
	IntegrityImpact
	ConfidentialityImpact
	CustomImpact
)
