


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArcAddon(ctx *pulumi.Context, args *LookupArcAddonArgs, opts ...pulumi.InvokeOption) (*LookupArcAddonResult, error) {
	var rv LookupArcAddonResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901preview:getArcAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArcAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupArcAddonResult struct {
	HostPlatform      string             `pulumi:"hostPlatform"`
	HostPlatformType  string             `pulumi:"hostPlatformType"`
	Id                string             `pulumi:"id"`
	Kind              string             `pulumi:"kind"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	ResourceLocation  string             `pulumi:"resourceLocation"`
	ResourceName      string             `pulumi:"resourceName"`
	SubscriptionId    string             `pulumi:"subscriptionId"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
	Version           string             `pulumi:"version"`
}

func LookupArcAddonOutput(ctx *pulumi.Context, args LookupArcAddonOutputArgs, opts ...pulumi.InvokeOption) LookupArcAddonResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArcAddonResult, error) {
			args := v.(LookupArcAddonArgs)
			r, err := LookupArcAddon(ctx, &args, opts...)
			var s LookupArcAddonResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArcAddonResultOutput)
}

type LookupArcAddonOutputArgs struct {
	AddonName         pulumi.StringInput `pulumi:"addonName"`
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoleName          pulumi.StringInput `pulumi:"roleName"`
}

func (LookupArcAddonOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcAddonArgs)(nil)).Elem()
}


type LookupArcAddonResultOutput struct{ *pulumi.OutputState }

func (LookupArcAddonResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcAddonResult)(nil)).Elem()
}

func (o LookupArcAddonResultOutput) ToLookupArcAddonResultOutput() LookupArcAddonResultOutput {
	return o
}

func (o LookupArcAddonResultOutput) ToLookupArcAddonResultOutputWithContext(ctx context.Context) LookupArcAddonResultOutput {
	return o
}

func (o LookupArcAddonResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ResourceLocation }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.ResourceName }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupArcAddonResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupArcAddonResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupArcAddonResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcAddonResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArcAddonResultOutput{})
}
