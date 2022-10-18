


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupInstance struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput          `pulumi:"name"`
	Properties BackupInstanceResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput     `pulumi:"systemData"`
	Type       pulumi.StringOutput          `pulumi:"type"`
}


func NewBackupInstance(ctx *pulumi.Context,
	name string, args *BackupInstanceArgs, opts ...pulumi.ResourceOption) (*BackupInstance, error) {
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
			Type: pulumi.String("azure-native:dataprotection:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210101:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210201preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210601preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220101:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220501:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220901preview:BackupInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupInstance
	err := ctx.RegisterResource("azure-native:dataprotection/v20210701:BackupInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupInstanceState, opts ...pulumi.ResourceOption) (*BackupInstance, error) {
	var resource BackupInstance
	err := ctx.ReadResource("azure-native:dataprotection/v20210701:BackupInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupInstanceState struct {
}

type BackupInstanceState struct {
}

func (BackupInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceState)(nil)).Elem()
}

type backupInstanceArgs struct {
	BackupInstanceName *string             `pulumi:"backupInstanceName"`
	Properties         *BackupInstanceType `pulumi:"properties"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	VaultName          string              `pulumi:"vaultName"`
}


type BackupInstanceArgs struct {
	BackupInstanceName pulumi.StringPtrInput
	Properties         BackupInstanceTypePtrInput
	ResourceGroupName  pulumi.StringInput
	VaultName          pulumi.StringInput
}

func (BackupInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceArgs)(nil)).Elem()
}

type BackupInstanceInput interface {
	pulumi.Input

	ToBackupInstanceOutput() BackupInstanceOutput
	ToBackupInstanceOutputWithContext(ctx context.Context) BackupInstanceOutput
}

func (*BackupInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstance)(nil)).Elem()
}

func (i *BackupInstance) ToBackupInstanceOutput() BackupInstanceOutput {
	return i.ToBackupInstanceOutputWithContext(context.Background())
}

func (i *BackupInstance) ToBackupInstanceOutputWithContext(ctx context.Context) BackupInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceOutput)
}

type BackupInstanceOutput struct{ *pulumi.OutputState }

func (BackupInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstance)(nil)).Elem()
}

func (o BackupInstanceOutput) ToBackupInstanceOutput() BackupInstanceOutput {
	return o
}

func (o BackupInstanceOutput) ToBackupInstanceOutputWithContext(ctx context.Context) BackupInstanceOutput {
	return o
}

func (o BackupInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupInstanceOutput) Properties() BackupInstanceResponseOutput {
	return o.ApplyT(func(v *BackupInstance) BackupInstanceResponseOutput { return v.Properties }).(BackupInstanceResponseOutput)
}

func (o BackupInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BackupInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BackupInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupInstanceOutput{})
}
