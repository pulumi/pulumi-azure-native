


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupPolicy struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput        `pulumi:"name"`
	Properties BackupPolicyResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput   `pulumi:"systemData"`
	Type       pulumi.StringOutput        `pulumi:"type"`
}


func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dataprotection:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210201preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210601preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220501:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupPolicy
	err := ctx.RegisterResource("azure-native:dataprotection/v20210701:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("azure-native:dataprotection/v20210701:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupPolicyState struct {
}

type BackupPolicyState struct {
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	BackupPolicyName  *string           `pulumi:"backupPolicyName"`
	Properties        *BackupPolicyType `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	VaultName         string            `pulumi:"vaultName"`
}


type BackupPolicyArgs struct {
	BackupPolicyName  pulumi.StringPtrInput
	Properties        BackupPolicyTypePtrInput
	ResourceGroupName pulumi.StringInput
	VaultName         pulumi.StringInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupPolicyOutput) Properties() BackupPolicyResponseOutput {
	return o.ApplyT(func(v *BackupPolicy) BackupPolicyResponseOutput { return v.Properties }).(BackupPolicyResponseOutput)
}

func (o BackupPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BackupPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupPolicyOutput{})
}
