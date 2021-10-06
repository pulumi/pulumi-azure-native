


package datafactory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFactoryGitHubAccessToken(ctx *pulumi.Context, args *GetFactoryGitHubAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetFactoryGitHubAccessTokenResult, error) {
	var rv GetFactoryGitHubAccessTokenResult
	err := ctx.Invoke("azure-native:datafactory:getFactoryGitHubAccessToken", args, &rv, opts...)
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
