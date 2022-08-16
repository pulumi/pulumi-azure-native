


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDigitalTwin(ctx *pulumi.Context, args *LookupDigitalTwinArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinResult, error) {
	var rv LookupDigitalTwinResult
	err := ctx.Invoke("azure-native:digitaltwins/v20200301preview:getDigitalTwin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDigitalTwinResult struct {
	CreatedTime       string                       `pulumi:"createdTime"`
	HostName          string                       `pulumi:"hostName"`
	Id                string                       `pulumi:"id"`
	LastUpdatedTime   string                       `pulumi:"lastUpdatedTime"`
	Location          string                       `pulumi:"location"`
	Name              string                       `pulumi:"name"`
	ProvisioningState string                       `pulumi:"provisioningState"`
	Sku               *DigitalTwinsSkuInfoResponse `pulumi:"sku"`
	Tags              map[string]string            `pulumi:"tags"`
	Type              string                       `pulumi:"type"`
}

func LookupDigitalTwinOutput(ctx *pulumi.Context, args LookupDigitalTwinOutputArgs, opts ...pulumi.InvokeOption) LookupDigitalTwinResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDigitalTwinResult, error) {
			args := v.(LookupDigitalTwinArgs)
			r, err := LookupDigitalTwin(ctx, &args, opts...)
			var s LookupDigitalTwinResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDigitalTwinResultOutput)
}

type LookupDigitalTwinOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupDigitalTwinOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinArgs)(nil)).Elem()
}


type LookupDigitalTwinResultOutput struct{ *pulumi.OutputState }

func (LookupDigitalTwinResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinResult)(nil)).Elem()
}

func (o LookupDigitalTwinResultOutput) ToLookupDigitalTwinResultOutput() LookupDigitalTwinResultOutput {
	return o
}

func (o LookupDigitalTwinResultOutput) ToLookupDigitalTwinResultOutputWithContext(ctx context.Context) LookupDigitalTwinResultOutput {
	return o
}

func (o LookupDigitalTwinResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Sku() DigitalTwinsSkuInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) *DigitalTwinsSkuInfoResponse { return v.Sku }).(DigitalTwinsSkuInfoResponsePtrOutput)
}

func (o LookupDigitalTwinResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDigitalTwinResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDigitalTwinResultOutput{})
}
