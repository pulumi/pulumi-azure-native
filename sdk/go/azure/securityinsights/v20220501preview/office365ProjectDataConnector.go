


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Office365ProjectDataConnector struct {
	pulumi.CustomResourceState

	DataTypes  Office365ProjectConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                           `pulumi:"etag"`
	Kind       pulumi.StringOutput                              `pulumi:"kind"`
	Name       pulumi.StringOutput                              `pulumi:"name"`
	SystemData SystemDataResponseOutput                         `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                              `pulumi:"tenantId"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewOffice365ProjectDataConnector(ctx *pulumi.Context,
	name string, args *Office365ProjectDataConnectorArgs, opts ...pulumi.ResourceOption) (*Office365ProjectDataConnector, error) {
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
	args.Kind = pulumi.String("Office365Project")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Office365ProjectDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource Office365ProjectDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220501preview:Office365ProjectDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOffice365ProjectDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Office365ProjectDataConnectorState, opts ...pulumi.ResourceOption) (*Office365ProjectDataConnector, error) {
	var resource Office365ProjectDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220501preview:Office365ProjectDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type office365ProjectDataConnectorState struct {
}

type Office365ProjectDataConnectorState struct {
}

func (Office365ProjectDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*office365ProjectDataConnectorState)(nil)).Elem()
}

type office365ProjectDataConnectorArgs struct {
	DataConnectorId   *string                            `pulumi:"dataConnectorId"`
	DataTypes         Office365ProjectConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                             `pulumi:"kind"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	TenantId          string                             `pulumi:"tenantId"`
	WorkspaceName     string                             `pulumi:"workspaceName"`
}


type Office365ProjectDataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         Office365ProjectConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (Office365ProjectDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*office365ProjectDataConnectorArgs)(nil)).Elem()
}

type Office365ProjectDataConnectorInput interface {
	pulumi.Input

	ToOffice365ProjectDataConnectorOutput() Office365ProjectDataConnectorOutput
	ToOffice365ProjectDataConnectorOutputWithContext(ctx context.Context) Office365ProjectDataConnectorOutput
}

func (*Office365ProjectDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**Office365ProjectDataConnector)(nil)).Elem()
}

func (i *Office365ProjectDataConnector) ToOffice365ProjectDataConnectorOutput() Office365ProjectDataConnectorOutput {
	return i.ToOffice365ProjectDataConnectorOutputWithContext(context.Background())
}

func (i *Office365ProjectDataConnector) ToOffice365ProjectDataConnectorOutputWithContext(ctx context.Context) Office365ProjectDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Office365ProjectDataConnectorOutput)
}

type Office365ProjectDataConnectorOutput struct{ *pulumi.OutputState }

func (Office365ProjectDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Office365ProjectDataConnector)(nil)).Elem()
}

func (o Office365ProjectDataConnectorOutput) ToOffice365ProjectDataConnectorOutput() Office365ProjectDataConnectorOutput {
	return o
}

func (o Office365ProjectDataConnectorOutput) ToOffice365ProjectDataConnectorOutputWithContext(ctx context.Context) Office365ProjectDataConnectorOutput {
	return o
}

func (o Office365ProjectDataConnectorOutput) DataTypes() Office365ProjectConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) Office365ProjectConnectorDataTypesResponseOutput {
		return v.DataTypes
	}).(Office365ProjectConnectorDataTypesResponseOutput)
}

func (o Office365ProjectDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o Office365ProjectDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o Office365ProjectDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Office365ProjectDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o Office365ProjectDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o Office365ProjectDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Office365ProjectDataConnectorOutput{})
}
