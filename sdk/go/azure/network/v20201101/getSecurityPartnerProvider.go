


package v20201101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityPartnerProvider(ctx *pulumi.Context, args *LookupSecurityPartnerProviderArgs, opts ...pulumi.InvokeOption) (*LookupSecurityPartnerProviderResult, error) {
	var rv LookupSecurityPartnerProviderResult
	err := ctx.Invoke("azure-native:network/v20201101:getSecurityPartnerProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityPartnerProviderArgs struct {
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	SecurityPartnerProviderName string `pulumi:"securityPartnerProviderName"`
}


type LookupSecurityPartnerProviderResult struct {
	ConnectionStatus     string               `pulumi:"connectionStatus"`
	Etag                 string               `pulumi:"etag"`
	Id                   *string              `pulumi:"id"`
	Location             *string              `pulumi:"location"`
	Name                 string               `pulumi:"name"`
	ProvisioningState    string               `pulumi:"provisioningState"`
	SecurityProviderName *string              `pulumi:"securityProviderName"`
	Tags                 map[string]string    `pulumi:"tags"`
	Type                 string               `pulumi:"type"`
	VirtualHub           *SubResourceResponse `pulumi:"virtualHub"`
}

func LookupSecurityPartnerProviderOutput(ctx *pulumi.Context, args LookupSecurityPartnerProviderOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityPartnerProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityPartnerProviderResult, error) {
			args := v.(LookupSecurityPartnerProviderArgs)
			r, err := LookupSecurityPartnerProvider(ctx, &args, opts...)
			var s LookupSecurityPartnerProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityPartnerProviderResultOutput)
}

type LookupSecurityPartnerProviderOutputArgs struct {
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityPartnerProviderName pulumi.StringInput `pulumi:"securityPartnerProviderName"`
}

func (LookupSecurityPartnerProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityPartnerProviderArgs)(nil)).Elem()
}


type LookupSecurityPartnerProviderResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityPartnerProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityPartnerProviderResult)(nil)).Elem()
}

func (o LookupSecurityPartnerProviderResultOutput) ToLookupSecurityPartnerProviderResultOutput() LookupSecurityPartnerProviderResultOutput {
	return o
}

func (o LookupSecurityPartnerProviderResultOutput) ToLookupSecurityPartnerProviderResultOutputWithContext(ctx context.Context) LookupSecurityPartnerProviderResultOutput {
	return o
}

func (o LookupSecurityPartnerProviderResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *string { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSecurityPartnerProviderResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityPartnerProviderResultOutput{})
}
