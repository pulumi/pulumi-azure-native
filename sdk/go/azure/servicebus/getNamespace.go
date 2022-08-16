


package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus:getNamespace", args, &rv, opts...)
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
	CreatedAt          string            `pulumi:"createdAt"`
	Id                 string            `pulumi:"id"`
	Location           string            `pulumi:"location"`
	MetricId           string            `pulumi:"metricId"`
	Name               string            `pulumi:"name"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	ServiceBusEndpoint string            `pulumi:"serviceBusEndpoint"`
	Sku                *SBSkuResponse    `pulumi:"sku"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
	UpdatedAt          string            `pulumi:"updatedAt"`
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

func (o LookupNamespaceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.MetricId }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) ServiceBusEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ServiceBusEndpoint }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Sku() SBSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *SBSkuResponse { return v.Sku }).(SBSkuResponsePtrOutput)
}

func (o LookupNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}
