


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DataConnector struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewDataConnector(ctx *pulumi.Context,
	name string, args *DataConnectorArgs, opts ...pulumi.ResourceOption) (*DataConnector, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220601preview:DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorState, opts ...pulumi.ResourceOption) (*DataConnector, error) {
	var resource DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220601preview:DataConnector", name, id, state, &resource, opts...)
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
	DataConnectorId   *string `pulumi:"dataConnectorId"`
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type DataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
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
	return reflect.TypeOf((**DataConnector)(nil)).Elem()
}

func (i *DataConnector) ToDataConnectorOutput() DataConnectorOutput {
	return i.ToDataConnectorOutputWithContext(context.Background())
}

func (i *DataConnector) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorOutput)
}

type DataConnectorOutput struct{ *pulumi.OutputState }

func (DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnector)(nil)).Elem()
}

func (o DataConnectorOutput) ToDataConnectorOutput() DataConnectorOutput {
	return o
}

func (o DataConnectorOutput) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return o
}

func (o DataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataConnectorOutput{})
}
