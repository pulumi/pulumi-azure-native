


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualHubRouteTableV2 struct {
	pulumi.CustomResourceState

	AttachedConnections pulumi.StringArrayOutput             `pulumi:"attachedConnections"`
	Etag                pulumi.StringOutput                  `pulumi:"etag"`
	Name                pulumi.StringPtrOutput               `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                  `pulumi:"provisioningState"`
	Routes              VirtualHubRouteV2ResponseArrayOutput `pulumi:"routes"`
}


func NewVirtualHubRouteTableV2(ctx *pulumi.Context,
	name string, args *VirtualHubRouteTableV2Args, opts ...pulumi.ResourceOption) (*VirtualHubRouteTableV2, error) {
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
			Type: pulumi.String("azure-native:network:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualHubRouteTableV2"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualHubRouteTableV2"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHubRouteTableV2
	err := ctx.RegisterResource("azure-native:network/v20210201:VirtualHubRouteTableV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualHubRouteTableV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubRouteTableV2State, opts ...pulumi.ResourceOption) (*VirtualHubRouteTableV2, error) {
	var resource VirtualHubRouteTableV2
	err := ctx.ReadResource("azure-native:network/v20210201:VirtualHubRouteTableV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualHubRouteTableV2State struct {
}

type VirtualHubRouteTableV2State struct {
}

func (VirtualHubRouteTableV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubRouteTableV2State)(nil)).Elem()
}

type virtualHubRouteTableV2Args struct {
	AttachedConnections []string            `pulumi:"attachedConnections"`
	Id                  *string             `pulumi:"id"`
	Name                *string             `pulumi:"name"`
	ResourceGroupName   string              `pulumi:"resourceGroupName"`
	RouteTableName      *string             `pulumi:"routeTableName"`
	Routes              []VirtualHubRouteV2 `pulumi:"routes"`
	VirtualHubName      string              `pulumi:"virtualHubName"`
}


type VirtualHubRouteTableV2Args struct {
	AttachedConnections pulumi.StringArrayInput
	Id                  pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	RouteTableName      pulumi.StringPtrInput
	Routes              VirtualHubRouteV2ArrayInput
	VirtualHubName      pulumi.StringInput
}

func (VirtualHubRouteTableV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubRouteTableV2Args)(nil)).Elem()
}

type VirtualHubRouteTableV2Input interface {
	pulumi.Input

	ToVirtualHubRouteTableV2Output() VirtualHubRouteTableV2Output
	ToVirtualHubRouteTableV2OutputWithContext(ctx context.Context) VirtualHubRouteTableV2Output
}

func (*VirtualHubRouteTableV2) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTableV2)(nil)).Elem()
}

func (i *VirtualHubRouteTableV2) ToVirtualHubRouteTableV2Output() VirtualHubRouteTableV2Output {
	return i.ToVirtualHubRouteTableV2OutputWithContext(context.Background())
}

func (i *VirtualHubRouteTableV2) ToVirtualHubRouteTableV2OutputWithContext(ctx context.Context) VirtualHubRouteTableV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableV2Output)
}

type VirtualHubRouteTableV2Output struct{ *pulumi.OutputState }

func (VirtualHubRouteTableV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTableV2)(nil)).Elem()
}

func (o VirtualHubRouteTableV2Output) ToVirtualHubRouteTableV2Output() VirtualHubRouteTableV2Output {
	return o
}

func (o VirtualHubRouteTableV2Output) ToVirtualHubRouteTableV2OutputWithContext(ctx context.Context) VirtualHubRouteTableV2Output {
	return o
}

func (o VirtualHubRouteTableV2Output) AttachedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableV2) pulumi.StringArrayOutput { return v.AttachedConnections }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteTableV2Output) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableV2) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualHubRouteTableV2Output) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableV2) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteTableV2Output) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableV2) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualHubRouteTableV2Output) Routes() VirtualHubRouteV2ResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableV2) VirtualHubRouteV2ResponseArrayOutput { return v.Routes }).(VirtualHubRouteV2ResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubRouteTableV2Output{})
}
