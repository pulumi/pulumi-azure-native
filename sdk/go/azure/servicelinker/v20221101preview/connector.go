


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connector struct {
	pulumi.CustomResourceState

	AuthInfo              pulumi.AnyOutput                       `pulumi:"authInfo"`
	ClientType            pulumi.StringPtrOutput                 `pulumi:"clientType"`
	ConfigurationInfo     ConfigurationInfoResponsePtrOutput     `pulumi:"configurationInfo"`
	Name                  pulumi.StringOutput                    `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                    `pulumi:"provisioningState"`
	PublicNetworkSolution PublicNetworkSolutionResponsePtrOutput `pulumi:"publicNetworkSolution"`
	Scope                 pulumi.StringPtrOutput                 `pulumi:"scope"`
	SecretStore           SecretStoreResponsePtrOutput           `pulumi:"secretStore"`
	SystemData            SystemDataResponseOutput               `pulumi:"systemData"`
	TargetService         pulumi.AnyOutput                       `pulumi:"targetService"`
	Type                  pulumi.StringOutput                    `pulumi:"type"`
	VNetSolution          VNetSolutionResponsePtrOutput          `pulumi:"vNetSolution"`
}


func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Connector
	err := ctx.RegisterResource("azure-native:servicelinker/v20221101preview:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azure-native:servicelinker/v20221101preview:Connector", name, id, state, &resource, opts...)
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
	AuthInfo              interface{}            `pulumi:"authInfo"`
	ClientType            *string                `pulumi:"clientType"`
	ConfigurationInfo     *ConfigurationInfo     `pulumi:"configurationInfo"`
	ConnectorName         *string                `pulumi:"connectorName"`
	Location              string                 `pulumi:"location"`
	PublicNetworkSolution *PublicNetworkSolution `pulumi:"publicNetworkSolution"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
	Scope                 *string                `pulumi:"scope"`
	SecretStore           *SecretStore           `pulumi:"secretStore"`
	TargetService         interface{}            `pulumi:"targetService"`
	VNetSolution          *VNetSolution          `pulumi:"vNetSolution"`
}


type ConnectorArgs struct {
	AuthInfo              pulumi.Input
	ClientType            pulumi.StringPtrInput
	ConfigurationInfo     ConfigurationInfoPtrInput
	ConnectorName         pulumi.StringPtrInput
	Location              pulumi.StringInput
	PublicNetworkSolution PublicNetworkSolutionPtrInput
	ResourceGroupName     pulumi.StringInput
	Scope                 pulumi.StringPtrInput
	SecretStore           SecretStorePtrInput
	TargetService         pulumi.Input
	VNetSolution          VNetSolutionPtrInput
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
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (i *Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i *Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct{ *pulumi.OutputState }

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

func (o ConnectorOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *Connector) pulumi.AnyOutput { return v.AuthInfo }).(pulumi.AnyOutput)
}

func (o ConnectorOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.ClientType }).(pulumi.StringPtrOutput)
}

func (o ConnectorOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *Connector) ConfigurationInfoResponsePtrOutput { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

func (o ConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConnectorOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v *Connector) PublicNetworkSolutionResponsePtrOutput { return v.PublicNetworkSolution }).(PublicNetworkSolutionResponsePtrOutput)
}

func (o ConnectorOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o ConnectorOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v *Connector) SecretStoreResponsePtrOutput { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

func (o ConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Connector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectorOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v *Connector) pulumi.AnyOutput { return v.TargetService }).(pulumi.AnyOutput)
}

func (o ConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ConnectorOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v *Connector) VNetSolutionResponsePtrOutput { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorOutput{})
}
