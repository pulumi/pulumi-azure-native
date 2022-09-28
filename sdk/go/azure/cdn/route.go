


package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Route struct {
	pulumi.CustomResourceState

	CompressionSettings        CompressionSettingsResponsePtrOutput `pulumi:"compressionSettings"`
	CustomDomains              ResourceReferenceResponseArrayOutput `pulumi:"customDomains"`
	DeploymentStatus           pulumi.StringOutput                  `pulumi:"deploymentStatus"`
	EnabledState               pulumi.StringPtrOutput               `pulumi:"enabledState"`
	ForwardingProtocol         pulumi.StringPtrOutput               `pulumi:"forwardingProtocol"`
	HttpsRedirect              pulumi.StringPtrOutput               `pulumi:"httpsRedirect"`
	LinkToDefaultDomain        pulumi.StringPtrOutput               `pulumi:"linkToDefaultDomain"`
	Name                       pulumi.StringOutput                  `pulumi:"name"`
	OriginGroup                ResourceReferenceResponseOutput      `pulumi:"originGroup"`
	OriginPath                 pulumi.StringPtrOutput               `pulumi:"originPath"`
	PatternsToMatch            pulumi.StringArrayOutput             `pulumi:"patternsToMatch"`
	ProvisioningState          pulumi.StringOutput                  `pulumi:"provisioningState"`
	QueryStringCachingBehavior pulumi.StringPtrOutput               `pulumi:"queryStringCachingBehavior"`
	RuleSets                   ResourceReferenceResponseArrayOutput `pulumi:"ruleSets"`
	SupportedProtocols         pulumi.StringArrayOutput             `pulumi:"supportedProtocols"`
	SystemData                 SystemDataResponseOutput             `pulumi:"systemData"`
	Type                       pulumi.StringOutput                  `pulumi:"type"`
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
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Route"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Route"),
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
	CompressionSettings        *CompressionSettings           `pulumi:"compressionSettings"`
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
	CompressionSettings        CompressionSettingsPtrInput
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

func (o RouteOutput) CompressionSettings() CompressionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Route) CompressionSettingsResponsePtrOutput { return v.CompressionSettings }).(CompressionSettingsResponsePtrOutput)
}

func (o RouteOutput) CustomDomains() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *Route) ResourceReferenceResponseArrayOutput { return v.CustomDomains }).(ResourceReferenceResponseArrayOutput)
}

func (o RouteOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o RouteOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) ForwardingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.ForwardingProtocol }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) HttpsRedirect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.HttpsRedirect }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) LinkToDefaultDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.LinkToDefaultDomain }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouteOutput) OriginGroup() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *Route) ResourceReferenceResponseOutput { return v.OriginGroup }).(ResourceReferenceResponseOutput)
}

func (o RouteOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.OriginPath }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Route) pulumi.StringArrayOutput { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

func (o RouteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteOutput) QueryStringCachingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.QueryStringCachingBehavior }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) RuleSets() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *Route) ResourceReferenceResponseArrayOutput { return v.RuleSets }).(ResourceReferenceResponseArrayOutput)
}

func (o RouteOutput) SupportedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Route) pulumi.StringArrayOutput { return v.SupportedProtocols }).(pulumi.StringArrayOutput)
}

func (o RouteOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Route) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RouteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
