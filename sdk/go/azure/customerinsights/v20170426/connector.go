


package v20170426

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connector struct {
	pulumi.CustomResourceState

	ConnectorId         pulumi.IntOutput       `pulumi:"connectorId"`
	ConnectorName       pulumi.StringPtrOutput `pulumi:"connectorName"`
	ConnectorProperties pulumi.MapOutput       `pulumi:"connectorProperties"`
	ConnectorType       pulumi.StringOutput    `pulumi:"connectorType"`
	Created             pulumi.StringOutput    `pulumi:"created"`
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName         pulumi.StringPtrOutput `pulumi:"displayName"`
	IsInternal          pulumi.BoolPtrOutput   `pulumi:"isInternal"`
	LastModified        pulumi.StringOutput    `pulumi:"lastModified"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	State               pulumi.StringOutput    `pulumi:"state"`
	TenantId            pulumi.StringOutput    `pulumi:"tenantId"`
	Type                pulumi.StringOutput    `pulumi:"type"`
}


func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorProperties == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorProperties'")
	}
	if args.ConnectorType == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorType'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights:Connector"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:Connector"),
		},
	})
	opts = append(opts, aliases)
	var resource Connector
	err := ctx.RegisterResource("azure-native:customerinsights/v20170426:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azure-native:customerinsights/v20170426:Connector", name, id, state, &resource, opts...)
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
	ConnectorName       *string                `pulumi:"connectorName"`
	ConnectorProperties map[string]interface{} `pulumi:"connectorProperties"`
	ConnectorType       string                 `pulumi:"connectorType"`
	Description         *string                `pulumi:"description"`
	DisplayName         *string                `pulumi:"displayName"`
	HubName             string                 `pulumi:"hubName"`
	IsInternal          *bool                  `pulumi:"isInternal"`
	ResourceGroupName   string                 `pulumi:"resourceGroupName"`
}


type ConnectorArgs struct {
	ConnectorName       pulumi.StringPtrInput
	ConnectorProperties pulumi.MapInput
	ConnectorType       pulumi.StringInput
	Description         pulumi.StringPtrInput
	DisplayName         pulumi.StringPtrInput
	HubName             pulumi.StringInput
	IsInternal          pulumi.BoolPtrInput
	ResourceGroupName   pulumi.StringInput
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

func (o ConnectorOutput) ConnectorId() pulumi.IntOutput {
	return o.ApplyT(func(v *Connector) pulumi.IntOutput { return v.ConnectorId }).(pulumi.IntOutput)
}

func (o ConnectorOutput) ConnectorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.ConnectorName }).(pulumi.StringPtrOutput)
}

func (o ConnectorOutput) ConnectorProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *Connector) pulumi.MapOutput { return v.ConnectorProperties }).(pulumi.MapOutput)
}

func (o ConnectorOutput) ConnectorType() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ConnectorType }).(pulumi.StringOutput)
}

func (o ConnectorOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o ConnectorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectorOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ConnectorOutput) IsInternal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.BoolPtrOutput { return v.IsInternal }).(pulumi.BoolPtrOutput)
}

func (o ConnectorOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o ConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectorOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o ConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorOutput{})
}
