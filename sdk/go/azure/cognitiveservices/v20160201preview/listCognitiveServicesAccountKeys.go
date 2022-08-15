


package v20160201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCognitiveServicesAccountKeys(ctx *pulumi.Context, args *ListCognitiveServicesAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListCognitiveServicesAccountKeysResult, error) {
	var rv ListCognitiveServicesAccountKeysResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20160201preview:listCognitiveServicesAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCognitiveServicesAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListCognitiveServicesAccountKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}

func ListCognitiveServicesAccountKeysOutput(ctx *pulumi.Context, args ListCognitiveServicesAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListCognitiveServicesAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCognitiveServicesAccountKeysResult, error) {
			args := v.(ListCognitiveServicesAccountKeysArgs)
			r, err := ListCognitiveServicesAccountKeys(ctx, &args, opts...)
			var s ListCognitiveServicesAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCognitiveServicesAccountKeysResultOutput)
}

type ListCognitiveServicesAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListCognitiveServicesAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCognitiveServicesAccountKeysArgs)(nil)).Elem()
}


type ListCognitiveServicesAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListCognitiveServicesAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCognitiveServicesAccountKeysResult)(nil)).Elem()
}

func (o ListCognitiveServicesAccountKeysResultOutput) ToListCognitiveServicesAccountKeysResultOutput() ListCognitiveServicesAccountKeysResultOutput {
	return o
}

func (o ListCognitiveServicesAccountKeysResultOutput) ToListCognitiveServicesAccountKeysResultOutputWithContext(ctx context.Context) ListCognitiveServicesAccountKeysResultOutput {
	return o
}

func (o ListCognitiveServicesAccountKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCognitiveServicesAccountKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

func (o ListCognitiveServicesAccountKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCognitiveServicesAccountKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCognitiveServicesAccountKeysResultOutput{})
}
