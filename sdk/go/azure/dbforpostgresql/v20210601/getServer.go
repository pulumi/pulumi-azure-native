


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20210601:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin       *string                    `pulumi:"administratorLogin"`
	AvailabilityZone         *string                    `pulumi:"availabilityZone"`
	Backup                   *BackupResponse            `pulumi:"backup"`
	FullyQualifiedDomainName string                     `pulumi:"fullyQualifiedDomainName"`
	HighAvailability         *HighAvailabilityResponse  `pulumi:"highAvailability"`
	Id                       string                     `pulumi:"id"`
	Location                 string                     `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	MinorVersion             string                     `pulumi:"minorVersion"`
	Name                     string                     `pulumi:"name"`
	Network                  *NetworkResponse           `pulumi:"network"`
	Sku                      *SkuResponse               `pulumi:"sku"`
	State                    string                     `pulumi:"state"`
	Storage                  *StorageResponse           `pulumi:"storage"`
	SystemData               SystemDataResponse         `pulumi:"systemData"`
	Tags                     map[string]string          `pulumi:"tags"`
	Type                     string                     `pulumi:"type"`
	Version                  *string                    `pulumi:"version"`
}


func (val *LookupServerResult) Defaults() *LookupServerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AvailabilityZone) {
		availabilityZone_ := ""
		tmp.AvailabilityZone = &availabilityZone_
	}
	tmp.Backup = tmp.Backup.Defaults()

	tmp.HighAvailability = tmp.HighAvailability.Defaults()

	tmp.MaintenanceWindow = tmp.MaintenanceWindow.Defaults()

	tmp.Network = tmp.Network.Defaults()

	return &tmp
}
