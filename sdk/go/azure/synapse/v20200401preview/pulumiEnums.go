


package v20200401preview

type StorageRedundancy string

const (
	StorageRedundancyLocal   = StorageRedundancy("Local")
	StorageRedundancyGeo     = StorageRedundancy("Geo")
	StorageRedundancyZone    = StorageRedundancy("Zone")
	StorageRedundancyGeoZone = StorageRedundancy("GeoZone")
)

func init() {
}
