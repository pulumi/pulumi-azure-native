


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMoveResource(ctx *pulumi.Context, args *LookupMoveResourceArgs, opts ...pulumi.InvokeOption) (*LookupMoveResourceResult, error) {
	var rv LookupMoveResourceResult
	err := ctx.Invoke("azure-native:migrate/v20191001preview:getMoveResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMoveResourceArgs struct {
	MoveCollectionName string `pulumi:"moveCollectionName"`
	MoveResourceName   string `pulumi:"moveResourceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMoveResourceResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties MoveResourcePropertiesResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}

func LookupMoveResourceOutput(ctx *pulumi.Context, args LookupMoveResourceOutputArgs, opts ...pulumi.InvokeOption) LookupMoveResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMoveResourceResult, error) {
			args := v.(LookupMoveResourceArgs)
			r, err := LookupMoveResource(ctx, &args, opts...)
			var s LookupMoveResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMoveResourceResultOutput)
}

type LookupMoveResourceOutputArgs struct {
	MoveCollectionName pulumi.StringInput `pulumi:"moveCollectionName"`
	MoveResourceName   pulumi.StringInput `pulumi:"moveResourceName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMoveResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMoveResourceArgs)(nil)).Elem()
}


type LookupMoveResourceResultOutput struct{ *pulumi.OutputState }

func (LookupMoveResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMoveResourceResult)(nil)).Elem()
}

func (o LookupMoveResourceResultOutput) ToLookupMoveResourceResultOutput() LookupMoveResourceResultOutput {
	return o
}

func (o LookupMoveResourceResultOutput) ToLookupMoveResourceResultOutputWithContext(ctx context.Context) LookupMoveResourceResultOutput {
	return o
}

func (o LookupMoveResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMoveResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMoveResourceResultOutput) Properties() MoveResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) MoveResourcePropertiesResponse { return v.Properties }).(MoveResourcePropertiesResponseOutput)
}

func (o LookupMoveResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMoveResourceResultOutput{})
}
