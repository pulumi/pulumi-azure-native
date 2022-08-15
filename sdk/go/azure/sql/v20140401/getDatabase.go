


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20140401:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	DatabaseName      string  `pulumi:"databaseName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
}


type LookupDatabaseResult struct {
	Collation                     *string                             `pulumi:"collation"`
	ContainmentState              float64                             `pulumi:"containmentState"`
	CreationDate                  string                              `pulumi:"creationDate"`
	CurrentServiceObjectiveId     string                              `pulumi:"currentServiceObjectiveId"`
	DatabaseId                    string                              `pulumi:"databaseId"`
	DefaultSecondaryLocation      string                              `pulumi:"defaultSecondaryLocation"`
	EarliestRestoreDate           string                              `pulumi:"earliestRestoreDate"`
	Edition                       *string                             `pulumi:"edition"`
	ElasticPoolName               *string                             `pulumi:"elasticPoolName"`
	FailoverGroupId               string                              `pulumi:"failoverGroupId"`
	Id                            string                              `pulumi:"id"`
	Kind                          string                              `pulumi:"kind"`
	Location                      string                              `pulumi:"location"`
	MaxSizeBytes                  *string                             `pulumi:"maxSizeBytes"`
	Name                          string                              `pulumi:"name"`
	ReadScale                     *string                             `pulumi:"readScale"`
	RecommendedIndex              []RecommendedIndexResponse          `pulumi:"recommendedIndex"`
	RequestedServiceObjectiveId   *string                             `pulumi:"requestedServiceObjectiveId"`
	RequestedServiceObjectiveName *string                             `pulumi:"requestedServiceObjectiveName"`
	ServiceLevelObjective         string                              `pulumi:"serviceLevelObjective"`
	ServiceTierAdvisors           []ServiceTierAdvisorResponse        `pulumi:"serviceTierAdvisors"`
	Status                        string                              `pulumi:"status"`
	Tags                          map[string]string                   `pulumi:"tags"`
	TransparentDataEncryption     []TransparentDataEncryptionResponse `pulumi:"transparentDataEncryption"`
	Type                          string                              `pulumi:"type"`
	ZoneRedundant                 *bool                               `pulumi:"zoneRedundant"`
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
	DatabaseName      pulumi.StringInput    `pulumi:"databaseName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput    `pulumi:"serverName"`
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

func (o LookupDatabaseResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) ContainmentState() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDatabaseResult) float64 { return v.ContainmentState }).(pulumi.Float64Output)
}

func (o LookupDatabaseResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) CurrentServiceObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.CurrentServiceObjectiveId }).(pulumi.StringOutput)
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

func (o LookupDatabaseResultOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.FailoverGroupId }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) MaxSizeBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.MaxSizeBytes }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ReadScale }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) RecommendedIndex() RecommendedIndexResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []RecommendedIndexResponse { return v.RecommendedIndex }).(RecommendedIndexResponseArrayOutput)
}

func (o LookupDatabaseResultOutput) RequestedServiceObjectiveId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.RequestedServiceObjectiveId }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) RequestedServiceObjectiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.RequestedServiceObjectiveName }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseResultOutput) ServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ServiceLevelObjective }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ServiceTierAdvisors() ServiceTierAdvisorResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []ServiceTierAdvisorResponse { return v.ServiceTierAdvisors }).(ServiceTierAdvisorResponseArrayOutput)
}

func (o LookupDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseResultOutput) TransparentDataEncryption() TransparentDataEncryptionResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []TransparentDataEncryptionResponse { return v.TransparentDataEncryption }).(TransparentDataEncryptionResponseArrayOutput)
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
