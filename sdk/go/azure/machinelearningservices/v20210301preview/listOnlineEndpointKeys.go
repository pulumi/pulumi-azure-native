


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOnlineEndpointKeys(ctx *pulumi.Context, args *ListOnlineEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListOnlineEndpointKeysResult, error) {
	var rv ListOnlineEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listOnlineEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOnlineEndpointKeysArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListOnlineEndpointKeysResult struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListOnlineEndpointKeysOutput(ctx *pulumi.Context, args ListOnlineEndpointKeysOutputArgs, opts ...pulumi.InvokeOption) ListOnlineEndpointKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOnlineEndpointKeysResult, error) {
			args := v.(ListOnlineEndpointKeysArgs)
			r, err := ListOnlineEndpointKeys(ctx, &args, opts...)
			var s ListOnlineEndpointKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOnlineEndpointKeysResultOutput)
}

type ListOnlineEndpointKeysOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListOnlineEndpointKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOnlineEndpointKeysArgs)(nil)).Elem()
}


type ListOnlineEndpointKeysResultOutput struct{ *pulumi.OutputState }

func (ListOnlineEndpointKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOnlineEndpointKeysResult)(nil)).Elem()
}

func (o ListOnlineEndpointKeysResultOutput) ToListOnlineEndpointKeysResultOutput() ListOnlineEndpointKeysResultOutput {
	return o
}

func (o ListOnlineEndpointKeysResultOutput) ToListOnlineEndpointKeysResultOutputWithContext(ctx context.Context) ListOnlineEndpointKeysResultOutput {
	return o
}

func (o ListOnlineEndpointKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOnlineEndpointKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListOnlineEndpointKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOnlineEndpointKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOnlineEndpointKeysResultOutput{})
}
