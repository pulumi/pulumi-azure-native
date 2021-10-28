


package v20170426

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectorMapping struct {
	pulumi.CustomResourceState

	ConnectorMappingName pulumi.StringOutput                      `pulumi:"connectorMappingName"`
	ConnectorName        pulumi.StringOutput                      `pulumi:"connectorName"`
	ConnectorType        pulumi.StringPtrOutput                   `pulumi:"connectorType"`
	Created              pulumi.StringOutput                      `pulumi:"created"`
	DataFormatId         pulumi.StringOutput                      `pulumi:"dataFormatId"`
	Description          pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName          pulumi.StringPtrOutput                   `pulumi:"displayName"`
	EntityType           pulumi.StringOutput                      `pulumi:"entityType"`
	EntityTypeName       pulumi.StringOutput                      `pulumi:"entityTypeName"`
	LastModified         pulumi.StringOutput                      `pulumi:"lastModified"`
	MappingProperties    ConnectorMappingPropertiesResponseOutput `pulumi:"mappingProperties"`
	Name                 pulumi.StringOutput                      `pulumi:"name"`
	NextRunTime          pulumi.StringOutput                      `pulumi:"nextRunTime"`
	RunId                pulumi.StringOutput                      `pulumi:"runId"`
	State                pulumi.StringOutput                      `pulumi:"state"`
	TenantId             pulumi.StringOutput                      `pulumi:"tenantId"`
	Type                 pulumi.StringOutput                      `pulumi:"type"`
}


func NewConnectorMapping(ctx *pulumi.Context,
	name string, args *ConnectorMappingArgs, opts ...pulumi.ResourceOption) (*ConnectorMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorName'")
	}
	if args.EntityType == nil {
		return nil, errors.New("invalid value for required argument 'EntityType'")
	}
	if args.EntityTypeName == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeName'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.MappingProperties == nil {
		return nil, errors.New("invalid value for required argument 'MappingProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170426:ConnectorMapping"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights:ConnectorMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:customerinsights:ConnectorMapping"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:ConnectorMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170101:ConnectorMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectorMapping
	err := ctx.RegisterResource("azure-native:customerinsights/v20170426:ConnectorMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectorMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorMappingState, opts ...pulumi.ResourceOption) (*ConnectorMapping, error) {
	var resource ConnectorMapping
	err := ctx.ReadResource("azure-native:customerinsights/v20170426:ConnectorMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectorMappingState struct {
}

type ConnectorMappingState struct {
}

func (ConnectorMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorMappingState)(nil)).Elem()
}

type connectorMappingArgs struct {
	ConnectorName     string                     `pulumi:"connectorName"`
	ConnectorType     *string                    `pulumi:"connectorType"`
	Description       *string                    `pulumi:"description"`
	DisplayName       *string                    `pulumi:"displayName"`
	EntityType        EntityTypes                `pulumi:"entityType"`
	EntityTypeName    string                     `pulumi:"entityTypeName"`
	HubName           string                     `pulumi:"hubName"`
	MappingName       *string                    `pulumi:"mappingName"`
	MappingProperties ConnectorMappingProperties `pulumi:"mappingProperties"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
}


type ConnectorMappingArgs struct {
	ConnectorName     pulumi.StringInput
	ConnectorType     pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	EntityType        EntityTypesInput
	EntityTypeName    pulumi.StringInput
	HubName           pulumi.StringInput
	MappingName       pulumi.StringPtrInput
	MappingProperties ConnectorMappingPropertiesInput
	ResourceGroupName pulumi.StringInput
}

func (ConnectorMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorMappingArgs)(nil)).Elem()
}

type ConnectorMappingInput interface {
	pulumi.Input

	ToConnectorMappingOutput() ConnectorMappingOutput
	ToConnectorMappingOutputWithContext(ctx context.Context) ConnectorMappingOutput
}

func (*ConnectorMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMapping)(nil))
}

func (i *ConnectorMapping) ToConnectorMappingOutput() ConnectorMappingOutput {
	return i.ToConnectorMappingOutputWithContext(context.Background())
}

func (i *ConnectorMapping) ToConnectorMappingOutputWithContext(ctx context.Context) ConnectorMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingOutput)
}

type ConnectorMappingOutput struct{ *pulumi.OutputState }

func (ConnectorMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorMapping)(nil))
}

func (o ConnectorMappingOutput) ToConnectorMappingOutput() ConnectorMappingOutput {
	return o
}

func (o ConnectorMappingOutput) ToConnectorMappingOutputWithContext(ctx context.Context) ConnectorMappingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorMappingInput)(nil)).Elem(), &ConnectorMapping{})
	pulumi.RegisterOutputType(ConnectorMappingOutput{})
}
