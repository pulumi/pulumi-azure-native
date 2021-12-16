


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HubRouteTable struct {
	pulumi.CustomResourceState

	AssociatedConnections  pulumi.StringArrayOutput    `pulumi:"associatedConnections"`
	Etag                   pulumi.StringOutput         `pulumi:"etag"`
	Labels                 pulumi.StringArrayOutput    `pulumi:"labels"`
	Name                   pulumi.StringPtrOutput      `pulumi:"name"`
	PropagatingConnections pulumi.StringArrayOutput    `pulumi:"propagatingConnections"`
	ProvisioningState      pulumi.StringOutput         `pulumi:"provisioningState"`
	Routes                 HubRouteResponseArrayOutput `pulumi:"routes"`
	Type                   pulumi.StringOutput         `pulumi:"type"`
}


func NewHubRouteTable(ctx *pulumi.Context,
	name string, args *HubRouteTableArgs, opts ...pulumi.ResourceOption) (*HubRouteTable, error) {
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
			Type: pulumi.String("azure-native:network:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:HubRouteTable"),
		},
	})
	opts = append(opts, aliases)
	var resource HubRouteTable
	err := ctx.RegisterResource("azure-native:network/v20210201:HubRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHubRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubRouteTableState, opts ...pulumi.ResourceOption) (*HubRouteTable, error) {
	var resource HubRouteTable
	err := ctx.ReadResource("azure-native:network/v20210201:HubRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hubRouteTableState struct {
}

type HubRouteTableState struct {
}

func (HubRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubRouteTableState)(nil)).Elem()
}

type hubRouteTableArgs struct {
	Id                *string    `pulumi:"id"`
	Labels            []string   `pulumi:"labels"`
	Name              *string    `pulumi:"name"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
	RouteTableName    *string    `pulumi:"routeTableName"`
	Routes            []HubRoute `pulumi:"routes"`
	VirtualHubName    string     `pulumi:"virtualHubName"`
}


type HubRouteTableArgs struct {
	Id                pulumi.StringPtrInput
	Labels            pulumi.StringArrayInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RouteTableName    pulumi.StringPtrInput
	Routes            HubRouteArrayInput
	VirtualHubName    pulumi.StringInput
}

func (HubRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubRouteTableArgs)(nil)).Elem()
}

type HubRouteTableInput interface {
	pulumi.Input

	ToHubRouteTableOutput() HubRouteTableOutput
	ToHubRouteTableOutputWithContext(ctx context.Context) HubRouteTableOutput
}

func (*HubRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**HubRouteTable)(nil)).Elem()
}

func (i *HubRouteTable) ToHubRouteTableOutput() HubRouteTableOutput {
	return i.ToHubRouteTableOutputWithContext(context.Background())
}

func (i *HubRouteTable) ToHubRouteTableOutputWithContext(ctx context.Context) HubRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubRouteTableOutput)
}

type HubRouteTableOutput struct{ *pulumi.OutputState }

func (HubRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HubRouteTable)(nil)).Elem()
}

func (o HubRouteTableOutput) ToHubRouteTableOutput() HubRouteTableOutput {
	return o
}

func (o HubRouteTableOutput) ToHubRouteTableOutputWithContext(ctx context.Context) HubRouteTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HubRouteTableOutput{})
}
