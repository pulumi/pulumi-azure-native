


package v20210301preview

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
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
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
	})
	opts = append(opts, aliases)
	var resource CodelessUiDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:CodelessUiDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodelessUiDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodelessUiDataConnectorState, opts ...pulumi.ResourceOption) (*CodelessUiDataConnector, error) {
	var resource CodelessUiDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:CodelessUiDataConnector", name, id, state, &resource, opts...)
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
	ConnectorUiConfig                   *CodelessUiConnectorConfigProperties `pulumi:"connectorUiConfig"`
	DataConnectorId                     *string                              `pulumi:"dataConnectorId"`
	Etag                                *string                              `pulumi:"etag"`
	Kind                                string                               `pulumi:"kind"`
	OperationalInsightsResourceProvider string                               `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string                               `pulumi:"resourceGroupName"`
	WorkspaceName                       string                               `pulumi:"workspaceName"`
}


type CodelessUiDataConnectorArgs struct {
	ConnectorUiConfig                   CodelessUiConnectorConfigPropertiesPtrInput
	DataConnectorId                     pulumi.StringPtrInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
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
	return reflect.TypeOf((*CodelessUiDataConnector)(nil))
}

func (i *CodelessUiDataConnector) ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput {
	return i.ToCodelessUiDataConnectorOutputWithContext(context.Background())
}

func (i *CodelessUiDataConnector) ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiDataConnectorOutput)
}

type CodelessUiDataConnectorOutput struct{ *pulumi.OutputState }

func (CodelessUiDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiDataConnector)(nil))
}

func (o CodelessUiDataConnectorOutput) ToCodelessUiDataConnectorOutput() CodelessUiDataConnectorOutput {
	return o
}

func (o CodelessUiDataConnectorOutput) ToCodelessUiDataConnectorOutputWithContext(ctx context.Context) CodelessUiDataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CodelessUiDataConnectorOutput{})
}
