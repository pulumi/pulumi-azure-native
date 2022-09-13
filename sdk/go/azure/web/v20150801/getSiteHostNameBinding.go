


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteHostNameBinding(ctx *pulumi.Context, args *LookupSiteHostNameBindingArgs, opts ...pulumi.InvokeOption) (*LookupSiteHostNameBindingResult, error) {
	var rv LookupSiteHostNameBindingResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteHostNameBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteHostNameBindingArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteHostNameBindingResult struct {
	AzureResourceName           *string           `pulumi:"azureResourceName"`
	AzureResourceType           *string           `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *string           `pulumi:"customHostNameDnsRecordType"`
	DomainId                    *string           `pulumi:"domainId"`
	HostNameType                *string           `pulumi:"hostNameType"`
	Id                          *string           `pulumi:"id"`
	Kind                        *string           `pulumi:"kind"`
	Location                    string            `pulumi:"location"`
	Name                        *string           `pulumi:"name"`
	SiteName                    *string           `pulumi:"siteName"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        *string           `pulumi:"type"`
}

func LookupSiteHostNameBindingOutput(ctx *pulumi.Context, args LookupSiteHostNameBindingOutputArgs, opts ...pulumi.InvokeOption) LookupSiteHostNameBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteHostNameBindingResult, error) {
			args := v.(LookupSiteHostNameBindingArgs)
			r, err := LookupSiteHostNameBinding(ctx, &args, opts...)
			var s LookupSiteHostNameBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteHostNameBindingResultOutput)
}

type LookupSiteHostNameBindingOutputArgs struct {
	HostName          pulumi.StringInput `pulumi:"hostName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteHostNameBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteHostNameBindingArgs)(nil)).Elem()
}


type LookupSiteHostNameBindingResultOutput struct{ *pulumi.OutputState }

func (LookupSiteHostNameBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteHostNameBindingResult)(nil)).Elem()
}

func (o LookupSiteHostNameBindingResultOutput) ToLookupSiteHostNameBindingResultOutput() LookupSiteHostNameBindingResultOutput {
	return o
}

func (o LookupSiteHostNameBindingResultOutput) ToLookupSiteHostNameBindingResultOutputWithContext(ctx context.Context) LookupSiteHostNameBindingResultOutput {
	return o
}

func (o LookupSiteHostNameBindingResultOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteHostNameBindingResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteHostNameBindingResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteHostNameBindingResultOutput{})
}
