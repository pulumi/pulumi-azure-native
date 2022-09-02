


package networkfunction

type DestinationType string

const (
	DestinationTypeAzureMonitor = DestinationType("AzureMonitor")
)

type EmissionType string

const (
	EmissionTypeIPFIX = EmissionType("IPFIX")
)

type IngestionType string

const (
	IngestionTypeIPFIX = IngestionType("IPFIX")
)

type SourceType string

const (
	SourceTypeResource = SourceType("Resource")
)

func init() {
}
