


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceResult struct {
	CreatedAt          *string           `pulumi:"createdAt"`
	Critical           *bool             `pulumi:"critical"`
	DataCenter         *string           `pulumi:"dataCenter"`
	Enabled            *bool             `pulumi:"enabled"`
	Id                 string            `pulumi:"id"`
	Location           *string           `pulumi:"location"`
	MetricId           string            `pulumi:"metricId"`
	Name               string            `pulumi:"name"`
	NamespaceType      *string           `pulumi:"namespaceType"`
	ProvisioningState  *string           `pulumi:"provisioningState"`
	Region             *string           `pulumi:"region"`
	ScaleUnit          *string           `pulumi:"scaleUnit"`
	ServiceBusEndpoint *string           `pulumi:"serviceBusEndpoint"`
	Sku                *SkuResponse      `pulumi:"sku"`
	Status             *string           `pulumi:"status"`
	SubscriptionId     *string           `pulumi:"subscriptionId"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
	UpdatedAt          *string           `pulumi:"updatedAt"`
}

func LookupNamespaceOutput(ctx *pulumi.Context, args LookupNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceResult, error) {
			args := v.(LookupNamespaceArgs)
			r, err := LookupNamespace(ctx, &args, opts...)
			var s LookupNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceResultOutput)
}

type LookupNamespaceOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceArgs)(nil)).Elem()
}


type LookupNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceResult)(nil)).Elem()
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutput() LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutputWithContext(ctx context.Context) LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Critical() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.Critical }).(pulumi.BoolPtrOutput)
}

func (o LookupNamespaceResultOutput) DataCenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.DataCenter }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.MetricId }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) NamespaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.NamespaceType }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) ScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.ScaleUnit }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) ServiceBusEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.ServiceBusEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupNamespaceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}
