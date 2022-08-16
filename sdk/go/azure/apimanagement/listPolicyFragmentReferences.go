


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPolicyFragmentReferences(ctx *pulumi.Context, args *ListPolicyFragmentReferencesArgs, opts ...pulumi.InvokeOption) (*ListPolicyFragmentReferencesResult, error) {
	var rv ListPolicyFragmentReferencesResult
	err := ctx.Invoke("azure-native:apimanagement:listPolicyFragmentReferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPolicyFragmentReferencesArgs struct {
	Id                string `pulumi:"id"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	Skip              *int   `pulumi:"skip"`
	Top               *int   `pulumi:"top"`
}


type ListPolicyFragmentReferencesResult struct {
	Count    *float64                          `pulumi:"count"`
	NextLink *string                           `pulumi:"nextLink"`
	Value    []ResourceCollectionResponseValue `pulumi:"value"`
}

func ListPolicyFragmentReferencesOutput(ctx *pulumi.Context, args ListPolicyFragmentReferencesOutputArgs, opts ...pulumi.InvokeOption) ListPolicyFragmentReferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPolicyFragmentReferencesResult, error) {
			args := v.(ListPolicyFragmentReferencesArgs)
			r, err := ListPolicyFragmentReferences(ctx, &args, opts...)
			var s ListPolicyFragmentReferencesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPolicyFragmentReferencesResultOutput)
}

type ListPolicyFragmentReferencesOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	Skip              pulumi.IntPtrInput `pulumi:"skip"`
	Top               pulumi.IntPtrInput `pulumi:"top"`
}

func (ListPolicyFragmentReferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicyFragmentReferencesArgs)(nil)).Elem()
}


type ListPolicyFragmentReferencesResultOutput struct{ *pulumi.OutputState }

func (ListPolicyFragmentReferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicyFragmentReferencesResult)(nil)).Elem()
}

func (o ListPolicyFragmentReferencesResultOutput) ToListPolicyFragmentReferencesResultOutput() ListPolicyFragmentReferencesResultOutput {
	return o
}

func (o ListPolicyFragmentReferencesResultOutput) ToListPolicyFragmentReferencesResultOutputWithContext(ctx context.Context) ListPolicyFragmentReferencesResultOutput {
	return o
}

func (o ListPolicyFragmentReferencesResultOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListPolicyFragmentReferencesResult) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

func (o ListPolicyFragmentReferencesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPolicyFragmentReferencesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListPolicyFragmentReferencesResultOutput) Value() ResourceCollectionResponseValueArrayOutput {
	return o.ApplyT(func(v ListPolicyFragmentReferencesResult) []ResourceCollectionResponseValue { return v.Value }).(ResourceCollectionResponseValueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPolicyFragmentReferencesResultOutput{})
}
