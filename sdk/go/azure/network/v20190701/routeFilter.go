


package v20190701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouteFilter struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput                           `pulumi:"etag"`
	Ipv6Peerings      ExpressRouteCircuitPeeringResponseArrayOutput `pulumi:"ipv6Peerings"`
	Location          pulumi.StringOutput                           `pulumi:"location"`
	Name              pulumi.StringOutput                           `pulumi:"name"`
	Peerings          ExpressRouteCircuitPeeringResponseArrayOutput `pulumi:"peerings"`
	ProvisioningState pulumi.StringOutput                           `pulumi:"provisioningState"`
	Rules             RouteFilterRuleResponseArrayOutput            `pulumi:"rules"`
	Tags              pulumi.StringMapOutput                        `pulumi:"tags"`
	Type              pulumi.StringOutput                           `pulumi:"type"`
}


func NewRouteFilter(ctx *pulumi.Context,
	name string, args *RouteFilterArgs, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:RouteFilter"),
		},
	})
	opts = append(opts, aliases)
	var resource RouteFilter
	err := ctx.RegisterResource("azure-native:network/v20190701:RouteFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRouteFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteFilterState, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	var resource RouteFilter
	err := ctx.ReadResource("azure-native:network/v20190701:RouteFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type routeFilterState struct {
}

type RouteFilterState struct {
}

func (RouteFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterState)(nil)).Elem()
}

type routeFilterArgs struct {
	Id                *string                          `pulumi:"id"`
	Ipv6Peerings      []ExpressRouteCircuitPeeringType `pulumi:"ipv6Peerings"`
	Location          *string                          `pulumi:"location"`
	Peerings          []ExpressRouteCircuitPeeringType `pulumi:"peerings"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	RouteFilterName   *string                          `pulumi:"routeFilterName"`
	Rules             []RouteFilterRuleType            `pulumi:"rules"`
	Tags              map[string]string                `pulumi:"tags"`
}


type RouteFilterArgs struct {
	Id                pulumi.StringPtrInput
	Ipv6Peerings      ExpressRouteCircuitPeeringTypeArrayInput
	Location          pulumi.StringPtrInput
	Peerings          ExpressRouteCircuitPeeringTypeArrayInput
	ResourceGroupName pulumi.StringInput
	RouteFilterName   pulumi.StringPtrInput
	Rules             RouteFilterRuleTypeArrayInput
	Tags              pulumi.StringMapInput
}

func (RouteFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterArgs)(nil)).Elem()
}

type RouteFilterInput interface {
	pulumi.Input

	ToRouteFilterOutput() RouteFilterOutput
	ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput
}

func (*RouteFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilter)(nil)).Elem()
}

func (i *RouteFilter) ToRouteFilterOutput() RouteFilterOutput {
	return i.ToRouteFilterOutputWithContext(context.Background())
}

func (i *RouteFilter) ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterOutput)
}

type RouteFilterOutput struct{ *pulumi.OutputState }

func (RouteFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilter)(nil)).Elem()
}

func (o RouteFilterOutput) ToRouteFilterOutput() RouteFilterOutput {
	return o
}

func (o RouteFilterOutput) ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput {
	return o
}

func (o RouteFilterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteFilterOutput) Ipv6Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilter) ExpressRouteCircuitPeeringResponseArrayOutput { return v.Ipv6Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o RouteFilterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RouteFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouteFilterOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilter) ExpressRouteCircuitPeeringResponseArrayOutput { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o RouteFilterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteFilterOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilter) RouteFilterRuleResponseArrayOutput { return v.Rules }).(RouteFilterRuleResponseArrayOutput)
}

func (o RouteFilterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RouteFilterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteFilterOutput{})
}
