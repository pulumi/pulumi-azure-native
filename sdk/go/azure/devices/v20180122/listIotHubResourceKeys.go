


package v20180122

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotHubResourceKeys(ctx *pulumi.Context, args *ListIotHubResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysResult, error) {
	var rv ListIotHubResourceKeysResult
	err := ctx.Invoke("azure-native:devices/v20180122:listIotHubResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListIotHubResourceKeysResult struct {
	NextLink string                                           `pulumi:"nextLink"`
	Value    []SharedAccessSignatureAuthorizationRuleResponse `pulumi:"value"`
}

func ListIotHubResourceKeysOutput(ctx *pulumi.Context, args ListIotHubResourceKeysOutputArgs, opts ...pulumi.InvokeOption) ListIotHubResourceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotHubResourceKeysResult, error) {
			args := v.(ListIotHubResourceKeysArgs)
			r, err := ListIotHubResourceKeys(ctx, &args, opts...)
			var s ListIotHubResourceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotHubResourceKeysResultOutput)
}

type ListIotHubResourceKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListIotHubResourceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysArgs)(nil)).Elem()
}


type ListIotHubResourceKeysResultOutput struct{ *pulumi.OutputState }

func (ListIotHubResourceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysResult)(nil)).Elem()
}

func (o ListIotHubResourceKeysResultOutput) ToListIotHubResourceKeysResultOutput() ListIotHubResourceKeysResultOutput {
	return o
}

func (o ListIotHubResourceKeysResultOutput) ToListIotHubResourceKeysResultOutputWithContext(ctx context.Context) ListIotHubResourceKeysResultOutput {
	return o
}

func (o ListIotHubResourceKeysResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListIotHubResourceKeysResultOutput) Value() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysResult) []SharedAccessSignatureAuthorizationRuleResponse { return v.Value }).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotHubResourceKeysResultOutput{})
}
