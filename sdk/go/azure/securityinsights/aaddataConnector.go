


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AADDataConnector struct {
	pulumi.CustomResourceState

	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag      pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind      pulumi.StringOutput                            `pulumi:"kind"`
	Name      pulumi.StringOutput                            `pulumi:"name"`
	TenantId  pulumi.StringPtrOutput                         `pulumi:"tenantId"`
	Type      pulumi.StringOutput                            `pulumi:"type"`
}


func NewAADDataConnector(ctx *pulumi.Context,
	name string, args *AADDataConnectorArgs, opts ...pulumi.ResourceOption) (*AADDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AzureActiveDirectory")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AADDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AADDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights:AADDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAADDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AADDataConnectorState, opts ...pulumi.ResourceOption) (*AADDataConnector, error) {
	var resource AADDataConnector
	err := ctx.ReadResource("azure-native:securityinsights:AADDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type aaddataConnectorState struct {
}

type AADDataConnectorState struct {
}

func (AADDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*aaddataConnectorState)(nil)).Elem()
}

type aaddataConnectorArgs struct {
	DataConnectorId   *string                        `pulumi:"dataConnectorId"`
	DataTypes         *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Etag              *string                        `pulumi:"etag"`
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	TenantId          *string                        `pulumi:"tenantId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type AADDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Etag              pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (AADDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aaddataConnectorArgs)(nil)).Elem()
}

type AADDataConnectorInput interface {
	pulumi.Input

	ToAADDataConnectorOutput() AADDataConnectorOutput
	ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput
}

func (*AADDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*AADDataConnector)(nil))
}

func (i *AADDataConnector) ToAADDataConnectorOutput() AADDataConnectorOutput {
	return i.ToAADDataConnectorOutputWithContext(context.Background())
}

func (i *AADDataConnector) ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADDataConnectorOutput)
}

type AADDataConnectorOutput struct{ *pulumi.OutputState }

func (AADDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADDataConnector)(nil))
}

func (o AADDataConnectorOutput) ToAADDataConnectorOutput() AADDataConnectorOutput {
	return o
}

func (o AADDataConnectorOutput) ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AADDataConnectorOutput{})
}
