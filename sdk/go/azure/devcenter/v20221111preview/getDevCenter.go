


package v20221111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDevCenter(ctx *pulumi.Context, args *LookupDevCenterArgs, opts ...pulumi.InvokeOption) (*LookupDevCenterResult, error) {
	var rv LookupDevCenterResult
	err := ctx.Invoke("azure-native:devcenter/v20221111preview:getDevCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDevCenterArgs struct {
	DevCenterName     string `pulumi:"devCenterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDevCenterResult struct {
	DevCenterUri      string                          `pulumi:"devCenterUri"`
	Id                string                          `pulumi:"id"`
	Identity          *ManagedServiceIdentityResponse `pulumi:"identity"`
	Location          string                          `pulumi:"location"`
	Name              string                          `pulumi:"name"`
	ProvisioningState string                          `pulumi:"provisioningState"`
	SystemData        SystemDataResponse              `pulumi:"systemData"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
}

func LookupDevCenterOutput(ctx *pulumi.Context, args LookupDevCenterOutputArgs, opts ...pulumi.InvokeOption) LookupDevCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevCenterResult, error) {
			args := v.(LookupDevCenterArgs)
			r, err := LookupDevCenter(ctx, &args, opts...)
			var s LookupDevCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevCenterResultOutput)
}

type LookupDevCenterOutputArgs struct {
	DevCenterName     pulumi.StringInput `pulumi:"devCenterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDevCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevCenterArgs)(nil)).Elem()
}


type LookupDevCenterResultOutput struct{ *pulumi.OutputState }

func (LookupDevCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevCenterResult)(nil)).Elem()
}

func (o LookupDevCenterResultOutput) ToLookupDevCenterResultOutput() LookupDevCenterResultOutput {
	return o
}

func (o LookupDevCenterResultOutput) ToLookupDevCenterResultOutputWithContext(ctx context.Context) LookupDevCenterResultOutput {
	return o
}

func (o LookupDevCenterResultOutput) DevCenterUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.DevCenterUri }).(pulumi.StringOutput)
}

func (o LookupDevCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDevCenterResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDevCenterResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupDevCenterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDevCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDevCenterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDevCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDevCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDevCenterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDevCenterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDevCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevCenterResultOutput{})
}
