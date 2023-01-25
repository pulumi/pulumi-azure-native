


package v20210915preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTarget(ctx *pulumi.Context, args *LookupTargetArgs, opts ...pulumi.InvokeOption) (*LookupTargetResult, error) {
	var rv LookupTargetResult
	err := ctx.Invoke("azure-native:chaos/v20210915preview:getTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetArgs struct {
	ParentProviderNamespace string `pulumi:"parentProviderNamespace"`
	ParentResourceName      string `pulumi:"parentResourceName"`
	ParentResourceType      string `pulumi:"parentResourceType"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	TargetName              string `pulumi:"targetName"`
}


type LookupTargetResult struct {
	Id         string             `pulumi:"id"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupTargetOutput(ctx *pulumi.Context, args LookupTargetOutputArgs, opts ...pulumi.InvokeOption) LookupTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetResult, error) {
			args := v.(LookupTargetArgs)
			r, err := LookupTarget(ctx, &args, opts...)
			var s LookupTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTargetResultOutput)
}

type LookupTargetOutputArgs struct {
	ParentProviderNamespace pulumi.StringInput `pulumi:"parentProviderNamespace"`
	ParentResourceName      pulumi.StringInput `pulumi:"parentResourceName"`
	ParentResourceType      pulumi.StringInput `pulumi:"parentResourceType"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	TargetName              pulumi.StringInput `pulumi:"targetName"`
}

func (LookupTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetArgs)(nil)).Elem()
}


type LookupTargetResultOutput struct{ *pulumi.OutputState }

func (LookupTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetResult)(nil)).Elem()
}

func (o LookupTargetResultOutput) ToLookupTargetResultOutput() LookupTargetResultOutput {
	return o
}

func (o LookupTargetResultOutput) ToLookupTargetResultOutputWithContext(ctx context.Context) LookupTargetResultOutput {
	return o
}

func (o LookupTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTargetResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupTargetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTargetResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTargetResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupTargetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTargetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTargetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetResultOutput{})
}
