


package storagepool

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIscsiTarget(ctx *pulumi.Context, args *LookupIscsiTargetArgs, opts ...pulumi.InvokeOption) (*LookupIscsiTargetResult, error) {
	var rv LookupIscsiTargetResult
	err := ctx.Invoke("azure-native:storagepool:getIscsiTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiTargetArgs struct {
	DiskPoolName      string `pulumi:"diskPoolName"`
	IscsiTargetName   string `pulumi:"iscsiTargetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIscsiTargetResult struct {
	Id                string                      `pulumi:"id"`
	Name              string                      `pulumi:"name"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	Status            string                      `pulumi:"status"`
	TargetIqn         string                      `pulumi:"targetIqn"`
	Tpgs              []TargetPortalGroupResponse `pulumi:"tpgs"`
	Type              string                      `pulumi:"type"`
}

func LookupIscsiTargetOutput(ctx *pulumi.Context, args LookupIscsiTargetOutputArgs, opts ...pulumi.InvokeOption) LookupIscsiTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIscsiTargetResult, error) {
			args := v.(LookupIscsiTargetArgs)
			r, err := LookupIscsiTarget(ctx, &args, opts...)
			var s LookupIscsiTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIscsiTargetResultOutput)
}

type LookupIscsiTargetOutputArgs struct {
	DiskPoolName      pulumi.StringInput `pulumi:"diskPoolName"`
	IscsiTargetName   pulumi.StringInput `pulumi:"iscsiTargetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIscsiTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiTargetArgs)(nil)).Elem()
}


type LookupIscsiTargetResultOutput struct{ *pulumi.OutputState }

func (LookupIscsiTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiTargetResult)(nil)).Elem()
}

func (o LookupIscsiTargetResultOutput) ToLookupIscsiTargetResultOutput() LookupIscsiTargetResultOutput {
	return o
}

func (o LookupIscsiTargetResultOutput) ToLookupIscsiTargetResultOutputWithContext(ctx context.Context) LookupIscsiTargetResultOutput {
	return o
}

func (o LookupIscsiTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) TargetIqn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.TargetIqn }).(pulumi.StringOutput)
}

func (o LookupIscsiTargetResultOutput) Tpgs() TargetPortalGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) []TargetPortalGroupResponse { return v.Tpgs }).(TargetPortalGroupResponseArrayOutput)
}

func (o LookupIscsiTargetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiTargetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIscsiTargetResultOutput{})
}
