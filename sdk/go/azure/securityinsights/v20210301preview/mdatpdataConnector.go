


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MDATPDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                         `pulumi:"etag"`
	Kind       pulumi.StringOutput                            `pulumi:"kind"`
	Name       pulumi.StringOutput                            `pulumi:"name"`
	SystemData SystemDataResponseOutput                       `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                            `pulumi:"tenantId"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewMDATPDataConnector(ctx *pulumi.Context,
	name string, args *MDATPDataConnectorArgs, opts ...pulumi.ResourceOption) (*MDATPDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
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
	args.Kind = pulumi.String("MicrosoftDefenderAdvancedThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MDATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MDATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MDATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:MDATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMDATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MDATPDataConnectorState, opts ...pulumi.ResourceOption) (*MDATPDataConnector, error) {
	var resource MDATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:MDATPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mdatpdataConnectorState struct {
}

type MDATPDataConnectorState struct {
}

func (MDATPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdatpdataConnectorState)(nil)).Elem()
}

type mdatpdataConnectorArgs struct {
	DataConnectorId                     *string                        `pulumi:"dataConnectorId"`
	DataTypes                           *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	Kind                                string                         `pulumi:"kind"`
	OperationalInsightsResourceProvider string                         `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string                         `pulumi:"resourceGroupName"`
	TenantId                            string                         `pulumi:"tenantId"`
	WorkspaceName                       string                         `pulumi:"workspaceName"`
}


type MDATPDataConnectorArgs struct {
	DataConnectorId                     pulumi.StringPtrInput
	DataTypes                           AlertsDataTypeOfDataConnectorPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	TenantId                            pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (MDATPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdatpdataConnectorArgs)(nil)).Elem()
}

type MDATPDataConnectorInput interface {
	pulumi.Input

	ToMDATPDataConnectorOutput() MDATPDataConnectorOutput
	ToMDATPDataConnectorOutputWithContext(ctx context.Context) MDATPDataConnectorOutput
}

func (*MDATPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MDATPDataConnector)(nil)).Elem()
}

func (i *MDATPDataConnector) ToMDATPDataConnectorOutput() MDATPDataConnectorOutput {
	return i.ToMDATPDataConnectorOutputWithContext(context.Background())
}

func (i *MDATPDataConnector) ToMDATPDataConnectorOutputWithContext(ctx context.Context) MDATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MDATPDataConnectorOutput)
}

type MDATPDataConnectorOutput struct{ *pulumi.OutputState }

func (MDATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MDATPDataConnector)(nil)).Elem()
}

func (o MDATPDataConnectorOutput) ToMDATPDataConnectorOutput() MDATPDataConnectorOutput {
	return o
}

func (o MDATPDataConnectorOutput) ToMDATPDataConnectorOutputWithContext(ctx context.Context) MDATPDataConnectorOutput {
	return o
}

func (o MDATPDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *MDATPDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o MDATPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MDATPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MDATPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MDATPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MDATPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MDATPDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o MDATPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MDATPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MDATPDataConnectorOutput{})
}
