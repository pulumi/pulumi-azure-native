


package v20180301beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSaasSubscriptionLevelAccessToken(ctx *pulumi.Context, args *ListSaasSubscriptionLevelAccessTokenArgs, opts ...pulumi.InvokeOption) (*ListSaasSubscriptionLevelAccessTokenResult, error) {
	var rv ListSaasSubscriptionLevelAccessTokenResult
	err := ctx.Invoke("azure-native:saas/v20180301beta:listSaasSubscriptionLevelAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSaasSubscriptionLevelAccessTokenArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListSaasSubscriptionLevelAccessTokenResult struct {
	PublisherOfferBaseUri *string `pulumi:"publisherOfferBaseUri"`
	Token                 *string `pulumi:"token"`
}

func ListSaasSubscriptionLevelAccessTokenOutput(ctx *pulumi.Context, args ListSaasSubscriptionLevelAccessTokenOutputArgs, opts ...pulumi.InvokeOption) ListSaasSubscriptionLevelAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSaasSubscriptionLevelAccessTokenResult, error) {
			args := v.(ListSaasSubscriptionLevelAccessTokenArgs)
			r, err := ListSaasSubscriptionLevelAccessToken(ctx, &args, opts...)
			var s ListSaasSubscriptionLevelAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSaasSubscriptionLevelAccessTokenResultOutput)
}

type ListSaasSubscriptionLevelAccessTokenOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListSaasSubscriptionLevelAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSaasSubscriptionLevelAccessTokenArgs)(nil)).Elem()
}


type ListSaasSubscriptionLevelAccessTokenResultOutput struct{ *pulumi.OutputState }

func (ListSaasSubscriptionLevelAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSaasSubscriptionLevelAccessTokenResult)(nil)).Elem()
}

func (o ListSaasSubscriptionLevelAccessTokenResultOutput) ToListSaasSubscriptionLevelAccessTokenResultOutput() ListSaasSubscriptionLevelAccessTokenResultOutput {
	return o
}

func (o ListSaasSubscriptionLevelAccessTokenResultOutput) ToListSaasSubscriptionLevelAccessTokenResultOutputWithContext(ctx context.Context) ListSaasSubscriptionLevelAccessTokenResultOutput {
	return o
}

func (o ListSaasSubscriptionLevelAccessTokenResultOutput) PublisherOfferBaseUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSaasSubscriptionLevelAccessTokenResult) *string { return v.PublisherOfferBaseUri }).(pulumi.StringPtrOutput)
}

func (o ListSaasSubscriptionLevelAccessTokenResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSaasSubscriptionLevelAccessTokenResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSaasSubscriptionLevelAccessTokenResultOutput{})
}
