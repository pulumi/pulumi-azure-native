


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Route struct {
	pulumi.CustomResourceState

	AddressPrefix     pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	Etag              pulumi.StringOutput    `pulumi:"etag"`
	HasBgpOverride    pulumi.BoolPtrOutput   `pulumi:"hasBgpOverride"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	NextHopIpAddress  pulumi.StringPtrOutput `pulumi:"nextHopIpAddress"`
	NextHopType       pulumi.StringOutput    `pulumi:"nextHopType"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Type              pulumi.StringPtrOutput `pulumi:"type"`
}


func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NextHopType == nil {
		return nil, errors.New("invalid value for required argument 'NextHopType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RouteTableName == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:Route"),
		},
	})
	opts = append(opts, aliases)
	var resource Route
	err := ctx.RegisterResource("azure-native:network/v20210301:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azure-native:network/v20210301:Route", name, id, state, &resource, opts...)
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
	AddressPrefix     *string `pulumi:"addressPrefix"`
	HasBgpOverride    *bool   `pulumi:"hasBgpOverride"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RouteName         *string `pulumi:"routeName"`
	RouteTableName    string  `pulumi:"routeTableName"`
	Type              *string `pulumi:"type"`
}


type RouteArgs struct {
	AddressPrefix     pulumi.StringPtrInput
	HasBgpOverride    pulumi.BoolPtrInput
	Id                pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	NextHopIpAddress  pulumi.StringPtrInput
	NextHopType       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RouteName         pulumi.StringPtrInput
	RouteTableName    pulumi.StringInput
	Type              pulumi.StringPtrInput
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

func (o RouteOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteOutput) HasBgpOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.BoolPtrOutput { return v.HasBgpOverride }).(pulumi.BoolPtrOutput)
}

func (o RouteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

func (o RouteOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.NextHopType }).(pulumi.StringOutput)
}

func (o RouteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
