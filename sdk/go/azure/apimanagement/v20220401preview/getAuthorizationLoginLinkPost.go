


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAuthorizationLoginLinkPost(ctx *pulumi.Context, args *GetAuthorizationLoginLinkPostArgs, opts ...pulumi.InvokeOption) (*GetAuthorizationLoginLinkPostResult, error) {
	var rv GetAuthorizationLoginLinkPostResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getAuthorizationLoginLinkPost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAuthorizationLoginLinkPostArgs struct {
	AuthorizationId         string  `pulumi:"authorizationId"`
	AuthorizationProviderId string  `pulumi:"authorizationProviderId"`
	PostLoginRedirectUrl    *string `pulumi:"postLoginRedirectUrl"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ServiceName             string  `pulumi:"serviceName"`
}


type GetAuthorizationLoginLinkPostResult struct {
	LoginLink *string `pulumi:"loginLink"`
}

func GetAuthorizationLoginLinkPostOutput(ctx *pulumi.Context, args GetAuthorizationLoginLinkPostOutputArgs, opts ...pulumi.InvokeOption) GetAuthorizationLoginLinkPostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAuthorizationLoginLinkPostResult, error) {
			args := v.(GetAuthorizationLoginLinkPostArgs)
			r, err := GetAuthorizationLoginLinkPost(ctx, &args, opts...)
			var s GetAuthorizationLoginLinkPostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAuthorizationLoginLinkPostResultOutput)
}

type GetAuthorizationLoginLinkPostOutputArgs struct {
	AuthorizationId         pulumi.StringInput    `pulumi:"authorizationId"`
	AuthorizationProviderId pulumi.StringInput    `pulumi:"authorizationProviderId"`
	PostLoginRedirectUrl    pulumi.StringPtrInput `pulumi:"postLoginRedirectUrl"`
	ResourceGroupName       pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName             pulumi.StringInput    `pulumi:"serviceName"`
}

func (GetAuthorizationLoginLinkPostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthorizationLoginLinkPostArgs)(nil)).Elem()
}


type GetAuthorizationLoginLinkPostResultOutput struct{ *pulumi.OutputState }

func (GetAuthorizationLoginLinkPostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthorizationLoginLinkPostResult)(nil)).Elem()
}

func (o GetAuthorizationLoginLinkPostResultOutput) ToGetAuthorizationLoginLinkPostResultOutput() GetAuthorizationLoginLinkPostResultOutput {
	return o
}

func (o GetAuthorizationLoginLinkPostResultOutput) ToGetAuthorizationLoginLinkPostResultOutputWithContext(ctx context.Context) GetAuthorizationLoginLinkPostResultOutput {
	return o
}

func (o GetAuthorizationLoginLinkPostResultOutput) LoginLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAuthorizationLoginLinkPostResult) *string { return v.LoginLink }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAuthorizationLoginLinkPostResultOutput{})
}
