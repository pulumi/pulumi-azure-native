


package netapp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupPolicy struct {
	pulumi.CustomResourceState

	DailyBackupsToKeep   pulumi.IntPtrOutput              `pulumi:"dailyBackupsToKeep"`
	Enabled              pulumi.BoolPtrOutput             `pulumi:"enabled"`
	Location             pulumi.StringOutput              `pulumi:"location"`
	MonthlyBackupsToKeep pulumi.IntPtrOutput              `pulumi:"monthlyBackupsToKeep"`
	Name                 pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput              `pulumi:"provisioningState"`
	Tags                 pulumi.StringMapOutput           `pulumi:"tags"`
	Type                 pulumi.StringOutput              `pulumi:"type"`
	VolumeBackups        VolumeBackupsResponseArrayOutput `pulumi:"volumeBackups"`
	VolumesAssigned      pulumi.IntPtrOutput              `pulumi:"volumesAssigned"`
	WeeklyBackupsToKeep  pulumi.IntPtrOutput              `pulumi:"weeklyBackupsToKeep"`
	YearlyBackupsToKeep  pulumi.IntPtrOutput              `pulumi:"yearlyBackupsToKeep"`
}


func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
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
			Type: pulumi.String("azure-native:netapp/v20200501:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupPolicy
	err := ctx.RegisterResource("azure-native:netapp:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("azure-native:netapp:BackupPolicy", name, id, state, &resource, opts...)
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
	AccountName          string            `pulumi:"accountName"`
	BackupPolicyName     *string           `pulumi:"backupPolicyName"`
	DailyBackupsToKeep   *int              `pulumi:"dailyBackupsToKeep"`
	Enabled              *bool             `pulumi:"enabled"`
	Location             *string           `pulumi:"location"`
	MonthlyBackupsToKeep *int              `pulumi:"monthlyBackupsToKeep"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Tags                 map[string]string `pulumi:"tags"`
	VolumeBackups        []VolumeBackups   `pulumi:"volumeBackups"`
	VolumesAssigned      *int              `pulumi:"volumesAssigned"`
	WeeklyBackupsToKeep  *int              `pulumi:"weeklyBackupsToKeep"`
	YearlyBackupsToKeep  *int              `pulumi:"yearlyBackupsToKeep"`
}


type BackupPolicyArgs struct {
	AccountName          pulumi.StringInput
	BackupPolicyName     pulumi.StringPtrInput
	DailyBackupsToKeep   pulumi.IntPtrInput
	Enabled              pulumi.BoolPtrInput
	Location             pulumi.StringPtrInput
	MonthlyBackupsToKeep pulumi.IntPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	VolumeBackups        VolumeBackupsArrayInput
	VolumesAssigned      pulumi.IntPtrInput
	WeeklyBackupsToKeep  pulumi.IntPtrInput
	YearlyBackupsToKeep  pulumi.IntPtrInput
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

func (o BackupPolicyOutput) DailyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.DailyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func (o BackupPolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BackupPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BackupPolicyOutput) MonthlyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.MonthlyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func (o BackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BackupPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BackupPolicyOutput) VolumeBackups() VolumeBackupsResponseArrayOutput {
	return o.ApplyT(func(v *BackupPolicy) VolumeBackupsResponseArrayOutput { return v.VolumeBackups }).(VolumeBackupsResponseArrayOutput)
}

func (o BackupPolicyOutput) VolumesAssigned() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.VolumesAssigned }).(pulumi.IntPtrOutput)
}

func (o BackupPolicyOutput) WeeklyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.WeeklyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func (o BackupPolicyOutput) YearlyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.YearlyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupPolicyOutput{})
}
