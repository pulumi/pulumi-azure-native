


package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	var rv LookupApiResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiResult struct {
	ApiRevision                   *string                                        `pulumi:"apiRevision"`
	ApiRevisionDescription        *string                                        `pulumi:"apiRevisionDescription"`
	ApiType                       *string                                        `pulumi:"apiType"`
	ApiVersion                    *string                                        `pulumi:"apiVersion"`
	ApiVersionDescription         *string                                        `pulumi:"apiVersionDescription"`
	ApiVersionSet                 *ApiVersionSetContractDetailsResponse          `pulumi:"apiVersionSet"`
	ApiVersionSetId               *string                                        `pulumi:"apiVersionSetId"`
	AuthenticationSettings        *AuthenticationSettingsContractResponse        `pulumi:"authenticationSettings"`
	Description                   *string                                        `pulumi:"description"`
	DisplayName                   *string                                        `pulumi:"displayName"`
	Id                            string                                         `pulumi:"id"`
	IsCurrent                     *bool                                          `pulumi:"isCurrent"`
	IsOnline                      bool                                           `pulumi:"isOnline"`
	Name                          string                                         `pulumi:"name"`
	Path                          string                                         `pulumi:"path"`
	Protocols                     []string                                       `pulumi:"protocols"`
	ServiceUrl                    *string                                        `pulumi:"serviceUrl"`
	SourceApiId                   *string                                        `pulumi:"sourceApiId"`
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContractResponse `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          *bool                                          `pulumi:"subscriptionRequired"`
	Type                          string                                         `pulumi:"type"`
}
