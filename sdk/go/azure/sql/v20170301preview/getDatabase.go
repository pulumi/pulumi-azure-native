


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20170301preview:getDatabase", args, &rv, opts...)
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
	CatalogCollation            *string           `pulumi:"catalogCollation"`
	Collation                   *string           `pulumi:"collation"`
	CreationDate                string            `pulumi:"creationDate"`
	CurrentServiceObjectiveName string            `pulumi:"currentServiceObjectiveName"`
	DatabaseId                  string            `pulumi:"databaseId"`
	DefaultSecondaryLocation    string            `pulumi:"defaultSecondaryLocation"`
	ElasticPoolId               *string           `pulumi:"elasticPoolId"`
	FailoverGroupId             string            `pulumi:"failoverGroupId"`
	Id                          string            `pulumi:"id"`
	Kind                        string            `pulumi:"kind"`
	Location                    string            `pulumi:"location"`
	MaxSizeBytes                *float64          `pulumi:"maxSizeBytes"`
	Name                        string            `pulumi:"name"`
	Sku                         *SkuResponse      `pulumi:"sku"`
	Status                      string            `pulumi:"status"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        string            `pulumi:"type"`
	ZoneRedundant               *bool             `pulumi:"zoneRedundant"`
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

func (o LookupDatabaseResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) ElasticPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ElasticPoolId }).(pulumi.StringPtrOutput)
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

func (o LookupDatabaseResultOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *float64 { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
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
