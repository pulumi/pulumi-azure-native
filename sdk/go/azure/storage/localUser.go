


package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LocalUser struct {
	pulumi.CustomResourceState

	HasSharedKey      pulumi.BoolPtrOutput               `pulumi:"hasSharedKey"`
	HasSshKey         pulumi.BoolPtrOutput               `pulumi:"hasSshKey"`
	HasSshPassword    pulumi.BoolPtrOutput               `pulumi:"hasSshPassword"`
	HomeDirectory     pulumi.StringPtrOutput             `pulumi:"homeDirectory"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	PermissionScopes  PermissionScopeResponseArrayOutput `pulumi:"permissionScopes"`
	Sid               pulumi.StringOutput                `pulumi:"sid"`
	SshAuthorizedKeys SshPublicKeyResponseArrayOutput    `pulumi:"sshAuthorizedKeys"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewLocalUser(ctx *pulumi.Context,
	name string, args *LocalUserArgs, opts ...pulumi.ResourceOption) (*LocalUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage/v20210801:LocalUser"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:LocalUser"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:LocalUser"),
		},
	})
	opts = append(opts, aliases)
	var resource LocalUser
	err := ctx.RegisterResource("azure-native:storage:LocalUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLocalUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalUserState, opts ...pulumi.ResourceOption) (*LocalUser, error) {
	var resource LocalUser
	err := ctx.ReadResource("azure-native:storage:LocalUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type localUserState struct {
}

type LocalUserState struct {
}

func (LocalUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*localUserState)(nil)).Elem()
}

type localUserArgs struct {
	AccountName       string            `pulumi:"accountName"`
	HasSharedKey      *bool             `pulumi:"hasSharedKey"`
	HasSshKey         *bool             `pulumi:"hasSshKey"`
	HasSshPassword    *bool             `pulumi:"hasSshPassword"`
	HomeDirectory     *string           `pulumi:"homeDirectory"`
	PermissionScopes  []PermissionScope `pulumi:"permissionScopes"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SshAuthorizedKeys []SshPublicKey    `pulumi:"sshAuthorizedKeys"`
	Username          *string           `pulumi:"username"`
}


type LocalUserArgs struct {
	AccountName       pulumi.StringInput
	HasSharedKey      pulumi.BoolPtrInput
	HasSshKey         pulumi.BoolPtrInput
	HasSshPassword    pulumi.BoolPtrInput
	HomeDirectory     pulumi.StringPtrInput
	PermissionScopes  PermissionScopeArrayInput
	ResourceGroupName pulumi.StringInput
	SshAuthorizedKeys SshPublicKeyArrayInput
	Username          pulumi.StringPtrInput
}

func (LocalUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localUserArgs)(nil)).Elem()
}

type LocalUserInput interface {
	pulumi.Input

	ToLocalUserOutput() LocalUserOutput
	ToLocalUserOutputWithContext(ctx context.Context) LocalUserOutput
}

func (*LocalUser) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalUser)(nil)).Elem()
}

func (i *LocalUser) ToLocalUserOutput() LocalUserOutput {
	return i.ToLocalUserOutputWithContext(context.Background())
}

func (i *LocalUser) ToLocalUserOutputWithContext(ctx context.Context) LocalUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalUserOutput)
}

type LocalUserOutput struct{ *pulumi.OutputState }

func (LocalUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalUser)(nil)).Elem()
}

func (o LocalUserOutput) ToLocalUserOutput() LocalUserOutput {
	return o
}

func (o LocalUserOutput) ToLocalUserOutputWithContext(ctx context.Context) LocalUserOutput {
	return o
}

func (o LocalUserOutput) HasSharedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.BoolPtrOutput { return v.HasSharedKey }).(pulumi.BoolPtrOutput)
}

func (o LocalUserOutput) HasSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.BoolPtrOutput { return v.HasSshKey }).(pulumi.BoolPtrOutput)
}

func (o LocalUserOutput) HasSshPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.BoolPtrOutput { return v.HasSshPassword }).(pulumi.BoolPtrOutput)
}

func (o LocalUserOutput) HomeDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringPtrOutput { return v.HomeDirectory }).(pulumi.StringPtrOutput)
}

func (o LocalUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LocalUserOutput) PermissionScopes() PermissionScopeResponseArrayOutput {
	return o.ApplyT(func(v *LocalUser) PermissionScopeResponseArrayOutput { return v.PermissionScopes }).(PermissionScopeResponseArrayOutput)
}

func (o LocalUserOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringOutput { return v.Sid }).(pulumi.StringOutput)
}

func (o LocalUserOutput) SshAuthorizedKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *LocalUser) SshPublicKeyResponseArrayOutput { return v.SshAuthorizedKeys }).(SshPublicKeyResponseArrayOutput)
}

func (o LocalUserOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LocalUser) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LocalUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LocalUserOutput{})
}
