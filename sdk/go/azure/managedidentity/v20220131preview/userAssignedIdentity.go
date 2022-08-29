


package v20220131preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserAssignedIdentity struct {
	pulumi.CustomResourceState

	ClientId    pulumi.StringOutput    `pulumi:"clientId"`
	Location    pulumi.StringOutput    `pulumi:"location"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	PrincipalId pulumi.StringOutput    `pulumi:"principalId"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	TenantId    pulumi.StringOutput    `pulumi:"tenantId"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewUserAssignedIdentity(ctx *pulumi.Context,
	name string, args *UserAssignedIdentityArgs, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managedidentity:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20150831preview:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20181130:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-native:managedidentity/v20210930preview:UserAssignedIdentity"),
		},
	})
	opts = append(opts, aliases)
	var resource UserAssignedIdentity
	err := ctx.RegisterResource("azure-native:managedidentity/v20220131preview:UserAssignedIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUserAssignedIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAssignedIdentityState, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	var resource UserAssignedIdentity
	err := ctx.ReadResource("azure-native:managedidentity/v20220131preview:UserAssignedIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userAssignedIdentityState struct {
}

type UserAssignedIdentityState struct {
}

func (UserAssignedIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityState)(nil)).Elem()
}

type userAssignedIdentityArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Tags              map[string]string `pulumi:"tags"`
}


type UserAssignedIdentityArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityArgs)(nil)).Elem()
}

type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput
}

func (*UserAssignedIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentity)(nil)).Elem()
}

func (i *UserAssignedIdentity) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i *UserAssignedIdentity) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}

type UserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o UserAssignedIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAssignedIdentity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
}
