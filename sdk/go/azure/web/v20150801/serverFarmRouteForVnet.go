


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerFarmRouteForVnet struct {
	pulumi.CustomResourceState

	EndAddress   pulumi.StringPtrOutput `pulumi:"endAddress"`
	Kind         pulumi.StringPtrOutput `pulumi:"kind"`
	Location     pulumi.StringOutput    `pulumi:"location"`
	Name         pulumi.StringPtrOutput `pulumi:"name"`
	RouteType    pulumi.StringPtrOutput `pulumi:"routeType"`
	StartAddress pulumi.StringPtrOutput `pulumi:"startAddress"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
	Type         pulumi.StringPtrOutput `pulumi:"type"`
}


func NewServerFarmRouteForVnet(ctx *pulumi.Context,
	name string, args *ServerFarmRouteForVnetArgs, opts ...pulumi.ResourceOption) (*ServerFarmRouteForVnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VnetName == nil {
		return nil, errors.New("invalid value for required argument 'VnetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160901:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:ServerFarmRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:ServerFarmRouteForVnet"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerFarmRouteForVnet
	err := ctx.RegisterResource("azure-native:web/v20150801:ServerFarmRouteForVnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerFarmRouteForVnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerFarmRouteForVnetState, opts ...pulumi.ResourceOption) (*ServerFarmRouteForVnet, error) {
	var resource ServerFarmRouteForVnet
	err := ctx.ReadResource("azure-native:web/v20150801:ServerFarmRouteForVnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverFarmRouteForVnetState struct {
}

type ServerFarmRouteForVnetState struct {
}

func (ServerFarmRouteForVnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverFarmRouteForVnetState)(nil)).Elem()
}

type serverFarmRouteForVnetArgs struct {
	EndAddress        *string           `pulumi:"endAddress"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RouteName         *string           `pulumi:"routeName"`
	RouteType         *string           `pulumi:"routeType"`
	StartAddress      *string           `pulumi:"startAddress"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	VnetName          string            `pulumi:"vnetName"`
}


type ServerFarmRouteForVnetArgs struct {
	EndAddress        pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RouteName         pulumi.StringPtrInput
	RouteType         pulumi.StringPtrInput
	StartAddress      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	VnetName          pulumi.StringInput
}

func (ServerFarmRouteForVnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverFarmRouteForVnetArgs)(nil)).Elem()
}

type ServerFarmRouteForVnetInput interface {
	pulumi.Input

	ToServerFarmRouteForVnetOutput() ServerFarmRouteForVnetOutput
	ToServerFarmRouteForVnetOutputWithContext(ctx context.Context) ServerFarmRouteForVnetOutput
}

func (*ServerFarmRouteForVnet) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerFarmRouteForVnet)(nil))
}

func (i *ServerFarmRouteForVnet) ToServerFarmRouteForVnetOutput() ServerFarmRouteForVnetOutput {
	return i.ToServerFarmRouteForVnetOutputWithContext(context.Background())
}

func (i *ServerFarmRouteForVnet) ToServerFarmRouteForVnetOutputWithContext(ctx context.Context) ServerFarmRouteForVnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerFarmRouteForVnetOutput)
}

type ServerFarmRouteForVnetOutput struct{ *pulumi.OutputState }

func (ServerFarmRouteForVnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerFarmRouteForVnet)(nil))
}

func (o ServerFarmRouteForVnetOutput) ToServerFarmRouteForVnetOutput() ServerFarmRouteForVnetOutput {
	return o
}

func (o ServerFarmRouteForVnetOutput) ToServerFarmRouteForVnetOutputWithContext(ctx context.Context) ServerFarmRouteForVnetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerFarmRouteForVnetInput)(nil)).Elem(), &ServerFarmRouteForVnet{})
	pulumi.RegisterOutputType(ServerFarmRouteForVnetOutput{})
}
