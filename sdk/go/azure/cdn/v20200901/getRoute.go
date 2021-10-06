


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	var rv LookupRouteResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteName         string `pulumi:"routeName"`
}


type LookupRouteResult struct {
	CompressionSettings        []CompressionSettingsResponse `pulumi:"compressionSettings"`
	CustomDomains              []ResourceReferenceResponse   `pulumi:"customDomains"`
	DeploymentStatus           string                        `pulumi:"deploymentStatus"`
	EnabledState               *string                       `pulumi:"enabledState"`
	ForwardingProtocol         *string                       `pulumi:"forwardingProtocol"`
	HttpsRedirect              *string                       `pulumi:"httpsRedirect"`
	Id                         string                        `pulumi:"id"`
	LinkToDefaultDomain        *string                       `pulumi:"linkToDefaultDomain"`
	Name                       string                        `pulumi:"name"`
	OriginGroup                ResourceReferenceResponse     `pulumi:"originGroup"`
	OriginPath                 *string                       `pulumi:"originPath"`
	PatternsToMatch            []string                      `pulumi:"patternsToMatch"`
	ProvisioningState          string                        `pulumi:"provisioningState"`
	QueryStringCachingBehavior *string                       `pulumi:"queryStringCachingBehavior"`
	RuleSets                   []ResourceReferenceResponse   `pulumi:"ruleSets"`
	SupportedProtocols         []string                      `pulumi:"supportedProtocols"`
	SystemData                 SystemDataResponse            `pulumi:"systemData"`
	Type                       string                        `pulumi:"type"`
}
