


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupDatabaseResult struct {
	AutoPauseDelay                   *int              `pulumi:"autoPauseDelay"`
	CatalogCollation                 *string           `pulumi:"catalogCollation"`
	Collation                        *string           `pulumi:"collation"`
	CreationDate                     string            `pulumi:"creationDate"`
	CurrentBackupStorageRedundancy   string            `pulumi:"currentBackupStorageRedundancy"`
	CurrentServiceObjectiveName      string            `pulumi:"currentServiceObjectiveName"`
	CurrentSku                       SkuResponse       `pulumi:"currentSku"`
	DatabaseId                       string            `pulumi:"databaseId"`
	DefaultSecondaryLocation         string            `pulumi:"defaultSecondaryLocation"`
	EarliestRestoreDate              string            `pulumi:"earliestRestoreDate"`
	ElasticPoolId                    *string           `pulumi:"elasticPoolId"`
	FailoverGroupId                  string            `pulumi:"failoverGroupId"`
	HighAvailabilityReplicaCount     *int              `pulumi:"highAvailabilityReplicaCount"`
	Id                               string            `pulumi:"id"`
	IsInfraEncryptionEnabled         bool              `pulumi:"isInfraEncryptionEnabled"`
	IsLedgerOn                       *bool             `pulumi:"isLedgerOn"`
	Kind                             string            `pulumi:"kind"`
	LicenseType                      *string           `pulumi:"licenseType"`
	Location                         string            `pulumi:"location"`
	MaintenanceConfigurationId       *string           `pulumi:"maintenanceConfigurationId"`
	ManagedBy                        string            `pulumi:"managedBy"`
	MaxLogSizeBytes                  float64           `pulumi:"maxLogSizeBytes"`
	MaxSizeBytes                     *float64          `pulumi:"maxSizeBytes"`
	MinCapacity                      *float64          `pulumi:"minCapacity"`
	Name                             string            `pulumi:"name"`
	PausedDate                       string            `pulumi:"pausedDate"`
	ReadScale                        *string           `pulumi:"readScale"`
	RequestedBackupStorageRedundancy *string           `pulumi:"requestedBackupStorageRedundancy"`
	RequestedServiceObjectiveName    string            `pulumi:"requestedServiceObjectiveName"`
	ResumedDate                      string            `pulumi:"resumedDate"`
	SecondaryType                    *string           `pulumi:"secondaryType"`
	Sku                              *SkuResponse      `pulumi:"sku"`
	Status                           string            `pulumi:"status"`
	Tags                             map[string]string `pulumi:"tags"`
	Type                             string            `pulumi:"type"`
	ZoneRedundant                    *bool             `pulumi:"zoneRedundant"`
}
