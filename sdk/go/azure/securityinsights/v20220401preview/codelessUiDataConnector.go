


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CodelessUiDataConnector struct {
	pulumi.CustomResourceState

	ConnectorUiConfig CodelessUiConnectorConfigPropertiesResponsePtrOutput `pulumi:"connectorUiConfig"`
	Etag              pulumi.StringPtrOutput                               `pulumi:"etag"`
	Kind              pulumi.StringOutput                                  `pulumi:"kind"`
	Name              pulumi.StringOutput                                  `pulumi:"name"`
	SystemData        SystemDataResponseOutput                             `pulumi:"systemData"`
	Type              pulumi.StringOutput                                  `pulumi:"type"`
}


func NewCodelessUiDataConnector(ctx *pulumi.Context,
	name string, args *CodelessUiDataConnectorArgs, opts ...pulumi.ResourceOption) (*CodelessUiDataConnector, error) {
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
	args.Kind = pulumi.String("GenericUI")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:CodelessUiDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:CodelessUiDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource CodelessUiDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:CodelessUiDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodelessUiDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodelessUiDataConnectorState, opts ...pulumi.ResourceOption) (*CodelessUiDataConnector, error) {
	var resource CodelessUiDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:CodelessUiDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type codelessUiDataConnectorState struct {
}

type CodelessUiDataConnectorState struct {
}

func (CodelessUiDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessUiDataConnectorState)(nil)).Elem()
}

type codelessUiDataConnectorArgs struct {
	ConnectorUiConfig *CodelessUiConnectorConfigProperties `pulumi:"connectorUiConfig"`
	DataConnectorId   *string                              `pulumi:"dataConnectorId"`
	Kind              string                               `pulumi:"kind"`
	ResourceGroupName string                               `pulumi:"resourceGroupName"`
	WorkspaceName     string                               `pulumi:"workspaceName"`
}


type CodelessUiDataConnectorArgs struct {
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesPtrInput
	DataConnectorId   pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (CodelessUiDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessUiDataConnectorArgs)(nil)).Elem()
}

type CodelessUiDataConnectorInput interface {
	pulumi.Input

	ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput
	ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput
}

func (*CodelessUiDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiDataConnector)(nil)).Elem()
}

func (i *CodelessUiDataConnector) ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput {
	return i.ToCodelessUiDataConnectorOutputWithContext(context.Background())
}

func (i *CodelessUiDataConnector) ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiDataConnectorOutput)
}

type CodelessUiDataConnectorOutput struct{ *pulumi.OutputState }

func (CodelessUiDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiDataConnector)(nil)).Elem()
}

func (o CodelessUiDataConnectorOutput) ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput {
	return o
}

func (o CodelessUiDataConnectorOutput) ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput {
	return o
}

func (o CodelessUiDataConnectorOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

func (o CodelessUiDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o CodelessUiDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o CodelessUiDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CodelessUiDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CodelessUiDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessUiDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodelessUiDataConnectorOutput{})
}
