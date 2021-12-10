


package v20210101

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
	})
	opts = append(opts, aliases)
	var resource BackupVault
	err := ctx.RegisterResource("azure-native:dataprotection/v20210101:BackupVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupVaultState, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	var resource BackupVault
	err := ctx.ReadResource("azure-native:dataprotection/v20210101:BackupVault", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*BackupVault)(nil))
}

func (i *BackupVault) ToBackupVaultOutput() BackupVaultOutput {
	return i.ToBackupVaultOutputWithContext(context.Background())
}

func (i *BackupVault) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultOutput)
}

type BackupVaultOutput struct{ *pulumi.OutputState }

func (BackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupVault)(nil))
}

func (o BackupVaultOutput) ToBackupVaultOutput() BackupVaultOutput {
	return o
}

func (o BackupVaultOutput) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackupVaultOutput{})
}
