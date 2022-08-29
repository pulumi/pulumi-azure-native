


package v20200717preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAddon(ctx *pulumi.Context, args *LookupAddonArgs, opts ...pulumi.InvokeOption) (*LookupAddonResult, error) {
	var rv LookupAddonResult
	err := ctx.Invoke("azure-native:avs/v20200717preview:getAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAddonResult struct {
	AddonType         *string `pulumi:"addonType"`
	Id                string  `pulumi:"id"`
	LicenseKey        *string `pulumi:"licenseKey"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}

func LookupAddonOutput(ctx *pulumi.Context, args LookupAddonOutputArgs, opts ...pulumi.InvokeOption) LookupAddonResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAddonResult, error) {
			args := v.(LookupAddonArgs)
			r, err := LookupAddon(ctx, &args, opts...)
			var s LookupAddonResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAddonResultOutput)
}

type LookupAddonOutputArgs struct {
	AddonName         pulumi.StringInput `pulumi:"addonName"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAddonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddonArgs)(nil)).Elem()
}


type LookupAddonResultOutput struct{ *pulumi.OutputState }

func (LookupAddonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddonResult)(nil)).Elem()
}

func (o LookupAddonResultOutput) ToLookupAddonResultOutput() LookupAddonResultOutput {
	return o
}

func (o LookupAddonResultOutput) ToLookupAddonResultOutputWithContext(ctx context.Context) LookupAddonResultOutput {
	return o
}

func (o LookupAddonResultOutput) AddonType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAddonResult) *string { return v.AddonType }).(pulumi.StringPtrOutput)
}

func (o LookupAddonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAddonResultOutput) LicenseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAddonResult) *string { return v.LicenseKey }).(pulumi.StringPtrOutput)
}

func (o LookupAddonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAddonResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAddonResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddonResultOutput{})
}
