


package v20180915

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
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:User"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:User"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:User"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:User"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:User", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*User)(nil))
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterOutputType(UserOutput{})
}
