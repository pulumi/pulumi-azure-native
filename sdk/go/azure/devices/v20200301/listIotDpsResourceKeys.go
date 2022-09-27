


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotDpsResourceKeys(ctx *pulumi.Context, args *ListIotDpsResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotDpsResourceKeysResult, error) {
	var rv ListIotDpsResourceKeysResult
	err := ctx.Invoke("azure-native:devices/v20200301:listIotDpsResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotDpsResourceKeysArgs struct {
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type ListIotDpsResourceKeysResult struct {
	NextLink string                                                                  `pulumi:"nextLink"`
	Value    []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"value"`
}

func ListIotDpsResourceKeysOutput(ctx *pulumi.Context, args ListIotDpsResourceKeysOutputArgs, opts ...pulumi.InvokeOption) ListIotDpsResourceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotDpsResourceKeysResult, error) {
			args := v.(ListIotDpsResourceKeysArgs)
			r, err := ListIotDpsResourceKeys(ctx, &args, opts...)
			var s ListIotDpsResourceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotDpsResourceKeysResultOutput)
}

type ListIotDpsResourceKeysOutputArgs struct {
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListIotDpsResourceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotDpsResourceKeysArgs)(nil)).Elem()
}


type ListIotDpsResourceKeysResultOutput struct{ *pulumi.OutputState }

func (ListIotDpsResourceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotDpsResourceKeysResult)(nil)).Elem()
}

func (o ListIotDpsResourceKeysResultOutput) ToListIotDpsResourceKeysResultOutput() ListIotDpsResourceKeysResultOutput {
	return o
}

func (o ListIotDpsResourceKeysResultOutput) ToListIotDpsResourceKeysResultOutputWithContext(ctx context.Context) ListIotDpsResourceKeysResultOutput {
	return o
}

func (o ListIotDpsResourceKeysResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListIotDpsResourceKeysResultOutput) Value() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysResult) []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		return v.Value
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotDpsResourceKeysResultOutput{})
}
