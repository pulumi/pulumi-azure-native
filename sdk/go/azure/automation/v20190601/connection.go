


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connection struct {
	pulumi.CustomResourceState

	ConnectionType        ConnectionTypeAssociationPropertyResponsePtrOutput `pulumi:"connectionType"`
	CreationTime          pulumi.StringOutput                                `pulumi:"creationTime"`
	Description           pulumi.StringPtrOutput                             `pulumi:"description"`
	FieldDefinitionValues pulumi.StringMapOutput                             `pulumi:"fieldDefinitionValues"`
	LastModifiedTime      pulumi.StringOutput                                `pulumi:"lastModifiedTime"`
	Name                  pulumi.StringOutput                                `pulumi:"name"`
	Type                  pulumi.StringOutput                                `pulumi:"type"`
}


func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ConnectionType == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionType'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Connection"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Connection"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Connection"),
		},
	})
	opts = append(opts, aliases)
	var resource Connection
	err := ctx.RegisterResource("azure-native:automation/v20190601:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azure-native:automation/v20190601:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	AutomationAccountName string                            `pulumi:"automationAccountName"`
	ConnectionName        *string                           `pulumi:"connectionName"`
	ConnectionType        ConnectionTypeAssociationProperty `pulumi:"connectionType"`
	Description           *string                           `pulumi:"description"`
	FieldDefinitionValues map[string]string                 `pulumi:"fieldDefinitionValues"`
	Name                  string                            `pulumi:"name"`
	ResourceGroupName     string                            `pulumi:"resourceGroupName"`
}


type ConnectionArgs struct {
	AutomationAccountName pulumi.StringInput
	ConnectionName        pulumi.StringPtrInput
	ConnectionType        ConnectionTypeAssociationPropertyInput
	Description           pulumi.StringPtrInput
	FieldDefinitionValues pulumi.StringMapInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ConnectionType() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *Connection) ConnectionTypeAssociationPropertyResponsePtrOutput { return v.ConnectionType }).(ConnectionTypeAssociationPropertyResponsePtrOutput)
}

func (o ConnectionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) FieldDefinitionValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.FieldDefinitionValues }).(pulumi.StringMapOutput)
}

func (o ConnectionOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}
