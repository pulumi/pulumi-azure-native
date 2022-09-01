


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ASCDataConnector struct {
	pulumi.CustomResourceState

	DataTypes      AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag           pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind           pulumi.StringOutput                            `pulumi:"kind"`
	Name           pulumi.StringOutput                            `pulumi:"name"`
	SubscriptionId pulumi.StringPtrOutput                         `pulumi:"subscriptionId"`
	SystemData     SystemDataResponseOutput                       `pulumi:"systemData"`
	Type           pulumi.StringOutput                            `pulumi:"type"`
}


func NewASCDataConnector(ctx *pulumi.Context,
	name string, args *ASCDataConnectorArgs, opts ...pulumi.ResourceOption) (*ASCDataConnector, error) {
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
	args.Kind = pulumi.String("AzureSecurityCenter")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ASCDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ASCDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource ASCDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:ASCDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetASCDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ASCDataConnectorState, opts ...pulumi.ResourceOption) (*ASCDataConnector, error) {
	var resource ASCDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:ASCDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ascdataConnectorState struct {
}

type ASCDataConnectorState struct {
}

func (ASCDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ascdataConnectorState)(nil)).Elem()
}

type ascdataConnectorArgs struct {
	DataConnectorId   *string                        `pulumi:"dataConnectorId"`
	DataTypes         *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	SubscriptionId    *string                        `pulumi:"subscriptionId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type ASCDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SubscriptionId    pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (ASCDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ascdataConnectorArgs)(nil)).Elem()
}

type ASCDataConnectorInput interface {
	pulumi.Input

	ToASCDataConnectorOutput() ASCDataConnectorOutput
	ToASCDataConnectorOutputWithContext(ctx context.Context) ASCDataConnectorOutput
}

func (*ASCDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**ASCDataConnector)(nil)).Elem()
}

func (i *ASCDataConnector) ToASCDataConnectorOutput() ASCDataConnectorOutput {
	return i.ToASCDataConnectorOutputWithContext(context.Background())
}

func (i *ASCDataConnector) ToASCDataConnectorOutputWithContext(ctx context.Context) ASCDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ASCDataConnectorOutput)
}

type ASCDataConnectorOutput struct{ *pulumi.OutputState }

func (ASCDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ASCDataConnector)(nil)).Elem()
}

func (o ASCDataConnectorOutput) ToASCDataConnectorOutput() ASCDataConnectorOutput {
	return o
}

func (o ASCDataConnectorOutput) ToASCDataConnectorOutputWithContext(ctx context.Context) ASCDataConnectorOutput {
	return o
}

func (o ASCDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *ASCDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o ASCDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ASCDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ASCDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ASCDataConnectorOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ASCDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ASCDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ASCDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ASCDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ASCDataConnectorOutput{})
}
