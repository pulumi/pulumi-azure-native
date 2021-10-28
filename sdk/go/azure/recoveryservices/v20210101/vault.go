


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vault struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput        `pulumi:"etag"`
	Identity   IdentityDataResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringOutput           `pulumi:"location"`
	Name       pulumi.StringOutput           `pulumi:"name"`
	Properties VaultPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput          `pulumi:"sku"`
	SystemData SystemDataResponseOutput      `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput        `pulumi:"tags"`
	Type       pulumi.StringOutput           `pulumi:"type"`
}


func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210101:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160601:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20160601:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20200202:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20200202:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20201001:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210210:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210301:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210401:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210601:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210701:Vault"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210801:Vault"),
		},
	})
	opts = append(opts, aliases)
	var resource Vault
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210101:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("azure-native:recoveryservices/v20210101:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vaultState struct {
}

type VaultState struct {
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	Etag              *string           `pulumi:"etag"`
	Identity          *IdentityData     `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	Properties        *VaultProperties  `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         *string           `pulumi:"vaultName"`
}


type VaultArgs struct {
	Etag              pulumi.StringPtrInput
	Identity          IdentityDataPtrInput
	Location          pulumi.StringPtrInput
	Properties        VaultPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringPtrInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

type VaultOutput struct{ *pulumi.OutputState }

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VaultInput)(nil)).Elem(), &Vault{})
	pulumi.RegisterOutputType(VaultOutput{})
}
