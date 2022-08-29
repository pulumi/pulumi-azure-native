


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppHostNameBindingSlot(ctx *pulumi.Context, args *LookupWebAppHostNameBindingSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHostNameBindingSlotResult, error) {
	var rv LookupWebAppHostNameBindingSlotResult
	err := ctx.Invoke("azure-native:web/v20200601:getWebAppHostNameBindingSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHostNameBindingSlotArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppHostNameBindingSlotResult struct {
	AzureResourceName           *string `pulumi:"azureResourceName"`
	AzureResourceType           *string `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *string `pulumi:"customHostNameDnsRecordType"`
	DomainId                    *string `pulumi:"domainId"`
	HostNameType                *string `pulumi:"hostNameType"`
	Id                          string  `pulumi:"id"`
	Kind                        *string `pulumi:"kind"`
	Name                        string  `pulumi:"name"`
	SiteName                    *string `pulumi:"siteName"`
	SslState                    *string `pulumi:"sslState"`
	Thumbprint                  *string `pulumi:"thumbprint"`
	Type                        string  `pulumi:"type"`
	VirtualIP                   string  `pulumi:"virtualIP"`
}

func LookupWebAppHostNameBindingSlotOutput(ctx *pulumi.Context, args LookupWebAppHostNameBindingSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppHostNameBindingSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppHostNameBindingSlotResult, error) {
			args := v.(LookupWebAppHostNameBindingSlotArgs)
			r, err := LookupWebAppHostNameBindingSlot(ctx, &args, opts...)
			var s LookupWebAppHostNameBindingSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppHostNameBindingSlotResultOutput)
}

type LookupWebAppHostNameBindingSlotOutputArgs struct {
	HostName          pulumi.StringInput `pulumi:"hostName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppHostNameBindingSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHostNameBindingSlotArgs)(nil)).Elem()
}


type LookupWebAppHostNameBindingSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppHostNameBindingSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHostNameBindingSlotResult)(nil)).Elem()
}

func (o LookupWebAppHostNameBindingSlotResultOutput) ToLookupWebAppHostNameBindingSlotResultOutput() LookupWebAppHostNameBindingSlotResultOutput {
	return o
}

func (o LookupWebAppHostNameBindingSlotResultOutput) ToLookupWebAppHostNameBindingSlotResultOutputWithContext(ctx context.Context) LookupWebAppHostNameBindingSlotResultOutput {
	return o
}

func (o LookupWebAppHostNameBindingSlotResultOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.SslState }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppHostNameBindingSlotResultOutput) VirtualIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingSlotResult) string { return v.VirtualIP }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppHostNameBindingSlotResultOutput{})
}
