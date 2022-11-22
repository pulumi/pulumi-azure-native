


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CodelessApiPollingDataConnector struct {
	pulumi.CustomResourceState

	ConnectorUiConfig CodelessUiConnectorConfigPropertiesResponsePtrOutput      `pulumi:"connectorUiConfig"`
	Etag              pulumi.StringPtrOutput                                    `pulumi:"etag"`
	Kind              pulumi.StringOutput                                       `pulumi:"kind"`
	Name              pulumi.StringOutput                                       `pulumi:"name"`
	PollingConfig     CodelessConnectorPollingConfigPropertiesResponsePtrOutput `pulumi:"pollingConfig"`
	SystemData        SystemDataResponseOutput                                  `pulumi:"systemData"`
	Type              pulumi.StringOutput                                       `pulumi:"type"`
}


func NewCodelessApiPollingDataConnector(ctx *pulumi.Context,
	name string, args *CodelessApiPollingDataConnectorArgs, opts ...pulumi.ResourceOption) (*CodelessApiPollingDataConnector, error) {
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
	args.Kind = pulumi.String("APIPolling")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:CodelessApiPollingDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:CodelessApiPollingDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource CodelessApiPollingDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:CodelessApiPollingDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodelessApiPollingDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodelessApiPollingDataConnectorState, opts ...pulumi.ResourceOption) (*CodelessApiPollingDataConnector, error) {
	var resource CodelessApiPollingDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:CodelessApiPollingDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type codelessApiPollingDataConnectorState struct {
}

type CodelessApiPollingDataConnectorState struct {
}

func (CodelessApiPollingDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessApiPollingDataConnectorState)(nil)).Elem()
}

type codelessApiPollingDataConnectorArgs struct {
	ConnectorUiConfig *CodelessUiConnectorConfigProperties      `pulumi:"connectorUiConfig"`
	DataConnectorId   *string                                   `pulumi:"dataConnectorId"`
	Kind              string                                    `pulumi:"kind"`
	PollingConfig     *CodelessConnectorPollingConfigProperties `pulumi:"pollingConfig"`
	ResourceGroupName string                                    `pulumi:"resourceGroupName"`
	WorkspaceName     string                                    `pulumi:"workspaceName"`
}


type CodelessApiPollingDataConnectorArgs struct {
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesPtrInput
	DataConnectorId   pulumi.StringPtrInput
	Kind              pulumi.StringInput
	PollingConfig     CodelessConnectorPollingConfigPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (CodelessApiPollingDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codelessApiPollingDataConnectorArgs)(nil)).Elem()
}

type CodelessApiPollingDataConnectorInput interface {
	pulumi.Input

	ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput
	ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput
}

func (*CodelessApiPollingDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessApiPollingDataConnector)(nil)).Elem()
}

func (i *CodelessApiPollingDataConnector) ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput {
	return i.ToCodelessApiPollingDataConnectorOutputWithContext(context.Background())
}

func (i *CodelessApiPollingDataConnector) ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessApiPollingDataConnectorOutput)
}

type CodelessApiPollingDataConnectorOutput struct{ *pulumi.OutputState }

func (CodelessApiPollingDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessApiPollingDataConnector)(nil)).Elem()
}

func (o CodelessApiPollingDataConnectorOutput) ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput {
	return o
}

func (o CodelessApiPollingDataConnectorOutput) ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput {
	return o
}

func (o CodelessApiPollingDataConnectorOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

func (o CodelessApiPollingDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o CodelessApiPollingDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o CodelessApiPollingDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CodelessApiPollingDataConnectorOutput) PollingConfig() CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
		return v.PollingConfig
	}).(CodelessConnectorPollingConfigPropertiesResponsePtrOutput)
}

func (o CodelessApiPollingDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CodelessApiPollingDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodelessApiPollingDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodelessApiPollingDataConnectorOutput{})
}
