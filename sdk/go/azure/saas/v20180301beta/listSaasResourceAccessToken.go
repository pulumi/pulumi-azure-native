


package v20180301beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSaasResourceAccessToken(ctx *pulumi.Context, args *ListSaasResourceAccessTokenArgs, opts ...pulumi.InvokeOption) (*ListSaasResourceAccessTokenResult, error) {
	var rv ListSaasResourceAccessTokenResult
	err := ctx.Invoke("azure-native:saas/v20180301beta:listSaasResourceAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSaasResourceAccessTokenArgs struct {
	ResourceId string `pulumi:"resourceId"`
}


type ListSaasResourceAccessTokenResult struct {
	PublisherOfferBaseUri *string `pulumi:"publisherOfferBaseUri"`
	Token                 *string `pulumi:"token"`
}

func ListSaasResourceAccessTokenOutput(ctx *pulumi.Context, args ListSaasResourceAccessTokenOutputArgs, opts ...pulumi.InvokeOption) ListSaasResourceAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSaasResourceAccessTokenResult, error) {
			args := v.(ListSaasResourceAccessTokenArgs)
			r, err := ListSaasResourceAccessToken(ctx, &args, opts...)
			var s ListSaasResourceAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSaasResourceAccessTokenResultOutput)
}

type ListSaasResourceAccessTokenOutputArgs struct {
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
}

func (ListSaasResourceAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSaasResourceAccessTokenArgs)(nil)).Elem()
}


type ListSaasResourceAccessTokenResultOutput struct{ *pulumi.OutputState }

func (ListSaasResourceAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSaasResourceAccessTokenResult)(nil)).Elem()
}

func (o ListSaasResourceAccessTokenResultOutput) ToListSaasResourceAccessTokenResultOutput() ListSaasResourceAccessTokenResultOutput {
	return o
}

func (o ListSaasResourceAccessTokenResultOutput) ToListSaasResourceAccessTokenResultOutputWithContext(ctx context.Context) ListSaasResourceAccessTokenResultOutput {
	return o
}

func (o ListSaasResourceAccessTokenResultOutput) PublisherOfferBaseUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSaasResourceAccessTokenResult) *string { return v.PublisherOfferBaseUri }).(pulumi.StringPtrOutput)
}

func (o ListSaasResourceAccessTokenResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSaasResourceAccessTokenResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSaasResourceAccessTokenResultOutput{})
}
