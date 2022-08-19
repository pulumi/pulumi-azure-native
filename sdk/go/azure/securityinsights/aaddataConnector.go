


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
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AADDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AADDataConnector"),
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
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	TenantId          *string                        `pulumi:"tenantId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type AADDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
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
	return reflect.TypeOf((**AADDataConnector)(nil)).Elem()
}

func (i *AADDataConnector) ToAADDataConnectorOutput() AADDataConnectorOutput {
	return i.ToAADDataConnectorOutputWithContext(context.Background())
}

func (i *AADDataConnector) ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADDataConnectorOutput)
}

type AADDataConnectorOutput struct{ *pulumi.OutputState }

func (AADDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADDataConnector)(nil)).Elem()
}

func (o AADDataConnectorOutput) ToAADDataConnectorOutput() AADDataConnectorOutput {
	return o
}

func (o AADDataConnectorOutput) ToAADDataConnectorOutputWithContext(ctx context.Context) AADDataConnectorOutput {
	return o
}

func (o AADDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *AADDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o AADDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AADDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AADDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AADDataConnectorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AADDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AADDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AADDataConnectorOutput{})
}
