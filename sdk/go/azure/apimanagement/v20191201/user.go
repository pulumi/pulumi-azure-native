


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type User struct {
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


func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if isZero(args.State) {
		args.State = pulumi.StringPtr("active")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:User"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:apimanagement/v20191201:User", name, id, state, &resource, opts...)
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
	AppType           *string                `pulumi:"appType"`
	Confirmation      *string                `pulumi:"confirmation"`
	Email             string                 `pulumi:"email"`
	FirstName         string                 `pulumi:"firstName"`
	Identities        []UserIdentityContract `pulumi:"identities"`
	LastName          string                 `pulumi:"lastName"`
	Note              *string                `pulumi:"note"`
	Notify            *bool                  `pulumi:"notify"`
	Password          *string                `pulumi:"password"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	ServiceName       string                 `pulumi:"serviceName"`
	State             *string                `pulumi:"state"`
	UserId            *string                `pulumi:"userId"`
}


type UserArgs struct {
	AppType           pulumi.StringPtrInput
	Confirmation      pulumi.StringPtrInput
	Email             pulumi.StringInput
	FirstName         pulumi.StringInput
	Identities        UserIdentityContractArrayInput
	LastName          pulumi.StringInput
	Note              pulumi.StringPtrInput
	Notify            pulumi.BoolPtrInput
	Password          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	State             pulumi.StringPtrInput
	UserId            pulumi.StringPtrInput
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

func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

func (o UserOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Groups() GroupContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *User) GroupContractPropertiesResponseArrayOutput { return v.Groups }).(GroupContractPropertiesResponseArrayOutput)
}

func (o UserOutput) Identities() UserIdentityContractResponseArrayOutput {
	return o.ApplyT(func(v *User) UserIdentityContractResponseArrayOutput { return v.Identities }).(UserIdentityContractResponseArrayOutput)
}

func (o UserOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.LastName }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

func (o UserOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.RegistrationDate }).(pulumi.StringPtrOutput)
}

func (o UserOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
