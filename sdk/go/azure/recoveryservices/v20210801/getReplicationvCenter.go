


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationvCenter(ctx *pulumi.Context, args *LookupReplicationvCenterArgs, opts ...pulumi.InvokeOption) (*LookupReplicationvCenterResult, error) {
	var rv LookupReplicationvCenterResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210801:getReplicationvCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationvCenterArgs struct {
	FabricName        string `pulumi:"fabricName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
	VcenterName       string `pulumi:"vcenterName"`
}


type LookupReplicationvCenterResult struct {
	Id         string                    `pulumi:"id"`
	Location   *string                   `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties VCenterPropertiesResponse `pulumi:"properties"`
	Type       string                    `pulumi:"type"`
}

func LookupReplicationvCenterOutput(ctx *pulumi.Context, args LookupReplicationvCenterOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationvCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationvCenterResult, error) {
			args := v.(LookupReplicationvCenterArgs)
			r, err := LookupReplicationvCenter(ctx, &args, opts...)
			var s LookupReplicationvCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationvCenterResultOutput)
}

type LookupReplicationvCenterOutputArgs struct {
	FabricName        pulumi.StringInput `pulumi:"fabricName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
	VcenterName       pulumi.StringInput `pulumi:"vcenterName"`
}

func (LookupReplicationvCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationvCenterArgs)(nil)).Elem()
}


type LookupReplicationvCenterResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationvCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationvCenterResult)(nil)).Elem()
}

func (o LookupReplicationvCenterResultOutput) ToLookupReplicationvCenterResultOutput() LookupReplicationvCenterResultOutput {
	return o
}

func (o LookupReplicationvCenterResultOutput) ToLookupReplicationvCenterResultOutputWithContext(ctx context.Context) LookupReplicationvCenterResultOutput {
	return o
}

func (o LookupReplicationvCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationvCenterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationvCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationvCenterResultOutput) Properties() VCenterPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) VCenterPropertiesResponse { return v.Properties }).(VCenterPropertiesResponseOutput)
}

func (o LookupReplicationvCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationvCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationvCenterResultOutput{})
}
