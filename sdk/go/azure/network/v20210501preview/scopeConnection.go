


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScopeConnection struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput   `pulumi:"description"`
	Etag        pulumi.StringOutput      `pulumi:"etag"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	ResourceId  pulumi.StringPtrOutput   `pulumi:"resourceId"`
	SystemData  SystemDataResponseOutput `pulumi:"systemData"`
	TenantId    pulumi.StringPtrOutput   `pulumi:"tenantId"`
	Type        pulumi.StringOutput      `pulumi:"type"`
}


func NewScopeConnection(ctx *pulumi.Context,
	name string, args *ScopeConnectionArgs, opts ...pulumi.ResourceOption) (*ScopeConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ScopeConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ScopeConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:ScopeConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:ScopeConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ScopeConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ScopeConnection
	err := ctx.RegisterResource("azure-native:network/v20210501preview:ScopeConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScopeConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeConnectionState, opts ...pulumi.ResourceOption) (*ScopeConnection, error) {
	var resource ScopeConnection
	err := ctx.ReadResource("azure-native:network/v20210501preview:ScopeConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scopeConnectionState struct {
}

type ScopeConnectionState struct {
}

func (ScopeConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeConnectionState)(nil)).Elem()
}

type scopeConnectionArgs struct {
	Description         *string `pulumi:"description"`
	NetworkManagerName  string  `pulumi:"networkManagerName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	ResourceId          *string `pulumi:"resourceId"`
	ScopeConnectionName *string `pulumi:"scopeConnectionName"`
	TenantId            *string `pulumi:"tenantId"`
}


type ScopeConnectionArgs struct {
	Description         pulumi.StringPtrInput
	NetworkManagerName  pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	ResourceId          pulumi.StringPtrInput
	ScopeConnectionName pulumi.StringPtrInput
	TenantId            pulumi.StringPtrInput
}

func (ScopeConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeConnectionArgs)(nil)).Elem()
}

type ScopeConnectionInput interface {
	pulumi.Input

	ToScopeConnectionOutput() ScopeConnectionOutput
	ToScopeConnectionOutputWithContext(ctx context.Context) ScopeConnectionOutput
}

func (*ScopeConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeConnection)(nil)).Elem()
}

func (i *ScopeConnection) ToScopeConnectionOutput() ScopeConnectionOutput {
	return i.ToScopeConnectionOutputWithContext(context.Background())
}

func (i *ScopeConnection) ToScopeConnectionOutputWithContext(ctx context.Context) ScopeConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeConnectionOutput)
}

type ScopeConnectionOutput struct{ *pulumi.OutputState }

func (ScopeConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeConnection)(nil)).Elem()
}

func (o ScopeConnectionOutput) ToScopeConnectionOutput() ScopeConnectionOutput {
	return o
}

func (o ScopeConnectionOutput) ToScopeConnectionOutputWithContext(ctx context.Context) ScopeConnectionOutput {
	return o
}

func (o ScopeConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ScopeConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ScopeConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScopeConnectionOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeConnection) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ScopeConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScopeConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScopeConnectionOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeConnection) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ScopeConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScopeConnectionOutput{})
}
