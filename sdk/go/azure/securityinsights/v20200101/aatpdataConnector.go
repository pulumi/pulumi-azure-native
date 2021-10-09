


package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AATPDataConnector struct {
	pulumi.CustomResourceState

	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag      pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind      pulumi.StringOutput                            `pulumi:"kind"`
	Name      pulumi.StringOutput                            `pulumi:"name"`
	TenantId  pulumi.StringPtrOutput                         `pulumi:"tenantId"`
	Type      pulumi.StringOutput                            `pulumi:"type"`
}


func NewAATPDataConnector(ctx *pulumi.Context,
	name string, args *AATPDataConnectorArgs, opts ...pulumi.ResourceOption) (*AATPDataConnector, error) {
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
	args.Kind = pulumi.String("AzureAdvancedThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:AATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20200101:AATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AATPDataConnectorState, opts ...pulumi.ResourceOption) (*AATPDataConnector, error) {
	var resource AATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20200101:AATPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type aatpdataConnectorState struct {
}

type AATPDataConnectorState struct {
}

func (AATPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*aatpdataConnectorState)(nil)).Elem()
}

type aatpdataConnectorArgs struct {
	DataConnectorId   *string                        `pulumi:"dataConnectorId"`
	DataTypes         *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Etag              *string                        `pulumi:"etag"`
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	TenantId          *string                        `pulumi:"tenantId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type AATPDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Etag              pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (AATPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aatpdataConnectorArgs)(nil)).Elem()
}

type AATPDataConnectorInput interface {
	pulumi.Input

	ToAATPDataConnectorOutput() AATPDataConnectorOutput
	ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput
}

func (*AATPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*AATPDataConnector)(nil))
}

func (i *AATPDataConnector) ToAATPDataConnectorOutput() AATPDataConnectorOutput {
	return i.ToAATPDataConnectorOutputWithContext(context.Background())
}

func (i *AATPDataConnector) ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AATPDataConnectorOutput)
}

type AATPDataConnectorOutput struct{ *pulumi.OutputState }

func (AATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AATPDataConnector)(nil))
}

func (o AATPDataConnectorOutput) ToAATPDataConnectorOutput() AATPDataConnectorOutput {
	return o
}

func (o AATPDataConnectorOutput) ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AATPDataConnectorOutput{})
}
