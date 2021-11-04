


package v20151031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionType struct {
	pulumi.CustomResourceState

	CreationTime     pulumi.StringOutput              `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput           `pulumi:"description"`
	FieldDefinitions FieldDefinitionResponseMapOutput `pulumi:"fieldDefinitions"`
	IsGlobal         pulumi.BoolPtrOutput             `pulumi:"isGlobal"`
	LastModifiedTime pulumi.StringPtrOutput           `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput              `pulumi:"name"`
	Type             pulumi.StringOutput              `pulumi:"type"`
}


func NewConnectionType(ctx *pulumi.Context,
	name string, args *ConnectionTypeArgs, opts ...pulumi.ResourceOption) (*ConnectionType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.FieldDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'FieldDefinitions'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:ConnectionType"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20200113preview:ConnectionType"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectionType
	err := ctx.RegisterResource("azure-native:automation/v20151031:ConnectionType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectionType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionTypeState, opts ...pulumi.ResourceOption) (*ConnectionType, error) {
	var resource ConnectionType
	err := ctx.ReadResource("azure-native:automation/v20151031:ConnectionType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionTypeState struct {
}

type ConnectionTypeState struct {
}

func (ConnectionTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionTypeState)(nil)).Elem()
}

type connectionTypeArgs struct {
	AutomationAccountName string                     `pulumi:"automationAccountName"`
	ConnectionTypeName    *string                    `pulumi:"connectionTypeName"`
	FieldDefinitions      map[string]FieldDefinition `pulumi:"fieldDefinitions"`
	IsGlobal              *bool                      `pulumi:"isGlobal"`
	Name                  string                     `pulumi:"name"`
	ResourceGroupName     string                     `pulumi:"resourceGroupName"`
}


type ConnectionTypeArgs struct {
	AutomationAccountName pulumi.StringInput
	ConnectionTypeName    pulumi.StringPtrInput
	FieldDefinitions      FieldDefinitionMapInput
	IsGlobal              pulumi.BoolPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (ConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionTypeArgs)(nil)).Elem()
}

type ConnectionTypeInput interface {
	pulumi.Input

	ToConnectionTypeOutput() ConnectionTypeOutput
	ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput
}

func (*ConnectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionType)(nil))
}

func (i *ConnectionType) ToConnectionTypeOutput() ConnectionTypeOutput {
	return i.ToConnectionTypeOutputWithContext(context.Background())
}

func (i *ConnectionType) ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeOutput)
}

type ConnectionTypeOutput struct{ *pulumi.OutputState }

func (ConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionType)(nil))
}

func (o ConnectionTypeOutput) ToConnectionTypeOutput() ConnectionTypeOutput {
	return o
}

func (o ConnectionTypeOutput) ToConnectionTypeOutputWithContext(ctx context.Context) ConnectionTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionTypeOutput{})
}
