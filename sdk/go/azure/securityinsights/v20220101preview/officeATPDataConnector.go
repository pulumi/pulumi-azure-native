


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OfficeATPDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind       pulumi.StringOutput                            `pulumi:"kind"`
	Name       pulumi.StringOutput                            `pulumi:"name"`
	SystemData SystemDataResponseOutput                       `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                            `pulumi:"tenantId"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewOfficeATPDataConnector(ctx *pulumi.Context,
	name string, args *OfficeATPDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeATPDataConnector, error) {
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
	args.Kind = pulumi.String("OfficeATP")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficeATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:OfficeATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficeATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:OfficeATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOfficeATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeATPDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeATPDataConnector, error) {
	var resource OfficeATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:OfficeATPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type officeATPDataConnectorState struct {
}

type OfficeATPDataConnectorState struct {
}

func (OfficeATPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeATPDataConnectorState)(nil)).Elem()
}

type officeATPDataConnectorArgs struct {
	DataConnectorId   *string                        `pulumi:"dataConnectorId"`
	DataTypes         *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Kind              string                         `pulumi:"kind"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	TenantId          string                         `pulumi:"tenantId"`
	WorkspaceName     string                         `pulumi:"workspaceName"`
}


type OfficeATPDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         AlertsDataTypeOfDataConnectorPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (OfficeATPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeATPDataConnectorArgs)(nil)).Elem()
}

type OfficeATPDataConnectorInput interface {
	pulumi.Input

	ToOfficeATPDataConnectorOutput() OfficeATPDataConnectorOutput
	ToOfficeATPDataConnectorOutputWithContext(ctx context.Context) OfficeATPDataConnectorOutput
}

func (*OfficeATPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeATPDataConnector)(nil)).Elem()
}

func (i *OfficeATPDataConnector) ToOfficeATPDataConnectorOutput() OfficeATPDataConnectorOutput {
	return i.ToOfficeATPDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeATPDataConnector) ToOfficeATPDataConnectorOutputWithContext(ctx context.Context) OfficeATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeATPDataConnectorOutput)
}

type OfficeATPDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeATPDataConnector)(nil)).Elem()
}

func (o OfficeATPDataConnectorOutput) ToOfficeATPDataConnectorOutput() OfficeATPDataConnectorOutput {
	return o
}

func (o OfficeATPDataConnectorOutput) ToOfficeATPDataConnectorOutputWithContext(ctx context.Context) OfficeATPDataConnectorOutput {
	return o
}

func (o OfficeATPDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o OfficeATPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OfficeATPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o OfficeATPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OfficeATPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OfficeATPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o OfficeATPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeATPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficeATPDataConnectorOutput{})
}
