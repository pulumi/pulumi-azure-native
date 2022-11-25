


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Group struct {
	pulumi.CustomResourceState

	BuiltIn     pulumi.BoolOutput      `pulumi:"builtIn"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName pulumi.StringOutput    `pulumi:"displayName"`
	ExternalId  pulumi.StringPtrOutput `pulumi:"externalId"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Group"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Group"),
		},
	})
	opts = append(opts, aliases)
	var resource Group
	err := ctx.RegisterResource("azure-native:apimanagement/v20220401preview:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure-native:apimanagement/v20220401preview:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	Description       *string    `pulumi:"description"`
	DisplayName       string     `pulumi:"displayName"`
	ExternalId        *string    `pulumi:"externalId"`
	GroupId           *string    `pulumi:"groupId"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
	ServiceName       string     `pulumi:"serviceName"`
	Type              *GroupType `pulumi:"type"`
}


type GroupArgs struct {
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	ExternalId        pulumi.StringPtrInput
	GroupId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Type              GroupTypePtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolOutput { return v.BuiltIn }).(pulumi.BoolOutput)
}

func (o GroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GroupOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
