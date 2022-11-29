


package v20171012

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	var rv LookupCustomDomainResult
	err := ctx.Invoke("azure-native:cdn/v20171012:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	CustomDomainName  string `pulumi:"customDomainName"`
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCustomDomainResult struct {
	CustomHttpsProvisioningState    string  `pulumi:"customHttpsProvisioningState"`
	CustomHttpsProvisioningSubstate string  `pulumi:"customHttpsProvisioningSubstate"`
	HostName                        string  `pulumi:"hostName"`
	Id                              string  `pulumi:"id"`
	Name                            string  `pulumi:"name"`
	ProvisioningState               string  `pulumi:"provisioningState"`
	ResourceState                   string  `pulumi:"resourceState"`
	Type                            string  `pulumi:"type"`
	ValidationData                  *string `pulumi:"validationData"`
}

func LookupCustomDomainOutput(ctx *pulumi.Context, args LookupCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomDomainResult, error) {
			args := v.(LookupCustomDomainArgs)
			r, err := LookupCustomDomain(ctx, &args, opts...)
			var s LookupCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomDomainResultOutput)
}

type LookupCustomDomainOutputArgs struct {
	CustomDomainName  pulumi.StringInput `pulumi:"customDomainName"`
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainArgs)(nil)).Elem()
}


type LookupCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainResult)(nil)).Elem()
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutput() LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutputWithContext(ctx context.Context) LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) CustomHttpsProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.CustomHttpsProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) CustomHttpsProvisioningSubstate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.CustomHttpsProvisioningSubstate }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) ValidationData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) *string { return v.ValidationData }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomDomainResultOutput{})
}
