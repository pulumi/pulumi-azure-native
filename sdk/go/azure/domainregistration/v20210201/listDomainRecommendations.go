


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDomainRecommendations(ctx *pulumi.Context, args *ListDomainRecommendationsArgs, opts ...pulumi.InvokeOption) (*ListDomainRecommendationsResult, error) {
	var rv ListDomainRecommendationsResult
	err := ctx.Invoke("azure-native:domainregistration/v20210201:listDomainRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainRecommendationsArgs struct {
	Keywords                 *string `pulumi:"keywords"`
	MaxDomainRecommendations *int    `pulumi:"maxDomainRecommendations"`
}


type ListDomainRecommendationsResult struct {
	NextLink string                   `pulumi:"nextLink"`
	Value    []NameIdentifierResponse `pulumi:"value"`
}

func ListDomainRecommendationsOutput(ctx *pulumi.Context, args ListDomainRecommendationsOutputArgs, opts ...pulumi.InvokeOption) ListDomainRecommendationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDomainRecommendationsResult, error) {
			args := v.(ListDomainRecommendationsArgs)
			r, err := ListDomainRecommendations(ctx, &args, opts...)
			var s ListDomainRecommendationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDomainRecommendationsResultOutput)
}

type ListDomainRecommendationsOutputArgs struct {
	Keywords                 pulumi.StringPtrInput `pulumi:"keywords"`
	MaxDomainRecommendations pulumi.IntPtrInput    `pulumi:"maxDomainRecommendations"`
}

func (ListDomainRecommendationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainRecommendationsArgs)(nil)).Elem()
}


type ListDomainRecommendationsResultOutput struct{ *pulumi.OutputState }

func (ListDomainRecommendationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainRecommendationsResult)(nil)).Elem()
}

func (o ListDomainRecommendationsResultOutput) ToListDomainRecommendationsResultOutput() ListDomainRecommendationsResultOutput {
	return o
}

func (o ListDomainRecommendationsResultOutput) ToListDomainRecommendationsResultOutputWithContext(ctx context.Context) ListDomainRecommendationsResultOutput {
	return o
}

func (o ListDomainRecommendationsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListDomainRecommendationsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListDomainRecommendationsResultOutput) Value() NameIdentifierResponseArrayOutput {
	return o.ApplyT(func(v ListDomainRecommendationsResult) []NameIdentifierResponse { return v.Value }).(NameIdentifierResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDomainRecommendationsResultOutput{})
}
