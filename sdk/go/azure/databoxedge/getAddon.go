


package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAddon(ctx *pulumi.Context, args *LookupAddonArgs, opts ...pulumi.InvokeOption) (*LookupAddonResult, error) {
	var rv LookupAddonResult
	err := ctx.Invoke("azure-native:databoxedge:getAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupAddonResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
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
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoleName          pulumi.StringInput `pulumi:"roleName"`
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

func (o LookupAddonResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAddonResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAddonResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAddonResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAddonResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAddonResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddonResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddonResultOutput{})
}
