


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OfficeIRMDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind       pulumi.StringOutput                            `pulumi:"kind"`
	Name       pulumi.StringOutput                            `pulumi:"name"`
	SystemData SystemDataResponseOutput                       `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                            `pulumi:"tenantId"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewOfficeIRMDataConnector(ctx *pulumi.Context,
	name string, args *OfficeIRMDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeIRMDataConnector, error) {
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
	args.Kind = pulumi.String("OfficeIRM")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:OfficeIRMDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficeIRMDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220601preview:OfficeIRMDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOfficeIRMDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeIRMDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeIRMDataConnector, error) {
	var resource OfficeIRMDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220601preview:OfficeIRMDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type officeIRMDataConnectorState struct {
}

type OfficeIRMDataConnectorState struct {
}

func (OfficeIRMDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeIRMDataConnectorState)(nil)).Elem()
}

type officeIRMDataConnectorArgs struct {
	DataConnectorId   *string                        `pulumi:"dataConnectorId"`
	DataTypes         *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	TenantId          string                         `pulumi:"tenantId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type OfficeIRMDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (OfficeIRMDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeIRMDataConnectorArgs)(nil)).Elem()
}

type OfficeIRMDataConnectorInput interface {
	pulumi.Input

	ToOfficeIRMDataConnectorOutput() OfficeIRMDataConnectorOutput
	ToOfficeIRMDataConnectorOutputWithContext(ctx context.Context) OfficeIRMDataConnectorOutput
}

func (*OfficeIRMDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeIRMDataConnector)(nil)).Elem()
}

func (i *OfficeIRMDataConnector) ToOfficeIRMDataConnectorOutput() OfficeIRMDataConnectorOutput {
	return i.ToOfficeIRMDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeIRMDataConnector) ToOfficeIRMDataConnectorOutputWithContext(ctx context.Context) OfficeIRMDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeIRMDataConnectorOutput)
}

type OfficeIRMDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeIRMDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeIRMDataConnector)(nil)).Elem()
}

func (o OfficeIRMDataConnectorOutput) ToOfficeIRMDataConnectorOutput() OfficeIRMDataConnectorOutput {
	return o
}

func (o OfficeIRMDataConnectorOutput) ToOfficeIRMDataConnectorOutputWithContext(ctx context.Context) OfficeIRMDataConnectorOutput {
	return o
}

func (o OfficeIRMDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o OfficeIRMDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OfficeIRMDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o OfficeIRMDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OfficeIRMDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OfficeIRMDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o OfficeIRMDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficeIRMDataConnectorOutput{})
}
