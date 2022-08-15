


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getDatabase", args, &rv, opts...)
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
	AutoPauseDelay                *int              `pulumi:"autoPauseDelay"`
	CatalogCollation              *string           `pulumi:"catalogCollation"`
	Collation                     *string           `pulumi:"collation"`
	CreationDate                  string            `pulumi:"creationDate"`
	CurrentServiceObjectiveName   string            `pulumi:"currentServiceObjectiveName"`
	CurrentSku                    SkuResponse       `pulumi:"currentSku"`
	DatabaseId                    string            `pulumi:"databaseId"`
	DefaultSecondaryLocation      string            `pulumi:"defaultSecondaryLocation"`
	EarliestRestoreDate           string            `pulumi:"earliestRestoreDate"`
	ElasticPoolId                 *string           `pulumi:"elasticPoolId"`
	FailoverGroupId               string            `pulumi:"failoverGroupId"`
	HighAvailabilityReplicaCount  *int              `pulumi:"highAvailabilityReplicaCount"`
	Id                            string            `pulumi:"id"`
	Kind                          string            `pulumi:"kind"`
	LicenseType                   *string           `pulumi:"licenseType"`
	Location                      string            `pulumi:"location"`
	MaintenanceConfigurationId    *string           `pulumi:"maintenanceConfigurationId"`
	ManagedBy                     string            `pulumi:"managedBy"`
	MaxLogSizeBytes               float64           `pulumi:"maxLogSizeBytes"`
	MaxSizeBytes                  *float64          `pulumi:"maxSizeBytes"`
	MinCapacity                   *float64          `pulumi:"minCapacity"`
	Name                          string            `pulumi:"name"`
	PausedDate                    string            `pulumi:"pausedDate"`
	ReadScale                     *string           `pulumi:"readScale"`
	RequestedServiceObjectiveName string            `pulumi:"requestedServiceObjectiveName"`
	ResumedDate                   string            `pulumi:"resumedDate"`
	SecondaryType                 *string           `pulumi:"secondaryType"`
	Sku                           *SkuResponse      `pulumi:"sku"`
	Status                        string            `pulumi:"status"`
	StorageAccountType            *string           `pulumi:"storageAccountType"`
	Tags                          map[string]string `pulumi:"tags"`
	Type                          string            `pulumi:"type"`
	ZoneRedundant                 *bool             `pulumi:"zoneRedundant"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}


type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) AutoPauseDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.AutoPauseDelay }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseResultOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) CurrentServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CurrentServiceObjectiveName }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) CurrentSku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDatabaseResult) SkuResponse { return v.CurrentSku }).(SkuResponseOutput)
}

func (o LookupDatabaseResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ElasticPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ElasticPoolId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.FailoverGroupId }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) HighAvailabilityReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.HighAvailabilityReplicaCount }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) MaxLogSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDatabaseResult) float64 { return v.MaxLogSizeBytes }).(pulumi.Float64Output)
}

func (o LookupDatabaseResultOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *float64 { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupDatabaseResultOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *float64 { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) PausedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.PausedDate }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ReadScale }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) RequestedServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.RequestedServiceObjectiveName }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ResumedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ResumedDate }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) SecondaryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.SecondaryType }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
