


package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Route struct {
	pulumi.CustomResourceState

	CompressionSettings        CompressionSettingsResponseArrayOutput `pulumi:"compressionSettings"`
	CustomDomains              ResourceReferenceResponseArrayOutput   `pulumi:"customDomains"`
	DeploymentStatus           pulumi.StringOutput                    `pulumi:"deploymentStatus"`
	EnabledState               pulumi.StringPtrOutput                 `pulumi:"enabledState"`
	ForwardingProtocol         pulumi.StringPtrOutput                 `pulumi:"forwardingProtocol"`
	HttpsRedirect              pulumi.StringPtrOutput                 `pulumi:"httpsRedirect"`
	LinkToDefaultDomain        pulumi.StringPtrOutput                 `pulumi:"linkToDefaultDomain"`
	Name                       pulumi.StringOutput                    `pulumi:"name"`
	OriginGroup                ResourceReferenceResponseOutput        `pulumi:"originGroup"`
	OriginPath                 pulumi.StringPtrOutput                 `pulumi:"originPath"`
	PatternsToMatch            pulumi.StringArrayOutput               `pulumi:"patternsToMatch"`
	ProvisioningState          pulumi.StringOutput                    `pulumi:"provisioningState"`
	QueryStringCachingBehavior pulumi.StringPtrOutput                 `pulumi:"queryStringCachingBehavior"`
	RuleSets                   ResourceReferenceResponseArrayOutput   `pulumi:"ruleSets"`
	SupportedProtocols         pulumi.StringArrayOutput               `pulumi:"supportedProtocols"`
	SystemData                 SystemDataResponseOutput               `pulumi:"systemData"`
	Type                       pulumi.StringOutput                    `pulumi:"type"`
}


func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.OriginGroup == nil {
		return nil, errors.New("invalid value for required argument 'OriginGroup'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Route"),
		},
	})
	opts = append(opts, aliases)
	var resource Route
	err := ctx.RegisterResource("azure-native:cdn:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azure-native:cdn:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	CompressionSettings        []CompressionSettings          `pulumi:"compressionSettings"`
	CustomDomains              []ResourceReference            `pulumi:"customDomains"`
	EnabledState               *string                        `pulumi:"enabledState"`
	EndpointName               string                         `pulumi:"endpointName"`
	ForwardingProtocol         *string                        `pulumi:"forwardingProtocol"`
	HttpsRedirect              *string                        `pulumi:"httpsRedirect"`
	LinkToDefaultDomain        *string                        `pulumi:"linkToDefaultDomain"`
	OriginGroup                ResourceReference              `pulumi:"originGroup"`
	OriginPath                 *string                        `pulumi:"originPath"`
	PatternsToMatch            []string                       `pulumi:"patternsToMatch"`
	ProfileName                string                         `pulumi:"profileName"`
	QueryStringCachingBehavior *AfdQueryStringCachingBehavior `pulumi:"queryStringCachingBehavior"`
	ResourceGroupName          string                         `pulumi:"resourceGroupName"`
	RouteName                  *string                        `pulumi:"routeName"`
	RuleSets                   []ResourceReference            `pulumi:"ruleSets"`
	SupportedProtocols         []string                       `pulumi:"supportedProtocols"`
}


type RouteArgs struct {
	CompressionSettings        CompressionSettingsArrayInput
	CustomDomains              ResourceReferenceArrayInput
	EnabledState               pulumi.StringPtrInput
	EndpointName               pulumi.StringInput
	ForwardingProtocol         pulumi.StringPtrInput
	HttpsRedirect              pulumi.StringPtrInput
	LinkToDefaultDomain        pulumi.StringPtrInput
	OriginGroup                ResourceReferenceInput
	OriginPath                 pulumi.StringPtrInput
	PatternsToMatch            pulumi.StringArrayInput
	ProfileName                pulumi.StringInput
	QueryStringCachingBehavior AfdQueryStringCachingBehaviorPtrInput
	ResourceGroupName          pulumi.StringInput
	RouteName                  pulumi.StringPtrInput
	RuleSets                   ResourceReferenceArrayInput
	SupportedProtocols         pulumi.StringArrayInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
