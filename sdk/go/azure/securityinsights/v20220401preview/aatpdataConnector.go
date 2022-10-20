


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AATPDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind       pulumi.StringOutput                            `pulumi:"kind"`
	Name       pulumi.StringOutput                            `pulumi:"name"`
	SystemData SystemDataResponseOutput                       `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                            `pulumi:"tenantId"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
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
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AzureAdvancedThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource AATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:AATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AATPDataConnectorState, opts ...pulumi.ResourceOption) (*AATPDataConnector, error) {
	var resource AATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:AATPDataConnector", name, id, state, &resource, opts...)
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
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	TenantId          string                         `pulumi:"tenantId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type AATPDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
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
	return reflect.TypeOf((**AATPDataConnector)(nil)).Elem()
}

func (i *AATPDataConnector) ToAATPDataConnectorOutput() AATPDataConnectorOutput {
	return i.ToAATPDataConnectorOutputWithContext(context.Background())
}

func (i *AATPDataConnector) ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AATPDataConnectorOutput)
}

type AATPDataConnectorOutput struct{ *pulumi.OutputState }

func (AATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AATPDataConnector)(nil)).Elem()
}

func (o AATPDataConnectorOutput) ToAATPDataConnectorOutput() AATPDataConnectorOutput {
	return o
}

func (o AATPDataConnectorOutput) ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput {
	return o
}

func (o AATPDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *AATPDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o AATPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AATPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AATPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AATPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AATPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AATPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o AATPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AATPDataConnectorOutput{})
}
