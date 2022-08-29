


package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGen1Environment(ctx *pulumi.Context, args *LookupGen1EnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupGen1EnvironmentResult, error) {
	var rv LookupGen1EnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getGen1Environment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGen1EnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupGen1EnvironmentResult struct {
	CreationTime                 string                         `pulumi:"creationTime"`
	DataAccessFqdn               string                         `pulumi:"dataAccessFqdn"`
	DataAccessId                 string                         `pulumi:"dataAccessId"`
	DataRetentionTime            string                         `pulumi:"dataRetentionTime"`
	Id                           string                         `pulumi:"id"`
	Kind                         string                         `pulumi:"kind"`
	Location                     string                         `pulumi:"location"`
	Name                         string                         `pulumi:"name"`
	PartitionKeyProperties       []TimeSeriesIdPropertyResponse `pulumi:"partitionKeyProperties"`
	ProvisioningState            string                         `pulumi:"provisioningState"`
	Sku                          SkuResponse                    `pulumi:"sku"`
	Status                       EnvironmentStatusResponse      `pulumi:"status"`
	StorageLimitExceededBehavior *string                        `pulumi:"storageLimitExceededBehavior"`
	Tags                         map[string]string              `pulumi:"tags"`
	Type                         string                         `pulumi:"type"`
}

func LookupGen1EnvironmentOutput(ctx *pulumi.Context, args LookupGen1EnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupGen1EnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGen1EnvironmentResult, error) {
			args := v.(LookupGen1EnvironmentArgs)
			r, err := LookupGen1Environment(ctx, &args, opts...)
			var s LookupGen1EnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGen1EnvironmentResultOutput)
}

type LookupGen1EnvironmentOutputArgs struct {
	EnvironmentName   pulumi.StringInput    `pulumi:"environmentName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupGen1EnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen1EnvironmentArgs)(nil)).Elem()
}


type LookupGen1EnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupGen1EnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGen1EnvironmentResult)(nil)).Elem()
}

func (o LookupGen1EnvironmentResultOutput) ToLookupGen1EnvironmentResultOutput() LookupGen1EnvironmentResultOutput {
	return o
}

func (o LookupGen1EnvironmentResultOutput) ToLookupGen1EnvironmentResultOutputWithContext(ctx context.Context) LookupGen1EnvironmentResultOutput {
	return o
}

func (o LookupGen1EnvironmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.DataAccessId }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) DataRetentionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.DataRetentionTime }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) PartitionKeyProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) []TimeSeriesIdPropertyResponse { return v.PartitionKeyProperties }).(TimeSeriesIdPropertyResponseArrayOutput)
}

func (o LookupGen1EnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGen1EnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupGen1EnvironmentResultOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) EnvironmentStatusResponse { return v.Status }).(EnvironmentStatusResponseOutput)
}

func (o LookupGen1EnvironmentResultOutput) StorageLimitExceededBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) *string { return v.StorageLimitExceededBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupGen1EnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGen1EnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGen1EnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGen1EnvironmentResultOutput{})
}
