


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OfficeDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  OfficeDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                     `pulumi:"etag"`
	Kind       pulumi.StringOutput                        `pulumi:"kind"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                        `pulumi:"tenantId"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewOfficeDataConnector(ctx *pulumi.Context,
	name string, args *OfficeDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
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
	args.Kind = pulumi.String("Office365")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:OfficeDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:OfficeDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficeDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:OfficeDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOfficeDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeDataConnector, error) {
	var resource OfficeDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:OfficeDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type officeDataConnectorState struct {
}

type OfficeDataConnectorState struct {
}

func (OfficeDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeDataConnectorState)(nil)).Elem()
}

type officeDataConnectorArgs struct {
	DataConnectorId   *string                      `pulumi:"dataConnectorId"`
	DataTypes         OfficeDataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                       `pulumi:"kind"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	TenantId          string                       `pulumi:"tenantId"`
	WorkspaceName     string                       `pulumi:"workspaceName"`
}


type OfficeDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         OfficeDataConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (OfficeDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeDataConnectorArgs)(nil)).Elem()
}

type OfficeDataConnectorInput interface {
	pulumi.Input

	ToOfficeDataConnectorOutput() OfficeDataConnectorOutput
	ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput
}

func (*OfficeDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnector)(nil)).Elem()
}

func (i *OfficeDataConnector) ToOfficeDataConnectorOutput() OfficeDataConnectorOutput {
	return i.ToOfficeDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeDataConnector) ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorOutput)
}

type OfficeDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnector)(nil)).Elem()
}

func (o OfficeDataConnectorOutput) ToOfficeDataConnectorOutput() OfficeDataConnectorOutput {
	return o
}

func (o OfficeDataConnectorOutput) ToOfficeDataConnectorOutputWithContext(ctx context.Context) OfficeDataConnectorOutput {
	return o
}

func (o OfficeDataConnectorOutput) DataTypes() OfficeDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *OfficeDataConnector) OfficeDataConnectorDataTypesResponseOutput { return v.DataTypes }).(OfficeDataConnectorDataTypesResponseOutput)
}

func (o OfficeDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OfficeDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o OfficeDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OfficeDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficeDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OfficeDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o OfficeDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficeDataConnectorOutput{})
}
