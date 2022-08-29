


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Backup struct {
	pulumi.CustomResourceState

	BackupId            pulumi.StringOutput    `pulumi:"backupId"`
	BackupType          pulumi.StringOutput    `pulumi:"backupType"`
	CreationDate        pulumi.StringOutput    `pulumi:"creationDate"`
	FailureReason       pulumi.StringOutput    `pulumi:"failureReason"`
	Label               pulumi.StringPtrOutput `pulumi:"label"`
	Location            pulumi.StringOutput    `pulumi:"location"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput    `pulumi:"provisioningState"`
	Size                pulumi.Float64Output   `pulumi:"size"`
	Type                pulumi.StringOutput    `pulumi:"type"`
	UseExistingSnapshot pulumi.BoolPtrOutput   `pulumi:"useExistingSnapshot"`
	VolumeName          pulumi.StringOutput    `pulumi:"volumeName"`
}


func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	if isZero(args.UseExistingSnapshot) {
		args.UseExistingSnapshot = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Backup"),
		},
	})
	opts = append(opts, aliases)
	var resource Backup
	err := ctx.RegisterResource("azure-native:netapp/v20210601:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("azure-native:netapp/v20210601:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupState struct {
}

type BackupState struct {
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	AccountName         string  `pulumi:"accountName"`
	BackupName          *string `pulumi:"backupName"`
	Label               *string `pulumi:"label"`
	Location            *string `pulumi:"location"`
	PoolName            string  `pulumi:"poolName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	UseExistingSnapshot *bool   `pulumi:"useExistingSnapshot"`
	VolumeName          string  `pulumi:"volumeName"`
}


type BackupArgs struct {
	AccountName         pulumi.StringInput
	BackupName          pulumi.StringPtrInput
	Label               pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	PoolName            pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	UseExistingSnapshot pulumi.BoolPtrInput
	VolumeName          pulumi.StringInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func (o BackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

func (o BackupOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

func (o BackupOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o BackupOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

func (o BackupOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

func (o BackupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BackupOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *Backup) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

func (o BackupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BackupOutput) UseExistingSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.UseExistingSnapshot }).(pulumi.BoolPtrOutput)
}

func (o BackupOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.VolumeName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupOutput{})
}
