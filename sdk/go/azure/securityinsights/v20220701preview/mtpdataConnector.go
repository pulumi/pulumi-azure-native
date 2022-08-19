


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MTPDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  MTPDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                  `pulumi:"etag"`
	Kind       pulumi.StringOutput                     `pulumi:"kind"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                     `pulumi:"tenantId"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewMTPDataConnector(ctx *pulumi.Context,
	name string, args *MTPDataConnectorArgs, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MTPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MTPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MTPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220701preview:MTPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMTPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MTPDataConnectorState, opts ...pulumi.ResourceOption) (*MTPDataConnector, error) {
	var resource MTPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220701preview:MTPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mtpdataConnectorState struct {
}

type MTPDataConnectorState struct {
}

func (MTPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorState)(nil)).Elem()
}

type mtpdataConnectorArgs struct {
	DataConnectorId   *string                   `pulumi:"dataConnectorId"`
	DataTypes         MTPDataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                    `pulumi:"kind"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	TenantId          string                    `pulumi:"tenantId"`
	WorkspaceName     string                    `pulumi:"workspaceName"`
}


type MTPDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         MTPDataConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (MTPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mtpdataConnectorArgs)(nil)).Elem()
}

type MTPDataConnectorInput interface {
	pulumi.Input

	ToMTPDataConnectorOutput() MTPDataConnectorOutput
	ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput
}

func (*MTPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnector)(nil)).Elem()
}

func (i *MTPDataConnector) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return i.ToMTPDataConnectorOutputWithContext(context.Background())
}

func (i *MTPDataConnector) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorOutput)
}

type MTPDataConnectorOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnector)(nil)).Elem()
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutput() MTPDataConnectorOutput {
	return o
}

func (o MTPDataConnectorOutput) ToMTPDataConnectorOutputWithContext(ctx context.Context) MTPDataConnectorOutput {
	return o
}

func (o MTPDataConnectorOutput) DataTypes() MTPDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MTPDataConnector) MTPDataConnectorDataTypesResponseOutput { return v.DataTypes }).(MTPDataConnectorDataTypesResponseOutput)
}

func (o MTPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MTPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MTPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MTPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MTPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MTPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o MTPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MTPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MTPDataConnectorOutput{})
}
