


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSharedKeys(ctx *pulumi.Context, args *GetSharedKeysArgs, opts ...pulumi.InvokeOption) (*GetSharedKeysResult, error) {
	var rv GetSharedKeysResult
	err := ctx.Invoke("azure-native:operationalinsights:getSharedKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSharedKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetSharedKeysResult struct {
	PrimarySharedKey   *string `pulumi:"primarySharedKey"`
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
}

func GetSharedKeysOutput(ctx *pulumi.Context, args GetSharedKeysOutputArgs, opts ...pulumi.InvokeOption) GetSharedKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSharedKeysResult, error) {
			args := v.(GetSharedKeysArgs)
			r, err := GetSharedKeys(ctx, &args, opts...)
			var s GetSharedKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSharedKeysResultOutput)
}

type GetSharedKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetSharedKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSharedKeysArgs)(nil)).Elem()
}


type GetSharedKeysResultOutput struct{ *pulumi.OutputState }

func (GetSharedKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSharedKeysResult)(nil)).Elem()
}

func (o GetSharedKeysResultOutput) ToGetSharedKeysResultOutput() GetSharedKeysResultOutput {
	return o
}

func (o GetSharedKeysResultOutput) ToGetSharedKeysResultOutputWithContext(ctx context.Context) GetSharedKeysResultOutput {
	return o
}

func (o GetSharedKeysResultOutput) PrimarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSharedKeysResult) *string { return v.PrimarySharedKey }).(pulumi.StringPtrOutput)
}

func (o GetSharedKeysResultOutput) SecondarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSharedKeysResult) *string { return v.SecondarySharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSharedKeysResultOutput{})
}
