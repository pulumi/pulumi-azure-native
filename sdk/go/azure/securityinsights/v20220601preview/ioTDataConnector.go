


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IoTDataConnector struct {
	pulumi.CustomResourceState

	DataTypes      AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag           pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind           pulumi.StringOutput                            `pulumi:"kind"`
	Name           pulumi.StringOutput                            `pulumi:"name"`
	SubscriptionId pulumi.StringPtrOutput                         `pulumi:"subscriptionId"`
	SystemData     SystemDataResponseOutput                       `pulumi:"systemData"`
	Type           pulumi.StringOutput                            `pulumi:"type"`
}


func NewIoTDataConnector(ctx *pulumi.Context,
	name string, args *IoTDataConnectorArgs, opts ...pulumi.ResourceOption) (*IoTDataConnector, error) {
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
	args.Kind = pulumi.String("IOT")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:IoTDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:IoTDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource IoTDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220601preview:IoTDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIoTDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTDataConnectorState, opts ...pulumi.ResourceOption) (*IoTDataConnector, error) {
	var resource IoTDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220601preview:IoTDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ioTDataConnectorState struct {
}

type IoTDataConnectorState struct {
}

func (IoTDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTDataConnectorState)(nil)).Elem()
}

type ioTDataConnectorArgs struct {
	DataConnectorId   *string                        `pulumi:"dataConnectorId"`
	DataTypes         *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	SubscriptionId    *string                        `pulumi:"subscriptionId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type IoTDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SubscriptionId    pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (IoTDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTDataConnectorArgs)(nil)).Elem()
}

type IoTDataConnectorInput interface {
	pulumi.Input

	ToIoTDataConnectorOutput() IoTDataConnectorOutput
	ToIoTDataConnectorOutputWithContext(ctx context.Context) IoTDataConnectorOutput
}

func (*IoTDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDataConnector)(nil)).Elem()
}

func (i *IoTDataConnector) ToIoTDataConnectorOutput() IoTDataConnectorOutput {
	return i.ToIoTDataConnectorOutputWithContext(context.Background())
}

func (i *IoTDataConnector) ToIoTDataConnectorOutputWithContext(ctx context.Context) IoTDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTDataConnectorOutput)
}

type IoTDataConnectorOutput struct{ *pulumi.OutputState }

func (IoTDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTDataConnector)(nil)).Elem()
}

func (o IoTDataConnectorOutput) ToIoTDataConnectorOutput() IoTDataConnectorOutput {
	return o
}

func (o IoTDataConnectorOutput) ToIoTDataConnectorOutputWithContext(ctx context.Context) IoTDataConnectorOutput {
	return o
}

func (o IoTDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *IoTDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o IoTDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IoTDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o IoTDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IoTDataConnectorOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o IoTDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IoTDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IoTDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IoTDataConnectorOutput{})
}
