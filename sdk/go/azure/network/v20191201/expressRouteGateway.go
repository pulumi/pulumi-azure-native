


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRouteGateway struct {
	pulumi.CustomResourceState

	AutoScaleConfiguration  ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput `pulumi:"autoScaleConfiguration"`
	Etag                    pulumi.StringOutput                                                  `pulumi:"etag"`
	ExpressRouteConnections ExpressRouteConnectionResponseArrayOutput                            `pulumi:"expressRouteConnections"`
	Location                pulumi.StringPtrOutput                                               `pulumi:"location"`
	Name                    pulumi.StringOutput                                                  `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                                                  `pulumi:"provisioningState"`
	Tags                    pulumi.StringMapOutput                                               `pulumi:"tags"`
	Type                    pulumi.StringOutput                                                  `pulumi:"type"`
	VirtualHub              VirtualHubIdResponseOutput                                           `pulumi:"virtualHub"`
}


func NewExpressRouteGateway(ctx *pulumi.Context,
	name string, args *ExpressRouteGatewayArgs, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHub == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHub'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRouteGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteGateway
	err := ctx.RegisterResource("azure-native:network/v20191201:ExpressRouteGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteGatewayState, opts ...pulumi.ResourceOption) (*ExpressRouteGateway, error) {
	var resource ExpressRouteGateway
	err := ctx.ReadResource("azure-native:network/v20191201:ExpressRouteGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteGatewayState struct {
}

type ExpressRouteGatewayState struct {
}

func (ExpressRouteGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayState)(nil)).Elem()
}

type expressRouteGatewayArgs struct {
	AutoScaleConfiguration  *ExpressRouteGatewayPropertiesAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	ExpressRouteGatewayName *string                                              `pulumi:"expressRouteGatewayName"`
	Id                      *string                                              `pulumi:"id"`
	Location                *string                                              `pulumi:"location"`
	ResourceGroupName       string                                               `pulumi:"resourceGroupName"`
	Tags                    map[string]string                                    `pulumi:"tags"`
	VirtualHub              VirtualHubId                                         `pulumi:"virtualHub"`
}


type ExpressRouteGatewayArgs struct {
	AutoScaleConfiguration  ExpressRouteGatewayPropertiesAutoScaleConfigurationPtrInput
	ExpressRouteGatewayName pulumi.StringPtrInput
	Id                      pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
	VirtualHub              VirtualHubIdInput
}

func (ExpressRouteGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteGatewayArgs)(nil)).Elem()
}

type ExpressRouteGatewayInput interface {
	pulumi.Input

	ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput
	ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput
}

func (*ExpressRouteGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteGateway)(nil)).Elem()
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput {
	return i.ToExpressRouteGatewayOutputWithContext(context.Background())
}

func (i *ExpressRouteGateway) ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteGatewayOutput)
}

type ExpressRouteGatewayOutput struct{ *pulumi.OutputState }

func (ExpressRouteGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteGateway)(nil)).Elem()
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayOutput() ExpressRouteGatewayOutput {
	return o
}

func (o ExpressRouteGatewayOutput) ToExpressRouteGatewayOutputWithContext(ctx context.Context) ExpressRouteGatewayOutput {
	return o
}

func (o ExpressRouteGatewayOutput) AutoScaleConfiguration() ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput {
		return v.AutoScaleConfiguration
	}).(ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput)
}

func (o ExpressRouteGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteGatewayOutput) ExpressRouteConnections() ExpressRouteConnectionResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) ExpressRouteConnectionResponseArrayOutput {
		return v.ExpressRouteConnections
	}).(ExpressRouteConnectionResponseArrayOutput)
}

func (o ExpressRouteGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExpressRouteGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExpressRouteGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExpressRouteGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ExpressRouteGatewayOutput) VirtualHub() VirtualHubIdResponseOutput {
	return o.ApplyT(func(v *ExpressRouteGateway) VirtualHubIdResponseOutput { return v.VirtualHub }).(VirtualHubIdResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteGatewayOutput{})
}
