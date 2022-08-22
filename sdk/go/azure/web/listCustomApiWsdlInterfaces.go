


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCustomApiWsdlInterfaces(ctx *pulumi.Context, args *ListCustomApiWsdlInterfacesArgs, opts ...pulumi.InvokeOption) (*ListCustomApiWsdlInterfacesResult, error) {
	var rv ListCustomApiWsdlInterfacesResult
	err := ctx.Invoke("azure-native:web:listCustomApiWsdlInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCustomApiWsdlInterfacesArgs struct {
	Content        *string      `pulumi:"content"`
	ImportMethod   *string      `pulumi:"importMethod"`
	Location       string       `pulumi:"location"`
	Service        *WsdlService `pulumi:"service"`
	SubscriptionId *string      `pulumi:"subscriptionId"`
	Url            *string      `pulumi:"url"`
}


type ListCustomApiWsdlInterfacesResult struct {
	Value []WsdlServiceResponse `pulumi:"value"`
}

func ListCustomApiWsdlInterfacesOutput(ctx *pulumi.Context, args ListCustomApiWsdlInterfacesOutputArgs, opts ...pulumi.InvokeOption) ListCustomApiWsdlInterfacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCustomApiWsdlInterfacesResult, error) {
			args := v.(ListCustomApiWsdlInterfacesArgs)
			r, err := ListCustomApiWsdlInterfaces(ctx, &args, opts...)
			var s ListCustomApiWsdlInterfacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCustomApiWsdlInterfacesResultOutput)
}

type ListCustomApiWsdlInterfacesOutputArgs struct {
	Content        pulumi.StringPtrInput `pulumi:"content"`
	ImportMethod   pulumi.StringPtrInput `pulumi:"importMethod"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Service        WsdlServicePtrInput   `pulumi:"service"`
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	Url            pulumi.StringPtrInput `pulumi:"url"`
}

func (ListCustomApiWsdlInterfacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCustomApiWsdlInterfacesArgs)(nil)).Elem()
}


type ListCustomApiWsdlInterfacesResultOutput struct{ *pulumi.OutputState }

func (ListCustomApiWsdlInterfacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCustomApiWsdlInterfacesResult)(nil)).Elem()
}

func (o ListCustomApiWsdlInterfacesResultOutput) ToListCustomApiWsdlInterfacesResultOutput() ListCustomApiWsdlInterfacesResultOutput {
	return o
}

func (o ListCustomApiWsdlInterfacesResultOutput) ToListCustomApiWsdlInterfacesResultOutputWithContext(ctx context.Context) ListCustomApiWsdlInterfacesResultOutput {
	return o
}

func (o ListCustomApiWsdlInterfacesResultOutput) Value() WsdlServiceResponseArrayOutput {
	return o.ApplyT(func(v ListCustomApiWsdlInterfacesResult) []WsdlServiceResponse { return v.Value }).(WsdlServiceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCustomApiWsdlInterfacesResultOutput{})
}
