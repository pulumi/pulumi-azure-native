


package v20220430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotHubResourceKeysForKeyName(ctx *pulumi.Context, args *ListIotHubResourceKeysForKeyNameArgs, opts ...pulumi.InvokeOption) (*ListIotHubResourceKeysForKeyNameResult, error) {
	var rv ListIotHubResourceKeysForKeyNameResult
	err := ctx.Invoke("azure-native:devices/v20220430preview:listIotHubResourceKeysForKeyName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotHubResourceKeysForKeyNameArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListIotHubResourceKeysForKeyNameResult struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListIotHubResourceKeysForKeyNameOutput(ctx *pulumi.Context, args ListIotHubResourceKeysForKeyNameOutputArgs, opts ...pulumi.InvokeOption) ListIotHubResourceKeysForKeyNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotHubResourceKeysForKeyNameResult, error) {
			args := v.(ListIotHubResourceKeysForKeyNameArgs)
			r, err := ListIotHubResourceKeysForKeyName(ctx, &args, opts...)
			var s ListIotHubResourceKeysForKeyNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotHubResourceKeysForKeyNameResultOutput)
}

type ListIotHubResourceKeysForKeyNameOutputArgs struct {
	KeyName           pulumi.StringInput `pulumi:"keyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListIotHubResourceKeysForKeyNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysForKeyNameArgs)(nil)).Elem()
}


type ListIotHubResourceKeysForKeyNameResultOutput struct{ *pulumi.OutputState }

func (ListIotHubResourceKeysForKeyNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotHubResourceKeysForKeyNameResult)(nil)).Elem()
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) ToListIotHubResourceKeysForKeyNameResultOutput() ListIotHubResourceKeysForKeyNameResultOutput {
	return o
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) ToListIotHubResourceKeysForKeyNameResultOutputWithContext(ctx context.Context) ListIotHubResourceKeysForKeyNameResultOutput {
	return o
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) string { return v.Rights }).(pulumi.StringOutput)
}

func (o ListIotHubResourceKeysForKeyNameResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIotHubResourceKeysForKeyNameResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotHubResourceKeysForKeyNameResultOutput{})
}
