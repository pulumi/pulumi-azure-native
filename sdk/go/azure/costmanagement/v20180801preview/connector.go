


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connector struct {
	pulumi.CustomResourceState

	Collection        ConnectorCollectionInfoResponseOutput `pulumi:"collection"`
	CreatedOn         pulumi.StringOutput                   `pulumi:"createdOn"`
	CredentialsKey    pulumi.StringPtrOutput                `pulumi:"credentialsKey"`
	DisplayName       pulumi.StringPtrOutput                `pulumi:"displayName"`
	Kind              pulumi.StringPtrOutput                `pulumi:"kind"`
	Location          pulumi.StringPtrOutput                `pulumi:"location"`
	ModifiedOn        pulumi.StringOutput                   `pulumi:"modifiedOn"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	ProviderAccountId pulumi.StringOutput                   `pulumi:"providerAccountId"`
	ReportId          pulumi.StringPtrOutput                `pulumi:"reportId"`
	Status            pulumi.StringPtrOutput                `pulumi:"status"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	Type              pulumi.StringOutput                   `pulumi:"type"`
}


func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:Connector"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190301preview:Connector"),
		},
	})
	opts = append(opts, aliases)
	var resource Connector
	err := ctx.RegisterResource("azure-native:costmanagement/v20180801preview:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azure-native:costmanagement/v20180801preview:Connector", name, id, state, &resource, opts...)
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
	ConnectorName     *string           `pulumi:"connectorName"`
	CredentialsKey    *string           `pulumi:"credentialsKey"`
	CredentialsSecret *string           `pulumi:"credentialsSecret"`
	DisplayName       *string           `pulumi:"displayName"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	ReportId          *string           `pulumi:"reportId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
}


type ConnectorArgs struct {
	ConnectorName     pulumi.StringPtrInput
	CredentialsKey    pulumi.StringPtrInput
	CredentialsSecret pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ReportId          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
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
