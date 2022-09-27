


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppHostNameBinding(ctx *pulumi.Context, args *LookupWebAppHostNameBindingArgs, opts ...pulumi.InvokeOption) (*LookupWebAppHostNameBindingResult, error) {
	var rv LookupWebAppHostNameBindingResult
	err := ctx.Invoke("azure-native:web/v20220301:getWebAppHostNameBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppHostNameBindingArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppHostNameBindingResult struct {
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

func LookupWebAppHostNameBindingOutput(ctx *pulumi.Context, args LookupWebAppHostNameBindingOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppHostNameBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppHostNameBindingResult, error) {
			args := v.(LookupWebAppHostNameBindingArgs)
			r, err := LookupWebAppHostNameBinding(ctx, &args, opts...)
			var s LookupWebAppHostNameBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppHostNameBindingResultOutput)
}

type LookupWebAppHostNameBindingOutputArgs struct {
	HostName          pulumi.StringInput `pulumi:"hostName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppHostNameBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHostNameBindingArgs)(nil)).Elem()
}


type LookupWebAppHostNameBindingResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppHostNameBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppHostNameBindingResult)(nil)).Elem()
}

func (o LookupWebAppHostNameBindingResultOutput) ToLookupWebAppHostNameBindingResultOutput() LookupWebAppHostNameBindingResultOutput {
	return o
}

func (o LookupWebAppHostNameBindingResultOutput) ToLookupWebAppHostNameBindingResultOutputWithContext(ctx context.Context) LookupWebAppHostNameBindingResultOutput {
	return o
}

func (o LookupWebAppHostNameBindingResultOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.SslState }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppHostNameBindingResultOutput) VirtualIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppHostNameBindingResult) string { return v.VirtualIP }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppHostNameBindingResultOutput{})
}
