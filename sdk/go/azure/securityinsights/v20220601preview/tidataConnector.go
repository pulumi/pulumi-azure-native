


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TIDataConnector struct {
	pulumi.CustomResourceState

	DataTypes         TIDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag              pulumi.StringPtrOutput                 `pulumi:"etag"`
	Kind              pulumi.StringOutput                    `pulumi:"kind"`
	Name              pulumi.StringOutput                    `pulumi:"name"`
	SystemData        SystemDataResponseOutput               `pulumi:"systemData"`
	TenantId          pulumi.StringOutput                    `pulumi:"tenantId"`
	TipLookbackPeriod pulumi.StringPtrOutput                 `pulumi:"tipLookbackPeriod"`
	Type              pulumi.StringOutput                    `pulumi:"type"`
}


func NewTIDataConnector(ctx *pulumi.Context,
	name string, args *TIDataConnectorArgs, opts ...pulumi.ResourceOption) (*TIDataConnector, error) {
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
	args.Kind = pulumi.String("ThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:TIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:TIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource TIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220601preview:TIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TIDataConnectorState, opts ...pulumi.ResourceOption) (*TIDataConnector, error) {
	var resource TIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220601preview:TIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tidataConnectorState struct {
}

type TIDataConnectorState struct {
}

func (TIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*tidataConnectorState)(nil)).Elem()
}

type tidataConnectorArgs struct {
	DataConnectorId   *string                  `pulumi:"dataConnectorId"`
	DataTypes         TIDataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                   `pulumi:"kind"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	TenantId          string                   `pulumi:"tenantId"`
	TipLookbackPeriod *string                  `pulumi:"tipLookbackPeriod"`
	WorkspaceName     string                   `pulumi:"workspaceName"`
}


type TIDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         TIDataConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	TipLookbackPeriod pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (TIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tidataConnectorArgs)(nil)).Elem()
}

type TIDataConnectorInput interface {
	pulumi.Input

	ToTIDataConnectorOutput() TIDataConnectorOutput
	ToTIDataConnectorOutputWithContext(ctx context.Context) TIDataConnectorOutput
}

func (*TIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnector)(nil)).Elem()
}

func (i *TIDataConnector) ToTIDataConnectorOutput() TIDataConnectorOutput {
	return i.ToTIDataConnectorOutputWithContext(context.Background())
}

func (i *TIDataConnector) ToTIDataConnectorOutputWithContext(ctx context.Context) TIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorOutput)
}

type TIDataConnectorOutput struct{ *pulumi.OutputState }

func (TIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnector)(nil)).Elem()
}

func (o TIDataConnectorOutput) ToTIDataConnectorOutput() TIDataConnectorOutput {
	return o
}

func (o TIDataConnectorOutput) ToTIDataConnectorOutputWithContext(ctx context.Context) TIDataConnectorOutput {
	return o
}

func (o TIDataConnectorOutput) DataTypes() TIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *TIDataConnector) TIDataConnectorDataTypesResponseOutput { return v.DataTypes }).(TIDataConnectorDataTypesResponseOutput)
}

func (o TIDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o TIDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o TIDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TIDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TIDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TIDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o TIDataConnectorOutput) TipLookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringPtrOutput { return v.TipLookbackPeriod }).(pulumi.StringPtrOutput)
}

func (o TIDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TIDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TIDataConnectorOutput{})
}
