


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFactoryGitHubAccessToken(ctx *pulumi.Context, args *GetFactoryGitHubAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetFactoryGitHubAccessTokenResult, error) {
	var rv GetFactoryGitHubAccessTokenResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getFactoryGitHubAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFactoryGitHubAccessTokenArgs struct {
	FactoryName              string              `pulumi:"factoryName"`
	GitHubAccessCode         string              `pulumi:"gitHubAccessCode"`
	GitHubAccessTokenBaseUrl string              `pulumi:"gitHubAccessTokenBaseUrl"`
	GitHubClientId           *string             `pulumi:"gitHubClientId"`
	GitHubClientSecret       *GitHubClientSecret `pulumi:"gitHubClientSecret"`
	ResourceGroupName        string              `pulumi:"resourceGroupName"`
}


type GetFactoryGitHubAccessTokenResult struct {
	GitHubAccessToken *string `pulumi:"gitHubAccessToken"`
}

func GetFactoryGitHubAccessTokenOutput(ctx *pulumi.Context, args GetFactoryGitHubAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetFactoryGitHubAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFactoryGitHubAccessTokenResult, error) {
			args := v.(GetFactoryGitHubAccessTokenArgs)
			r, err := GetFactoryGitHubAccessToken(ctx, &args, opts...)
			var s GetFactoryGitHubAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFactoryGitHubAccessTokenResultOutput)
}

type GetFactoryGitHubAccessTokenOutputArgs struct {
	FactoryName              pulumi.StringInput         `pulumi:"factoryName"`
	GitHubAccessCode         pulumi.StringInput         `pulumi:"gitHubAccessCode"`
	GitHubAccessTokenBaseUrl pulumi.StringInput         `pulumi:"gitHubAccessTokenBaseUrl"`
	GitHubClientId           pulumi.StringPtrInput      `pulumi:"gitHubClientId"`
	GitHubClientSecret       GitHubClientSecretPtrInput `pulumi:"gitHubClientSecret"`
	ResourceGroupName        pulumi.StringInput         `pulumi:"resourceGroupName"`
}

func (GetFactoryGitHubAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFactoryGitHubAccessTokenArgs)(nil)).Elem()
}


type GetFactoryGitHubAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetFactoryGitHubAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFactoryGitHubAccessTokenResult)(nil)).Elem()
}

func (o GetFactoryGitHubAccessTokenResultOutput) ToGetFactoryGitHubAccessTokenResultOutput() GetFactoryGitHubAccessTokenResultOutput {
	return o
}

func (o GetFactoryGitHubAccessTokenResultOutput) ToGetFactoryGitHubAccessTokenResultOutputWithContext(ctx context.Context) GetFactoryGitHubAccessTokenResultOutput {
	return o
}

func (o GetFactoryGitHubAccessTokenResultOutput) GitHubAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFactoryGitHubAccessTokenResult) *string { return v.GitHubAccessToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFactoryGitHubAccessTokenResultOutput{})
}
