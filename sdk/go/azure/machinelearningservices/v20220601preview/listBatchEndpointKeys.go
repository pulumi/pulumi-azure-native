


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBatchEndpointKeys(ctx *pulumi.Context, args *ListBatchEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListBatchEndpointKeysResult, error) {
	var rv ListBatchEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220601preview:listBatchEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBatchEndpointKeysArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListBatchEndpointKeysResult struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListBatchEndpointKeysOutput(ctx *pulumi.Context, args ListBatchEndpointKeysOutputArgs, opts ...pulumi.InvokeOption) ListBatchEndpointKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBatchEndpointKeysResult, error) {
			args := v.(ListBatchEndpointKeysArgs)
			r, err := ListBatchEndpointKeys(ctx, &args, opts...)
			var s ListBatchEndpointKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBatchEndpointKeysResultOutput)
}

type ListBatchEndpointKeysOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListBatchEndpointKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBatchEndpointKeysArgs)(nil)).Elem()
}


type ListBatchEndpointKeysResultOutput struct{ *pulumi.OutputState }

func (ListBatchEndpointKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBatchEndpointKeysResult)(nil)).Elem()
}

func (o ListBatchEndpointKeysResultOutput) ToListBatchEndpointKeysResultOutput() ListBatchEndpointKeysResultOutput {
	return o
}

func (o ListBatchEndpointKeysResultOutput) ToListBatchEndpointKeysResultOutputWithContext(ctx context.Context) ListBatchEndpointKeysResultOutput {
	return o
}

func (o ListBatchEndpointKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBatchEndpointKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListBatchEndpointKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBatchEndpointKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBatchEndpointKeysResultOutput{})
}
