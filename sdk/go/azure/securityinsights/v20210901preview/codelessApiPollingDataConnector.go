


package v20210901preview

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
	})
	opts = append(opts, aliases)
	var resource CodelessApiPollingDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:CodelessApiPollingDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodelessApiPollingDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodelessApiPollingDataConnectorState, opts ...pulumi.ResourceOption) (*CodelessApiPollingDataConnector, error) {
	var resource CodelessApiPollingDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:CodelessApiPollingDataConnector", name, id, state, &resource, opts...)
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
	Etag              *string                                   `pulumi:"etag"`
	Kind              string                                    `pulumi:"kind"`
	PollingConfig     *CodelessConnectorPollingConfigProperties `pulumi:"pollingConfig"`
	ResourceGroupName string                                    `pulumi:"resourceGroupName"`
	WorkspaceName     string                                    `pulumi:"workspaceName"`
}


type CodelessApiPollingDataConnectorArgs struct {
	ConnectorUiConfig CodelessUiConnectorConfigPropertiesPtrInput
	DataConnectorId   pulumi.StringPtrInput
	Etag              pulumi.StringPtrInput
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
	return reflect.TypeOf((*CodelessApiPollingDataConnector)(nil))
}

func (i *CodelessApiPollingDataConnector) ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput {
	return i.ToCodelessApiPollingDataConnectorOutputWithContext(context.Background())
}

func (i *CodelessApiPollingDataConnector) ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessApiPollingDataConnectorOutput)
}

type CodelessApiPollingDataConnectorOutput struct{ *pulumi.OutputState }

func (CodelessApiPollingDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessApiPollingDataConnector)(nil))
}

func (o CodelessApiPollingDataConnectorOutput) ToCodelessApiPollingDataConnectorOutput() CodelessApiPollingDataConnectorOutput {
	return o
}

func (o CodelessApiPollingDataConnectorOutput) ToCodelessApiPollingDataConnectorOutputWithContext(ctx context.Context) CodelessApiPollingDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CodelessApiPollingDataConnectorOutput{})
}
