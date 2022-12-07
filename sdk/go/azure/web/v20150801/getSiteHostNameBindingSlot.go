


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteHostNameBindingSlot(ctx *pulumi.Context, args *LookupSiteHostNameBindingSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteHostNameBindingSlotResult, error) {
	var rv LookupSiteHostNameBindingSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteHostNameBindingSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteHostNameBindingSlotArgs struct {
	HostName          string `pulumi:"hostName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteHostNameBindingSlotResult struct {
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

func LookupSiteHostNameBindingSlotOutput(ctx *pulumi.Context, args LookupSiteHostNameBindingSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteHostNameBindingSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteHostNameBindingSlotResult, error) {
			args := v.(LookupSiteHostNameBindingSlotArgs)
			r, err := LookupSiteHostNameBindingSlot(ctx, &args, opts...)
			var s LookupSiteHostNameBindingSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteHostNameBindingSlotResultOutput)
}

type LookupSiteHostNameBindingSlotOutputArgs struct {
	HostName          pulumi.StringInput `pulumi:"hostName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteHostNameBindingSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteHostNameBindingSlotArgs)(nil)).Elem()
}


type LookupSiteHostNameBindingSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteHostNameBindingSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteHostNameBindingSlotResult)(nil)).Elem()
}

func (o LookupSiteHostNameBindingSlotResultOutput) ToLookupSiteHostNameBindingSlotResultOutput() LookupSiteHostNameBindingSlotResultOutput {
	return o
}

func (o LookupSiteHostNameBindingSlotResultOutput) ToLookupSiteHostNameBindingSlotResultOutputWithContext(ctx context.Context) LookupSiteHostNameBindingSlotResultOutput {
	return o
}

func (o LookupSiteHostNameBindingSlotResultOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteHostNameBindingSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteHostNameBindingSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteHostNameBindingSlotResultOutput{})
}
