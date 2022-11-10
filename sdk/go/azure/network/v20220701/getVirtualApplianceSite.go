


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualApplianceSite(ctx *pulumi.Context, args *LookupVirtualApplianceSiteArgs, opts ...pulumi.InvokeOption) (*LookupVirtualApplianceSiteResult, error) {
	var rv LookupVirtualApplianceSiteResult
	err := ctx.Invoke("azure-native:network/v20220701:getVirtualApplianceSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualApplianceSiteArgs struct {
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	SiteName                    string `pulumi:"siteName"`
}


type LookupVirtualApplianceSiteResult struct {
	AddressPrefix     *string                            `pulumi:"addressPrefix"`
	Etag              string                             `pulumi:"etag"`
	Id                *string                            `pulumi:"id"`
	Name              *string                            `pulumi:"name"`
	O365Policy        *Office365PolicyPropertiesResponse `pulumi:"o365Policy"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Type              string                             `pulumi:"type"`
}

func LookupVirtualApplianceSiteOutput(ctx *pulumi.Context, args LookupVirtualApplianceSiteOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualApplianceSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualApplianceSiteResult, error) {
			args := v.(LookupVirtualApplianceSiteArgs)
			r, err := LookupVirtualApplianceSite(ctx, &args, opts...)
			var s LookupVirtualApplianceSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualApplianceSiteResultOutput)
}

type LookupVirtualApplianceSiteOutputArgs struct {
	NetworkVirtualApplianceName pulumi.StringInput `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	SiteName                    pulumi.StringInput `pulumi:"siteName"`
}

func (LookupVirtualApplianceSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualApplianceSiteArgs)(nil)).Elem()
}


type LookupVirtualApplianceSiteResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualApplianceSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualApplianceSiteResult)(nil)).Elem()
}

func (o LookupVirtualApplianceSiteResultOutput) ToLookupVirtualApplianceSiteResultOutput() LookupVirtualApplianceSiteResultOutput {
	return o
}

func (o LookupVirtualApplianceSiteResultOutput) ToLookupVirtualApplianceSiteResultOutputWithContext(ctx context.Context) LookupVirtualApplianceSiteResultOutput {
	return o
}

func (o LookupVirtualApplianceSiteResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualApplianceSiteResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualApplianceSiteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualApplianceSiteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualApplianceSiteResultOutput) O365Policy() Office365PolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *Office365PolicyPropertiesResponse { return v.O365Policy }).(Office365PolicyPropertiesResponsePtrOutput)
}

func (o LookupVirtualApplianceSiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualApplianceSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualApplianceSiteResultOutput{})
}
