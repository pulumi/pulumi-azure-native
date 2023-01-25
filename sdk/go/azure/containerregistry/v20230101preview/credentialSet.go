


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CredentialSet struct {
	pulumi.CustomResourceState

	AuthCredentials   AuthCredentialResponseArrayOutput   `pulumi:"authCredentials"`
	CreationDate      pulumi.StringOutput                 `pulumi:"creationDate"`
	Identity          IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	LoginServer       pulumi.StringPtrOutput              `pulumi:"loginServer"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewCredentialSet(ctx *pulumi.Context,
	name string, args *CredentialSetArgs, opts ...pulumi.ResourceOption) (*CredentialSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CredentialSet
	err := ctx.RegisterResource("azure-native:containerregistry/v20230101preview:CredentialSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCredentialSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialSetState, opts ...pulumi.ResourceOption) (*CredentialSet, error) {
	var resource CredentialSet
	err := ctx.ReadResource("azure-native:containerregistry/v20230101preview:CredentialSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type credentialSetState struct {
}

type CredentialSetState struct {
}

func (CredentialSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialSetState)(nil)).Elem()
}

type credentialSetArgs struct {
	AuthCredentials   []AuthCredential    `pulumi:"authCredentials"`
	CredentialSetName *string             `pulumi:"credentialSetName"`
	Identity          *IdentityProperties `pulumi:"identity"`
	LoginServer       *string             `pulumi:"loginServer"`
	RegistryName      string              `pulumi:"registryName"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
}


type CredentialSetArgs struct {
	AuthCredentials   AuthCredentialArrayInput
	CredentialSetName pulumi.StringPtrInput
	Identity          IdentityPropertiesPtrInput
	LoginServer       pulumi.StringPtrInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (CredentialSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialSetArgs)(nil)).Elem()
}

type CredentialSetInput interface {
	pulumi.Input

	ToCredentialSetOutput() CredentialSetOutput
	ToCredentialSetOutputWithContext(ctx context.Context) CredentialSetOutput
}

func (*CredentialSet) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialSet)(nil)).Elem()
}

func (i *CredentialSet) ToCredentialSetOutput() CredentialSetOutput {
	return i.ToCredentialSetOutputWithContext(context.Background())
}

func (i *CredentialSet) ToCredentialSetOutputWithContext(ctx context.Context) CredentialSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialSetOutput)
}

type CredentialSetOutput struct{ *pulumi.OutputState }

func (CredentialSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialSet)(nil)).Elem()
}

func (o CredentialSetOutput) ToCredentialSetOutput() CredentialSetOutput {
	return o
}

func (o CredentialSetOutput) ToCredentialSetOutputWithContext(ctx context.Context) CredentialSetOutput {
	return o
}

func (o CredentialSetOutput) AuthCredentials() AuthCredentialResponseArrayOutput {
	return o.ApplyT(func(v *CredentialSet) AuthCredentialResponseArrayOutput { return v.AuthCredentials }).(AuthCredentialResponseArrayOutput)
}

func (o CredentialSetOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o CredentialSetOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CredentialSet) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o CredentialSetOutput) LoginServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringPtrOutput { return v.LoginServer }).(pulumi.StringPtrOutput)
}

func (o CredentialSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CredentialSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CredentialSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CredentialSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CredentialSetOutput{})
}
