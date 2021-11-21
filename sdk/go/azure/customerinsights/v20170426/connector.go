


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
