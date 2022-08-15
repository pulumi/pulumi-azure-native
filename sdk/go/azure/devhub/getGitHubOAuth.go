


package devhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGitHubOAuth(ctx *pulumi.Context, args *GetGitHubOAuthArgs, opts ...pulumi.InvokeOption) (*GetGitHubOAuthResult, error) {
	var rv GetGitHubOAuthResult
	err := ctx.Invoke("azure-native:devhub:getGitHubOAuth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGitHubOAuthArgs struct {
	Location    string  `pulumi:"location"`
	RedirectUrl *string `pulumi:"redirectUrl"`
}


type GetGitHubOAuthResult struct {
	AuthURL *string `pulumi:"authURL"`
	Token   *string `pulumi:"token"`
}

func GetGitHubOAuthOutput(ctx *pulumi.Context, args GetGitHubOAuthOutputArgs, opts ...pulumi.InvokeOption) GetGitHubOAuthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGitHubOAuthResult, error) {
			args := v.(GetGitHubOAuthArgs)
			r, err := GetGitHubOAuth(ctx, &args, opts...)
			var s GetGitHubOAuthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGitHubOAuthResultOutput)
}

type GetGitHubOAuthOutputArgs struct {
	Location    pulumi.StringInput    `pulumi:"location"`
	RedirectUrl pulumi.StringPtrInput `pulumi:"redirectUrl"`
}

func (GetGitHubOAuthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitHubOAuthArgs)(nil)).Elem()
}


type GetGitHubOAuthResultOutput struct{ *pulumi.OutputState }

func (GetGitHubOAuthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitHubOAuthResult)(nil)).Elem()
}

func (o GetGitHubOAuthResultOutput) ToGetGitHubOAuthResultOutput() GetGitHubOAuthResultOutput {
	return o
}

func (o GetGitHubOAuthResultOutput) ToGetGitHubOAuthResultOutputWithContext(ctx context.Context) GetGitHubOAuthResultOutput {
	return o
}

func (o GetGitHubOAuthResultOutput) AuthURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGitHubOAuthResult) *string { return v.AuthURL }).(pulumi.StringPtrOutput)
}

func (o GetGitHubOAuthResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGitHubOAuthResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGitHubOAuthResultOutput{})
}
