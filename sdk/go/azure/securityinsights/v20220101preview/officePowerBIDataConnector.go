


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OfficePowerBIDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  OfficePowerBIConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                        `pulumi:"etag"`
	Kind       pulumi.StringOutput                           `pulumi:"kind"`
	Name       pulumi.StringOutput                           `pulumi:"name"`
	SystemData SystemDataResponseOutput                      `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                           `pulumi:"tenantId"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewOfficePowerBIDataConnector(ctx *pulumi.Context,
	name string, args *OfficePowerBIDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficePowerBIDataConnector, error) {
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
	args.Kind = pulumi.String("OfficePowerBI")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficePowerBIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficePowerBIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:OfficePowerBIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOfficePowerBIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficePowerBIDataConnectorState, opts ...pulumi.ResourceOption) (*OfficePowerBIDataConnector, error) {
	var resource OfficePowerBIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:OfficePowerBIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type officePowerBIDataConnectorState struct {
}

type OfficePowerBIDataConnectorState struct {
}

func (OfficePowerBIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officePowerBIDataConnectorState)(nil)).Elem()
}

type officePowerBIDataConnectorArgs struct {
	DataConnectorId   *string                         `pulumi:"dataConnectorId"`
	DataTypes         OfficePowerBIConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                          `pulumi:"kind"`
	ResourceGroupName string                          `pulumi:"resourceGroupName"`
	TenantId          string                          `pulumi:"tenantId"`
	WorkspaceName     string                          `pulumi:"workspaceName"`
}


type OfficePowerBIDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         OfficePowerBIConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (OfficePowerBIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officePowerBIDataConnectorArgs)(nil)).Elem()
}

type OfficePowerBIDataConnectorInput interface {
	pulumi.Input

	ToOfficePowerBIDataConnectorOutput() OfficePowerBIDataConnectorOutput
	ToOfficePowerBIDataConnectorOutputWithContext(ctx context.Context) OfficePowerBIDataConnectorOutput
}

func (*OfficePowerBIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficePowerBIDataConnector)(nil)).Elem()
}

func (i *OfficePowerBIDataConnector) ToOfficePowerBIDataConnectorOutput() OfficePowerBIDataConnectorOutput {
	return i.ToOfficePowerBIDataConnectorOutputWithContext(context.Background())
}

func (i *OfficePowerBIDataConnector) ToOfficePowerBIDataConnectorOutputWithContext(ctx context.Context) OfficePowerBIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficePowerBIDataConnectorOutput)
}

type OfficePowerBIDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficePowerBIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficePowerBIDataConnector)(nil)).Elem()
}

func (o OfficePowerBIDataConnectorOutput) ToOfficePowerBIDataConnectorOutput() OfficePowerBIDataConnectorOutput {
	return o
}

func (o OfficePowerBIDataConnectorOutput) ToOfficePowerBIDataConnectorOutputWithContext(ctx context.Context) OfficePowerBIDataConnectorOutput {
	return o
}

func (o OfficePowerBIDataConnectorOutput) DataTypes() OfficePowerBIConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) OfficePowerBIConnectorDataTypesResponseOutput { return v.DataTypes }).(OfficePowerBIConnectorDataTypesResponseOutput)
}

func (o OfficePowerBIDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OfficePowerBIDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o OfficePowerBIDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OfficePowerBIDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OfficePowerBIDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o OfficePowerBIDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficePowerBIDataConnectorOutput{})
}
