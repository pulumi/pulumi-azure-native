


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type User struct {
	pulumi.CustomResourceState

	CreatedDate       pulumi.StringOutput              `pulumi:"createdDate"`
	Identity          UserIdentityResponsePtrOutput    `pulumi:"identity"`
	Location          pulumi.StringPtrOutput           `pulumi:"location"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput              `pulumi:"provisioningState"`
	SecretStore       UserSecretStoreResponsePtrOutput `pulumi:"secretStore"`
	Tags              pulumi.StringMapOutput           `pulumi:"tags"`
	Type              pulumi.StringOutput              `pulumi:"type"`
	UniqueIdentifier  pulumi.StringOutput              `pulumi:"uniqueIdentifier"`
}


func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:User"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-native:devtestlab:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:devtestlab:User", name, id, state, &resource, opts...)
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
	Identity          *UserIdentity     `pulumi:"identity"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SecretStore       *UserSecretStore  `pulumi:"secretStore"`
	Tags              map[string]string `pulumi:"tags"`
}


type UserArgs struct {
	Identity          UserIdentityPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SecretStore       UserSecretStorePtrInput
	Tags              pulumi.StringMapInput
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

func (o UserOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o UserOutput) Identity() UserIdentityResponsePtrOutput {
	return o.ApplyT(func(v *User) UserIdentityResponsePtrOutput { return v.Identity }).(UserIdentityResponsePtrOutput)
}

func (o UserOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o UserOutput) SecretStore() UserSecretStoreResponsePtrOutput {
	return o.ApplyT(func(v *User) UserSecretStoreResponsePtrOutput { return v.SecretStore }).(UserSecretStoreResponsePtrOutput)
}

func (o UserOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *User) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o UserOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
