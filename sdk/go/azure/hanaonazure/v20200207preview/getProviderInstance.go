


package v20200207preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProviderInstance(ctx *pulumi.Context, args *LookupProviderInstanceArgs, opts ...pulumi.InvokeOption) (*LookupProviderInstanceResult, error) {
	var rv LookupProviderInstanceResult
	err := ctx.Invoke("azure-native:hanaonazure/v20200207preview:getProviderInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProviderInstanceArgs struct {
	ProviderInstanceName string `pulumi:"providerInstanceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	SapMonitorName       string `pulumi:"sapMonitorName"`
}


type LookupProviderInstanceResult struct {
	Id                string  `pulumi:"id"`
	Metadata          *string `pulumi:"metadata"`
	Name              string  `pulumi:"name"`
	Properties        string  `pulumi:"properties"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}

func LookupProviderInstanceOutput(ctx *pulumi.Context, args LookupProviderInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupProviderInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProviderInstanceResult, error) {
			args := v.(LookupProviderInstanceArgs)
			r, err := LookupProviderInstance(ctx, &args, opts...)
			var s LookupProviderInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProviderInstanceResultOutput)
}

type LookupProviderInstanceOutputArgs struct {
	ProviderInstanceName pulumi.StringInput `pulumi:"providerInstanceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	SapMonitorName       pulumi.StringInput `pulumi:"sapMonitorName"`
}

func (LookupProviderInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderInstanceArgs)(nil)).Elem()
}


type LookupProviderInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupProviderInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProviderInstanceResult)(nil)).Elem()
}

func (o LookupProviderInstanceResultOutput) ToLookupProviderInstanceResultOutput() LookupProviderInstanceResultOutput {
	return o
}

func (o LookupProviderInstanceResultOutput) ToLookupProviderInstanceResultOutputWithContext(ctx context.Context) LookupProviderInstanceResultOutput {
	return o
}

func (o LookupProviderInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) Metadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) *string { return v.Metadata }).(pulumi.StringPtrOutput)
}

func (o LookupProviderInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) Properties() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Properties }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProviderInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProviderInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProviderInstanceResultOutput{})
}
