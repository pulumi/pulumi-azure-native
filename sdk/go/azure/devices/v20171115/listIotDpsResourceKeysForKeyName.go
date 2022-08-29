


package v20171115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIotDpsResourceKeysForKeyName(ctx *pulumi.Context, args *ListIotDpsResourceKeysForKeyNameArgs, opts ...pulumi.InvokeOption) (*ListIotDpsResourceKeysForKeyNameResult, error) {
	var rv ListIotDpsResourceKeysForKeyNameResult
	err := ctx.Invoke("azure-native:devices/v20171115:listIotDpsResourceKeysForKeyName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotDpsResourceKeysForKeyNameArgs struct {
	KeyName                 string `pulumi:"keyName"`
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type ListIotDpsResourceKeysForKeyNameResult struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListIotDpsResourceKeysForKeyNameOutput(ctx *pulumi.Context, args ListIotDpsResourceKeysForKeyNameOutputArgs, opts ...pulumi.InvokeOption) ListIotDpsResourceKeysForKeyNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIotDpsResourceKeysForKeyNameResult, error) {
			args := v.(ListIotDpsResourceKeysForKeyNameArgs)
			r, err := ListIotDpsResourceKeysForKeyName(ctx, &args, opts...)
			var s ListIotDpsResourceKeysForKeyNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIotDpsResourceKeysForKeyNameResultOutput)
}

type ListIotDpsResourceKeysForKeyNameOutputArgs struct {
	KeyName                 pulumi.StringInput `pulumi:"keyName"`
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListIotDpsResourceKeysForKeyNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotDpsResourceKeysForKeyNameArgs)(nil)).Elem()
}


type ListIotDpsResourceKeysForKeyNameResultOutput struct{ *pulumi.OutputState }

func (ListIotDpsResourceKeysForKeyNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIotDpsResourceKeysForKeyNameResult)(nil)).Elem()
}

func (o ListIotDpsResourceKeysForKeyNameResultOutput) ToListIotDpsResourceKeysForKeyNameResultOutput() ListIotDpsResourceKeysForKeyNameResultOutput {
	return o
}

func (o ListIotDpsResourceKeysForKeyNameResultOutput) ToListIotDpsResourceKeysForKeyNameResultOutputWithContext(ctx context.Context) ListIotDpsResourceKeysForKeyNameResultOutput {
	return o
}

func (o ListIotDpsResourceKeysForKeyNameResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysForKeyNameResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ListIotDpsResourceKeysForKeyNameResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysForKeyNameResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListIotDpsResourceKeysForKeyNameResultOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysForKeyNameResult) string { return v.Rights }).(pulumi.StringOutput)
}

func (o ListIotDpsResourceKeysForKeyNameResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIotDpsResourceKeysForKeyNameResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIotDpsResourceKeysForKeyNameResultOutput{})
}
