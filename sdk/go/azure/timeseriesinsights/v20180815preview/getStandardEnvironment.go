


package v20180815preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStandardEnvironment(ctx *pulumi.Context, args *LookupStandardEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupStandardEnvironmentResult, error) {
	var rv LookupStandardEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20180815preview:getStandardEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStandardEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupStandardEnvironmentResult struct {
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

func LookupStandardEnvironmentOutput(ctx *pulumi.Context, args LookupStandardEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupStandardEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStandardEnvironmentResult, error) {
			args := v.(LookupStandardEnvironmentArgs)
			r, err := LookupStandardEnvironment(ctx, &args, opts...)
			var s LookupStandardEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStandardEnvironmentResultOutput)
}

type LookupStandardEnvironmentOutputArgs struct {
	EnvironmentName   pulumi.StringInput    `pulumi:"environmentName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupStandardEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardEnvironmentArgs)(nil)).Elem()
}


type LookupStandardEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupStandardEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardEnvironmentResult)(nil)).Elem()
}

func (o LookupStandardEnvironmentResultOutput) ToLookupStandardEnvironmentResultOutput() LookupStandardEnvironmentResultOutput {
	return o
}

func (o LookupStandardEnvironmentResultOutput) ToLookupStandardEnvironmentResultOutputWithContext(ctx context.Context) LookupStandardEnvironmentResultOutput {
	return o
}

func (o LookupStandardEnvironmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.DataAccessId }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) DataRetentionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.DataRetentionTime }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) PartitionKeyProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) []TimeSeriesIdPropertyResponse {
		return v.PartitionKeyProperties
	}).(TimeSeriesIdPropertyResponseArrayOutput)
}

func (o LookupStandardEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStandardEnvironmentResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupStandardEnvironmentResultOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) EnvironmentStatusResponse { return v.Status }).(EnvironmentStatusResponseOutput)
}

func (o LookupStandardEnvironmentResultOutput) StorageLimitExceededBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) *string { return v.StorageLimitExceededBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupStandardEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStandardEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStandardEnvironmentResultOutput{})
}
