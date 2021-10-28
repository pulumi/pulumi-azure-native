


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DataConnector struct {
	pulumi.CustomResourceState

	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	Kind pulumi.StringOutput    `pulumi:"kind"`
	Name pulumi.StringOutput    `pulumi:"name"`
	Type pulumi.StringOutput    `pulumi:"type"`
}


func NewDataConnector(ctx *pulumi.Context,
	name string, args *DataConnectorArgs, opts ...pulumi.ResourceOption) (*DataConnector, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:DataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorState, opts ...pulumi.ResourceOption) (*DataConnector, error) {
	var resource DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataConnectorState struct {
}

type DataConnectorState struct {
}

func (DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorState)(nil)).Elem()
}

type dataConnectorArgs struct {
	DataConnectorId                     *string `pulumi:"dataConnectorId"`
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type DataConnectorArgs struct {
	DataConnectorId                     pulumi.StringPtrInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorArgs)(nil)).Elem()
}

type DataConnectorInput interface {
	pulumi.Input

	ToDataConnectorOutput() DataConnectorOutput
	ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput
}

func (*DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnector)(nil))
}

func (i *DataConnector) ToDataConnectorOutput() DataConnectorOutput {
	return i.ToDataConnectorOutputWithContext(context.Background())
}

func (i *DataConnector) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorOutput)
}

type DataConnectorOutput struct{ *pulumi.OutputState }

func (DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnector)(nil))
}

func (o DataConnectorOutput) ToDataConnectorOutput() DataConnectorOutput {
	return o
}

func (o DataConnectorOutput) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataConnectorInput)(nil)).Elem(), &DataConnector{})
	pulumi.RegisterOutputType(DataConnectorOutput{})
}
