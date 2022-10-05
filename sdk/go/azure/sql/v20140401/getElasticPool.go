


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupElasticPool(ctx *pulumi.Context, args *LookupElasticPoolArgs, opts ...pulumi.InvokeOption) (*LookupElasticPoolResult, error) {
	var rv LookupElasticPoolResult
	err := ctx.Invoke("azure-native:sql/v20140401:getElasticPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupElasticPoolArgs struct {
	ElasticPoolName   string `pulumi:"elasticPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupElasticPoolResult struct {
	CreationDate   string            `pulumi:"creationDate"`
	DatabaseDtuMax *int              `pulumi:"databaseDtuMax"`
	DatabaseDtuMin *int              `pulumi:"databaseDtuMin"`
	Dtu            *int              `pulumi:"dtu"`
	Edition        *string           `pulumi:"edition"`
	Id             string            `pulumi:"id"`
	Kind           string            `pulumi:"kind"`
	Location       string            `pulumi:"location"`
	Name           string            `pulumi:"name"`
	State          string            `pulumi:"state"`
	StorageMB      *int              `pulumi:"storageMB"`
	Tags           map[string]string `pulumi:"tags"`
	Type           string            `pulumi:"type"`
	ZoneRedundant  *bool             `pulumi:"zoneRedundant"`
}

func LookupElasticPoolOutput(ctx *pulumi.Context, args LookupElasticPoolOutputArgs, opts ...pulumi.InvokeOption) LookupElasticPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticPoolResult, error) {
			args := v.(LookupElasticPoolArgs)
			r, err := LookupElasticPool(ctx, &args, opts...)
			var s LookupElasticPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupElasticPoolResultOutput)
}

type LookupElasticPoolOutputArgs struct {
	ElasticPoolName   pulumi.StringInput `pulumi:"elasticPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupElasticPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticPoolArgs)(nil)).Elem()
}


type LookupElasticPoolResultOutput struct{ *pulumi.OutputState }

func (LookupElasticPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticPoolResult)(nil)).Elem()
}

func (o LookupElasticPoolResultOutput) ToLookupElasticPoolResultOutput() LookupElasticPoolResultOutput {
	return o
}

func (o LookupElasticPoolResultOutput) ToLookupElasticPoolResultOutputWithContext(ctx context.Context) LookupElasticPoolResultOutput {
	return o
}

func (o LookupElasticPoolResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) DatabaseDtuMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *int { return v.DatabaseDtuMax }).(pulumi.IntPtrOutput)
}

func (o LookupElasticPoolResultOutput) DatabaseDtuMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *int { return v.DatabaseDtuMin }).(pulumi.IntPtrOutput)
}

func (o LookupElasticPoolResultOutput) Dtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *int { return v.Dtu }).(pulumi.IntPtrOutput)
}

func (o LookupElasticPoolResultOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o LookupElasticPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *int { return v.StorageMB }).(pulumi.IntPtrOutput)
}

func (o LookupElasticPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupElasticPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupElasticPoolResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupElasticPoolResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticPoolResultOutput{})
}
