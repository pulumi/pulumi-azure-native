


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MSTIDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  MSTIDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                   `pulumi:"etag"`
	Kind       pulumi.StringOutput                      `pulumi:"kind"`
	Name       pulumi.StringOutput                      `pulumi:"name"`
	SystemData SystemDataResponseOutput                 `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                      `pulumi:"tenantId"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewMSTIDataConnector(ctx *pulumi.Context,
	name string, args *MSTIDataConnectorArgs, opts ...pulumi.ResourceOption) (*MSTIDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MSTIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MSTIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20211001preview:MSTIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMSTIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MSTIDataConnectorState, opts ...pulumi.ResourceOption) (*MSTIDataConnector, error) {
	var resource MSTIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20211001preview:MSTIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mstidataConnectorState struct {
}

type MSTIDataConnectorState struct {
}

func (MSTIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mstidataConnectorState)(nil)).Elem()
}

type mstidataConnectorArgs struct {
	DataConnectorId   *string                    `pulumi:"dataConnectorId"`
	DataTypes         MSTIDataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                     `pulumi:"kind"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	TenantId          string                     `pulumi:"tenantId"`
	WorkspaceName     string                     `pulumi:"workspaceName"`
}


type MSTIDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         MSTIDataConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (MSTIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mstidataConnectorArgs)(nil)).Elem()
}

type MSTIDataConnectorInput interface {
	pulumi.Input

	ToMSTIDataConnectorOutput() MSTIDataConnectorOutput
	ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput
}

func (*MSTIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnector)(nil)).Elem()
}

func (i *MSTIDataConnector) ToMSTIDataConnectorOutput() MSTIDataConnectorOutput {
	return i.ToMSTIDataConnectorOutputWithContext(context.Background())
}

func (i *MSTIDataConnector) ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorOutput)
}

type MSTIDataConnectorOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnector)(nil)).Elem()
}

func (o MSTIDataConnectorOutput) ToMSTIDataConnectorOutput() MSTIDataConnectorOutput {
	return o
}

func (o MSTIDataConnectorOutput) ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput {
	return o
}

func (o MSTIDataConnectorOutput) DataTypes() MSTIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MSTIDataConnector) MSTIDataConnectorDataTypesResponseOutput { return v.DataTypes }).(MSTIDataConnectorDataTypesResponseOutput)
}

func (o MSTIDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MSTIDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MSTIDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MSTIDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MSTIDataConnectorOutput{})
}
