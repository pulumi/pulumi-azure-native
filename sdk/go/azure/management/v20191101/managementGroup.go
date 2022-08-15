


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ManagementGroup struct {
	pulumi.CustomResourceState

	Children    ManagementGroupChildInfoResponseArrayOutput   `pulumi:"children"`
	Details     ManagementGroupDetailsResponsePtrOutput       `pulumi:"details"`
	DisplayName pulumi.StringPtrOutput                        `pulumi:"displayName"`
	Name        pulumi.StringOutput                           `pulumi:"name"`
	Path        ManagementGroupPathElementResponseArrayOutput `pulumi:"path"`
	Roles       pulumi.StringArrayOutput                      `pulumi:"roles"`
	TenantId    pulumi.StringPtrOutput                        `pulumi:"tenantId"`
	Type        pulumi.StringOutput                           `pulumi:"type"`
}


func NewManagementGroup(ctx *pulumi.Context,
	name string, args *ManagementGroupArgs, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	if args == nil {
		args = &ManagementGroupArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:management:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20171101preview:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20180101preview:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20180301preview:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20200201:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20200501:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20201001:ManagementGroup"),
		},
		{
			Type: pulumi.String("azure-native:management/v20210401:ManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementGroup
	err := ctx.RegisterResource("azure-native:management/v20191101:ManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupState, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	var resource ManagementGroup
	err := ctx.ReadResource("azure-native:management/v20191101:ManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementGroupState struct {
}

type ManagementGroupState struct {
}

func (ManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupState)(nil)).Elem()
}

type managementGroupArgs struct {
	Details     *CreateManagementGroupDetails `pulumi:"details"`
	DisplayName *string                       `pulumi:"displayName"`
	GroupId     *string                       `pulumi:"groupId"`
	Name        *string                       `pulumi:"name"`
}


type ManagementGroupArgs struct {
	Details     CreateManagementGroupDetailsPtrInput
	DisplayName pulumi.StringPtrInput
	GroupId     pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
}

func (ManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupArgs)(nil)).Elem()
}

type ManagementGroupInput interface {
	pulumi.Input

	ToManagementGroupOutput() ManagementGroupOutput
	ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput
}

func (*ManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroup)(nil)).Elem()
}

func (i *ManagementGroup) ToManagementGroupOutput() ManagementGroupOutput {
	return i.ToManagementGroupOutputWithContext(context.Background())
}

func (i *ManagementGroup) ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupOutput)
}

type ManagementGroupOutput struct{ *pulumi.OutputState }

func (ManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroup)(nil)).Elem()
}

func (o ManagementGroupOutput) ToManagementGroupOutput() ManagementGroupOutput {
	return o
}

func (o ManagementGroupOutput) ToManagementGroupOutputWithContext(ctx context.Context) ManagementGroupOutput {
	return o
}

func (o ManagementGroupOutput) Children() ManagementGroupChildInfoResponseArrayOutput {
	return o.ApplyT(func(v *ManagementGroup) ManagementGroupChildInfoResponseArrayOutput { return v.Children }).(ManagementGroupChildInfoResponseArrayOutput)
}

func (o ManagementGroupOutput) Details() ManagementGroupDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroup) ManagementGroupDetailsResponsePtrOutput { return v.Details }).(ManagementGroupDetailsResponsePtrOutput)
}

func (o ManagementGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementGroupOutput) Path() ManagementGroupPathElementResponseArrayOutput {
	return o.ApplyT(func(v *ManagementGroup) ManagementGroupPathElementResponseArrayOutput { return v.Path }).(ManagementGroupPathElementResponseArrayOutput)
}

func (o ManagementGroupOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementGroup) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o ManagementGroupOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroup) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupOutput{})
}
