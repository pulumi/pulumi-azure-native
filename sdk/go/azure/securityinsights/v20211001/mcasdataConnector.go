


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MCASDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  MCASDataConnectorDataTypesResponsePtrOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                      `pulumi:"etag"`
	Kind       pulumi.StringOutput                         `pulumi:"kind"`
	Name       pulumi.StringOutput                         `pulumi:"name"`
	SystemData SystemDataResponseOutput                    `pulumi:"systemData"`
	TenantId   pulumi.StringPtrOutput                      `pulumi:"tenantId"`
	Type       pulumi.StringOutput                         `pulumi:"type"`
}


func NewMCASDataConnector(ctx *pulumi.Context,
	name string, args *MCASDataConnectorArgs, opts ...pulumi.ResourceOption) (*MCASDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftCloudAppSecurity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MCASDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MCASDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource MCASDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20211001:MCASDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMCASDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MCASDataConnectorState, opts ...pulumi.ResourceOption) (*MCASDataConnector, error) {
	var resource MCASDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20211001:MCASDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mcasdataConnectorState struct {
}

type MCASDataConnectorState struct {
}

func (MCASDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mcasdataConnectorState)(nil)).Elem()
}

type mcasdataConnectorArgs struct {
	DataConnectorId   *string                     `pulumi:"dataConnectorId"`
	DataTypes         *MCASDataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                      `pulumi:"kind"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	TenantId          *string                     `pulumi:"tenantId"`
	WorkspaceName     string                      `pulumi:"workspaceName"`
}


type MCASDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         MCASDataConnectorDataTypesPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (MCASDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mcasdataConnectorArgs)(nil)).Elem()
}

type MCASDataConnectorInput interface {
	pulumi.Input

	ToMCASDataConnectorOutput() MCASDataConnectorOutput
	ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput
}

func (*MCASDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnector)(nil)).Elem()
}

func (i *MCASDataConnector) ToMCASDataConnectorOutput() MCASDataConnectorOutput {
	return i.ToMCASDataConnectorOutputWithContext(context.Background())
}

func (i *MCASDataConnector) ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorOutput)
}

type MCASDataConnectorOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnector)(nil)).Elem()
}

func (o MCASDataConnectorOutput) ToMCASDataConnectorOutput() MCASDataConnectorOutput {
	return o
}

func (o MCASDataConnectorOutput) ToMCASDataConnectorOutputWithContext(ctx context.Context) MCASDataConnectorOutput {
	return o
}

func (o MCASDataConnectorOutput) DataTypes() MCASDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyT(func(v *MCASDataConnector) MCASDataConnectorDataTypesResponsePtrOutput { return v.DataTypes }).(MCASDataConnectorDataTypesResponsePtrOutput)
}

func (o MCASDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MCASDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MCASDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MCASDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MCASDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MCASDataConnectorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o MCASDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MCASDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MCASDataConnectorOutput{})
}
