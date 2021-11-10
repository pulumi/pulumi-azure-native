


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionGateway struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                              `pulumi:"etag"`
	Location   pulumi.StringPtrOutput                              `pulumi:"location"`
	Name       pulumi.StringOutput                                 `pulumi:"name"`
	Properties ConnectionGatewayDefinitionResponsePropertiesOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                              `pulumi:"tags"`
	Type       pulumi.StringOutput                                 `pulumi:"type"`
}


func NewConnectionGateway(ctx *pulumi.Context,
	name string, args *ConnectionGatewayArgs, opts ...pulumi.ResourceOption) (*ConnectionGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:ConnectionGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectionGateway
	err := ctx.RegisterResource("azure-native:web/v20160601:ConnectionGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectionGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionGatewayState, opts ...pulumi.ResourceOption) (*ConnectionGateway, error) {
	var resource ConnectionGateway
	err := ctx.ReadResource("azure-native:web/v20160601:ConnectionGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionGatewayState struct {
}

type ConnectionGatewayState struct {
}

func (ConnectionGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionGatewayState)(nil)).Elem()
}

type connectionGatewayArgs struct {
	ConnectionGatewayName *string                                `pulumi:"connectionGatewayName"`
	Etag                  *string                                `pulumi:"etag"`
	Location              *string                                `pulumi:"location"`
	Properties            *ConnectionGatewayDefinitionProperties `pulumi:"properties"`
	ResourceGroupName     string                                 `pulumi:"resourceGroupName"`
	SubscriptionId        *string                                `pulumi:"subscriptionId"`
	Tags                  map[string]string                      `pulumi:"tags"`
}


type ConnectionGatewayArgs struct {
	ConnectionGatewayName pulumi.StringPtrInput
	Etag                  pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Properties            ConnectionGatewayDefinitionPropertiesPtrInput
	ResourceGroupName     pulumi.StringInput
	SubscriptionId        pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (ConnectionGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionGatewayArgs)(nil)).Elem()
}

type ConnectionGatewayInput interface {
	pulumi.Input

	ToConnectionGatewayOutput() ConnectionGatewayOutput
	ToConnectionGatewayOutputWithContext(ctx context.Context) ConnectionGatewayOutput
}

func (*ConnectionGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGateway)(nil))
}

func (i *ConnectionGateway) ToConnectionGatewayOutput() ConnectionGatewayOutput {
	return i.ToConnectionGatewayOutputWithContext(context.Background())
}

func (i *ConnectionGateway) ToConnectionGatewayOutputWithContext(ctx context.Context) ConnectionGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionGatewayOutput)
}

type ConnectionGatewayOutput struct{ *pulumi.OutputState }

func (ConnectionGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionGateway)(nil))
}

func (o ConnectionGatewayOutput) ToConnectionGatewayOutput() ConnectionGatewayOutput {
	return o
}

func (o ConnectionGatewayOutput) ToConnectionGatewayOutputWithContext(ctx context.Context) ConnectionGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionGatewayOutput{})
}
