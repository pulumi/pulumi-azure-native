


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type User struct {
	pulumi.CustomResourceState

	AdditionalUsageQuota pulumi.StringPtrOutput   `pulumi:"additionalUsageQuota"`
	DisplayName          pulumi.StringOutput      `pulumi:"displayName"`
	Email                pulumi.StringOutput      `pulumi:"email"`
	InvitationSent       pulumi.StringOutput      `pulumi:"invitationSent"`
	InvitationState      pulumi.StringOutput      `pulumi:"invitationState"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	RegistrationState    pulumi.StringOutput      `pulumi:"registrationState"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	TotalUsage           pulumi.StringOutput      `pulumi:"totalUsage"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20211115preview:User"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20220801:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-native:labservices/v20211001preview:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:labservices/v20211001preview:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userState struct {
}

type UserState struct {
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	AdditionalUsageQuota *string `pulumi:"additionalUsageQuota"`
	Email                string  `pulumi:"email"`
	LabName              string  `pulumi:"labName"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
	UserName             *string `pulumi:"userName"`
}


type UserArgs struct {
	AdditionalUsageQuota pulumi.StringPtrInput
	Email                pulumi.StringInput
	LabName              pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	UserName             pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) AdditionalUsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.AdditionalUsageQuota }).(pulumi.StringPtrOutput)
}

func (o UserOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o UserOutput) InvitationSent() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.InvitationSent }).(pulumi.StringOutput)
}

func (o UserOutput) InvitationState() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.InvitationState }).(pulumi.StringOutput)
}

func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o UserOutput) RegistrationState() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.RegistrationState }).(pulumi.StringOutput)
}

func (o UserOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *User) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UserOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.TotalUsage }).(pulumi.StringOutput)
}

func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
