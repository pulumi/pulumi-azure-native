


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouteMap struct {
	pulumi.CustomResourceState

	AssociatedInboundConnections  pulumi.StringArrayOutput        `pulumi:"associatedInboundConnections"`
	AssociatedOutboundConnections pulumi.StringArrayOutput        `pulumi:"associatedOutboundConnections"`
	Etag                          pulumi.StringOutput             `pulumi:"etag"`
	Name                          pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState             pulumi.StringOutput             `pulumi:"provisioningState"`
	Rules                         RouteMapRuleResponseArrayOutput `pulumi:"rules"`
	Type                          pulumi.StringOutput             `pulumi:"type"`
}


func NewRouteMap(ctx *pulumi.Context,
	name string, args *RouteMapArgs, opts ...pulumi.ResourceOption) (*RouteMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20220701:RouteMap"),
		},
	})
	opts = append(opts, aliases)
	var resource RouteMap
	err := ctx.RegisterResource("azure-native:network/v20220501:RouteMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRouteMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteMapState, opts ...pulumi.ResourceOption) (*RouteMap, error) {
	var resource RouteMap
	err := ctx.ReadResource("azure-native:network/v20220501:RouteMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type routeMapState struct {
}

type RouteMapState struct {
}

func (RouteMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeMapState)(nil)).Elem()
}

type routeMapArgs struct {
	AssociatedInboundConnections  []string       `pulumi:"associatedInboundConnections"`
	AssociatedOutboundConnections []string       `pulumi:"associatedOutboundConnections"`
	Id                            *string        `pulumi:"id"`
	ResourceGroupName             string         `pulumi:"resourceGroupName"`
	RouteMapName                  *string        `pulumi:"routeMapName"`
	Rules                         []RouteMapRule `pulumi:"rules"`
	VirtualHubName                string         `pulumi:"virtualHubName"`
}


type RouteMapArgs struct {
	AssociatedInboundConnections  pulumi.StringArrayInput
	AssociatedOutboundConnections pulumi.StringArrayInput
	Id                            pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	RouteMapName                  pulumi.StringPtrInput
	Rules                         RouteMapRuleArrayInput
	VirtualHubName                pulumi.StringInput
}

func (RouteMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeMapArgs)(nil)).Elem()
}

type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput
}

func (*RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteMap)(nil)).Elem()
}

func (i *RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i *RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteMap)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) AssociatedInboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringArrayOutput { return v.AssociatedInboundConnections }).(pulumi.StringArrayOutput)
}

func (o RouteMapOutput) AssociatedOutboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringArrayOutput { return v.AssociatedOutboundConnections }).(pulumi.StringArrayOutput)
}

func (o RouteMapOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouteMapOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteMapOutput) Rules() RouteMapRuleResponseArrayOutput {
	return o.ApplyT(func(v *RouteMap) RouteMapRuleResponseArrayOutput { return v.Rules }).(RouteMapRuleResponseArrayOutput)
}

func (o RouteMapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteMapOutput{})
}
