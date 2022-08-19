


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Dynamics365DataConnector struct {
	pulumi.CustomResourceState

	DataTypes  Dynamics365DataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	Etag       pulumi.StringPtrOutput                          `pulumi:"etag"`
	Kind       pulumi.StringOutput                             `pulumi:"kind"`
	Name       pulumi.StringOutput                             `pulumi:"name"`
	SystemData SystemDataResponseOutput                        `pulumi:"systemData"`
	TenantId   pulumi.StringOutput                             `pulumi:"tenantId"`
	Type       pulumi.StringOutput                             `pulumi:"type"`
}


func NewDynamics365DataConnector(ctx *pulumi.Context,
	name string, args *Dynamics365DataConnectorArgs, opts ...pulumi.ResourceOption) (*Dynamics365DataConnector, error) {
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
	args.Kind = pulumi.String("Dynamics365")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Dynamics365DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Dynamics365DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource Dynamics365DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:Dynamics365DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDynamics365DataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Dynamics365DataConnectorState, opts ...pulumi.ResourceOption) (*Dynamics365DataConnector, error) {
	var resource Dynamics365DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:Dynamics365DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dynamics365DataConnectorState struct {
}

type Dynamics365DataConnectorState struct {
}

func (Dynamics365DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamics365DataConnectorState)(nil)).Elem()
}

type dynamics365DataConnectorArgs struct {
	DataConnectorId   *string                           `pulumi:"dataConnectorId"`
	DataTypes         Dynamics365DataConnectorDataTypes `pulumi:"dataTypes"`
	Kind              string                            `pulumi:"kind"`
	ResourceGroupName string                            `pulumi:"resourceGroupName"`
	TenantId          string                            `pulumi:"tenantId"`
	WorkspaceName     string                            `pulumi:"workspaceName"`
}


type Dynamics365DataConnectorArgs struct {
	DataConnectorId   pulumi.StringPtrInput
	DataTypes         Dynamics365DataConnectorDataTypesInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (Dynamics365DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamics365DataConnectorArgs)(nil)).Elem()
}

type Dynamics365DataConnectorInput interface {
	pulumi.Input

	ToDynamics365DataConnectorOutput() Dynamics365DataConnectorOutput
	ToDynamics365DataConnectorOutputWithContext(ctx context.Context) Dynamics365DataConnectorOutput
}

func (*Dynamics365DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnector)(nil)).Elem()
}

func (i *Dynamics365DataConnector) ToDynamics365DataConnectorOutput() Dynamics365DataConnectorOutput {
	return i.ToDynamics365DataConnectorOutputWithContext(context.Background())
}

func (i *Dynamics365DataConnector) ToDynamics365DataConnectorOutputWithContext(ctx context.Context) Dynamics365DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorOutput)
}

type Dynamics365DataConnectorOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnector)(nil)).Elem()
}

func (o Dynamics365DataConnectorOutput) ToDynamics365DataConnectorOutput() Dynamics365DataConnectorOutput {
	return o
}

func (o Dynamics365DataConnectorOutput) ToDynamics365DataConnectorOutputWithContext(ctx context.Context) Dynamics365DataConnectorOutput {
	return o
}

func (o Dynamics365DataConnectorOutput) DataTypes() Dynamics365DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) Dynamics365DataConnectorDataTypesResponseOutput { return v.DataTypes }).(Dynamics365DataConnectorDataTypesResponseOutput)
}

func (o Dynamics365DataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o Dynamics365DataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o Dynamics365DataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Dynamics365DataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o Dynamics365DataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o Dynamics365DataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynamics365DataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Dynamics365DataConnectorOutput{})
}
