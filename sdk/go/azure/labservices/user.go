


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type User struct {
	pulumi.CustomResourceState

	Email                 pulumi.StringOutput                 `pulumi:"email"`
	FamilyName            pulumi.StringOutput                 `pulumi:"familyName"`
	GivenName             pulumi.StringOutput                 `pulumi:"givenName"`
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	Location              pulumi.StringPtrOutput              `pulumi:"location"`
	Name                  pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	Tags                  pulumi.StringMapOutput              `pulumi:"tags"`
	TenantId              pulumi.StringOutput                 `pulumi:"tenantId"`
	TotalUsage            pulumi.StringOutput                 `pulumi:"totalUsage"`
	Type                  pulumi.StringOutput                 `pulumi:"type"`
	UniqueIdentifier      pulumi.StringPtrOutput              `pulumi:"uniqueIdentifier"`
}


func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20181015:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-native:labservices:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:labservices:User", name, id, state, &resource, opts...)
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
	LabAccountName    string            `pulumi:"labAccountName"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
	UserName          *string           `pulumi:"userName"`
}


type UserArgs struct {
	LabAccountName    pulumi.StringInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
	UserName          pulumi.StringPtrInput
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

func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o UserOutput) FamilyName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.FamilyName }).(pulumi.StringOutput)
}

func (o UserOutput) GivenName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.GivenName }).(pulumi.StringOutput)
}

func (o UserOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *User) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o UserOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *User) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o UserOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o UserOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.TotalUsage }).(pulumi.StringOutput)
}

func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o UserOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
