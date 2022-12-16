


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupVault struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput              `pulumi:"eTag"`
	Identity   DppIdentityDetailsResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties BackupVaultResponseOutput           `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewBackupVault(ctx *pulumi.Context,
	name string, args *BackupVaultArgs, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dataprotection:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210101:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210201preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210601preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210701:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220101:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220901preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221001preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221101preview:BackupVault"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupVault
	err := ctx.RegisterResource("azure-native:dataprotection/v20220501:BackupVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupVaultState, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	var resource BackupVault
	err := ctx.ReadResource("azure-native:dataprotection/v20220501:BackupVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupVaultState struct {
}

type BackupVaultState struct {
}

func (BackupVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultState)(nil)).Elem()
}

type backupVaultArgs struct {
	ETag              *string             `pulumi:"eTag"`
	Identity          *DppIdentityDetails `pulumi:"identity"`
	Location          *string             `pulumi:"location"`
	Properties        BackupVaultType     `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
	VaultName         *string             `pulumi:"vaultName"`
}


type BackupVaultArgs struct {
	ETag              pulumi.StringPtrInput
	Identity          DppIdentityDetailsPtrInput
	Location          pulumi.StringPtrInput
	Properties        BackupVaultTypeInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringPtrInput
}

func (BackupVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultArgs)(nil)).Elem()
}

type BackupVaultInput interface {
	pulumi.Input

	ToBackupVaultOutput() BackupVaultOutput
	ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput
}

func (*BackupVault) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVault)(nil)).Elem()
}

func (i *BackupVault) ToBackupVaultOutput() BackupVaultOutput {
	return i.ToBackupVaultOutputWithContext(context.Background())
}

func (i *BackupVault) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultOutput)
}

type BackupVaultOutput struct{ *pulumi.OutputState }

func (BackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVault)(nil)).Elem()
}

func (o BackupVaultOutput) ToBackupVaultOutput() BackupVaultOutput {
	return o
}

func (o BackupVaultOutput) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return o
}

func (o BackupVaultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o BackupVaultOutput) Identity() DppIdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v *BackupVault) DppIdentityDetailsResponsePtrOutput { return v.Identity }).(DppIdentityDetailsResponsePtrOutput)
}

func (o BackupVaultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BackupVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupVaultOutput) Properties() BackupVaultResponseOutput {
	return o.ApplyT(func(v *BackupVault) BackupVaultResponseOutput { return v.Properties }).(BackupVaultResponseOutput)
}

func (o BackupVaultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BackupVault) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BackupVaultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BackupVaultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupVaultOutput{})
}
