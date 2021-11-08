


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connector struct {
	pulumi.CustomResourceState

	AuthenticationDetails pulumi.AnyOutput                                 `pulumi:"authenticationDetails"`
	HybridComputeSettings HybridComputeSettingsPropertiesResponsePtrOutput `pulumi:"hybridComputeSettings"`
	Name                  pulumi.StringOutput                              `pulumi:"name"`
	Type                  pulumi.StringOutput                              `pulumi:"type"`
}


func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		args = &ConnectorArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security:Connector"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101preview:Connector"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20200101preview:Connector"),
		},
	})
	opts = append(opts, aliases)
	var resource Connector
	err := ctx.RegisterResource("azure-native:security:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azure-native:security:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectorState struct {
}

type ConnectorState struct {
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	AuthenticationDetails interface{}                      `pulumi:"authenticationDetails"`
	ConnectorName         *string                          `pulumi:"connectorName"`
	HybridComputeSettings *HybridComputeSettingsProperties `pulumi:"hybridComputeSettings"`
}


type ConnectorArgs struct {
	AuthenticationDetails pulumi.Input
	ConnectorName         pulumi.StringPtrInput
	HybridComputeSettings HybridComputeSettingsPropertiesPtrInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}

type ConnectorInput interface {
	pulumi.Input

	ToConnectorOutput() ConnectorOutput
	ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput
}

func (*Connector) ElementType() reflect.Type {
	return reflect.TypeOf((*Connector)(nil))
}

func (i *Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i *Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct{ *pulumi.OutputState }

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connector)(nil))
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectorOutput{})
}
