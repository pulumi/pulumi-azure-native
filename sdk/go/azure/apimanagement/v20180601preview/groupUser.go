


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GroupUser struct {
	pulumi.CustomResourceState

	Email            pulumi.StringPtrOutput                     `pulumi:"email"`
	FirstName        pulumi.StringPtrOutput                     `pulumi:"firstName"`
	Groups           GroupContractPropertiesResponseArrayOutput `pulumi:"groups"`
	Identities       UserIdentityContractResponseArrayOutput    `pulumi:"identities"`
	LastName         pulumi.StringPtrOutput                     `pulumi:"lastName"`
	Name             pulumi.StringOutput                        `pulumi:"name"`
	Note             pulumi.StringPtrOutput                     `pulumi:"note"`
	RegistrationDate pulumi.StringPtrOutput                     `pulumi:"registrationDate"`
	State            pulumi.StringPtrOutput                     `pulumi:"state"`
	Type             pulumi.StringOutput                        `pulumi:"type"`
}


func NewGroupUser(ctx *pulumi.Context,
	name string, args *GroupUserArgs, opts ...pulumi.ResourceOption) (*GroupUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:GroupUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:GroupUser"),
		},
	})
	opts = append(opts, aliases)
	var resource GroupUser
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:GroupUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGroupUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupUserState, opts ...pulumi.ResourceOption) (*GroupUser, error) {
	var resource GroupUser
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:GroupUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type groupUserState struct {
}

type GroupUserState struct {
}

func (GroupUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupUserState)(nil)).Elem()
}

type groupUserArgs struct {
	GroupId           string  `pulumi:"groupId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	UserId            *string `pulumi:"userId"`
}


type GroupUserArgs struct {
	GroupId           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	UserId            pulumi.StringPtrInput
}

func (GroupUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupUserArgs)(nil)).Elem()
}

type GroupUserInput interface {
	pulumi.Input

	ToGroupUserOutput() GroupUserOutput
	ToGroupUserOutputWithContext(ctx context.Context) GroupUserOutput
}

func (*GroupUser) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupUser)(nil)).Elem()
}

func (i *GroupUser) ToGroupUserOutput() GroupUserOutput {
	return i.ToGroupUserOutputWithContext(context.Background())
}

func (i *GroupUser) ToGroupUserOutputWithContext(ctx context.Context) GroupUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupUserOutput)
}

type GroupUserOutput struct{ *pulumi.OutputState }

func (GroupUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupUser)(nil)).Elem()
}

func (o GroupUserOutput) ToGroupUserOutput() GroupUserOutput {
	return o
}

func (o GroupUserOutput) ToGroupUserOutputWithContext(ctx context.Context) GroupUserOutput {
	return o
}

func (o GroupUserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

func (o GroupUserOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringPtrOutput { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o GroupUserOutput) Groups() GroupContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *GroupUser) GroupContractPropertiesResponseArrayOutput { return v.Groups }).(GroupContractPropertiesResponseArrayOutput)
}

func (o GroupUserOutput) Identities() UserIdentityContractResponseArrayOutput {
	return o.ApplyT(func(v *GroupUser) UserIdentityContractResponseArrayOutput { return v.Identities }).(UserIdentityContractResponseArrayOutput)
}

func (o GroupUserOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringPtrOutput { return v.LastName }).(pulumi.StringPtrOutput)
}

func (o GroupUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GroupUserOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

func (o GroupUserOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringPtrOutput { return v.RegistrationDate }).(pulumi.StringPtrOutput)
}

func (o GroupUserOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o GroupUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupUserOutput{})
}
